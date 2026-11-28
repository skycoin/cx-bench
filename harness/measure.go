package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

// measureSettings are the flags of harness measure.
type measureSettings struct {
	runs      int
	minimum   time.Duration
	startupX  int
	maximum   time.Duration
	benches   []string
	sizeLimit int
}

// measureCommand sizes each benchmark of each form until a run takes at
// least -min and -startup times the form's empty program, then records
// -runs runs, the size sweep, the layout b runs and the server mode runs.
func measureCommand(root string, args []string) error {
	flags := flag.NewFlagSet("measure", flag.ExitOnError)
	formList := flags.String("forms", "native", "form folders, comma separated, or native, wasm, all")
	out := flags.String("out", "", "results directory, results/<date> by default")
	runs := flags.Int("runs", 5, "timed runs per benchmark and layout")
	minimum := flags.Duration("min", 2*time.Second, "shortest run")
	startupX := flags.Int("startup", 100, "shortest run in multiples of the empty program")
	maximum := flags.Duration("max", 10*time.Second, "longest target run, for forms whose start up is long")
	skip := flags.String("skip", "", "form labels to leave out, comma separated, lua/5.4 for example")
	benchList := flags.String("benches", strings.Join(benchNames, ","), "benchmarks, comma separated")
	flags.Parse(args)
	if *out == "" {
		*out = filepath.Join(root, "results", time.Now().UTC().Format("2006-01-02"))
	}
	settings := measureSettings{runs: *runs, minimum: *minimum, startupX: *startupX, maximum: *maximum, benches: strings.Split(*benchList, ","), sizeLimit: 1 << 30}
	forms, err := formsLoad(root, formNamesParse(*formList))
	if err != nil {
		return err
	}
	forms = formsSkip(forms, *skip)
	expected, err := expectedLoad(root)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(*out, 0o755); err != nil {
		return err
	}
	rawFile, err := os.OpenFile(filepath.Join(*out, "raw.jsonl"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer rawFile.Close()
	encoder := json.NewEncoder(rawFile)
	record := func(r run) {
		encoder.Encode(r)
		status := "ok"
		if !r.Pass {
			status = "FAIL " + r.Error
		}
		fmt.Printf("%s %-7s %-18s %-15s %-1s size %-10d %9.1f ms load %.2f %s\n", time.Now().UTC().Format("15:04:05"), r.Kind, r.Form, r.Bench, r.Layout, r.Size, float64(r.TotalNs)/1e6, r.LoadAvg, status)
	}
	record(run{Kind: "machine", Form: "-", Pass: true, Note: fmt.Sprintf("%s/%s, %d logical CPUs, started %s", runtime.GOOS, runtime.GOARCH, runtime.NumCPU(), time.Now().UTC().Format(time.RFC3339))})
	for _, f := range forms {
		for _, layout := range layoutsOf(f) {
			var totals []int64
			for i := 0; i < 2*settings.runs; i++ {
				r := emptyRun(f, layout)
				record(r)
				if r.Pass {
					totals = append(totals, r.TotalNs)
				}
			}
			if len(totals) == 0 {
				break
			}
			if layout != "a" {
				continue
			}
			empty := time.Duration(median(totals))
			target := settings.minimum
			if floor := empty * time.Duration(settings.startupX); floor > target {
				target = floor
			}
			if target > settings.maximum {
				target = settings.maximum
				record(run{Kind: "note", Form: f.label(), Pass: true, Note: fmt.Sprintf("start up %v, runs capped at %v, under %dx start up", empty, settings.maximum, settings.startupX)})
			}
			for _, bench := range settings.benches {
				measureBench(f, bench, empty, target, settings, expected.want, record)
			}
		}
		if expected.added {
			if err := expected.save(); err != nil {
				return err
			}
		}
	}
	return nil
}

func layoutsOf(f form) []string {
	if f.caps["layout"] {
		return []string{"a", "b"}
	}
	return []string{"a"}
}

func measureBench(f form, bench string, empty, target time.Duration, settings measureSettings, want func(string, int) string, record func(run)) {
	size := 1
	for {
		r := oneShot(f, "a", bench, size, want)
		r.Kind = "size"
		if strings.Contains(r.Error, "no Lua 5.1 version") {
			return
		}
		record(r)
		if !r.Pass {
			return
		}
		total := time.Duration(r.TotalNs)
		if total >= target {
			break
		}
		if size >= settings.sizeLimit {
			record(run{Kind: "note", Form: f.label(), Bench: bench, Size: size, Pass: true, Note: "size limit reached below the target time, the work is precomputed"})
			break
		}
		work := total - empty
		if work < time.Millisecond {
			work = time.Millisecond
		}
		// A JIT makes small runs a poor guide, so the size grows at most 10
		// times a step, and 3 times once a run is a tenth of the target.
		factor := float64(target-empty) * 1.15 / float64(work)
		ceiling := 10.0
		if total*10 >= target {
			ceiling = 3
		}
		factor = min(max(factor, 1.2), ceiling)
		size = min(int(float64(size)*factor)+1, settings.sizeLimit)
	}
	for i := 0; i < settings.runs; i++ {
		record(oneShot(f, "a", bench, size, want))
	}
	for _, divisor := range []int{8, 4, 2} {
		if size/divisor >= 1 {
			r := oneShot(f, "a", bench, size/divisor, want)
			r.Kind = "sweep"
			record(r)
		}
	}
	if f.caps["layout"] {
		for i := 0; i < settings.runs; i++ {
			record(oneShot(f, "b", bench, size, want))
		}
	}
	if f.caps["server"] {
		r := serverRun(f, bench, size, 3)
		if expected := want(bench, size); r.Pass && expected != "" && r.Hash != expected {
			r.Pass, r.Error = false, "server hash "+r.Hash+", want "+expected
		}
		record(r)
	}
	if f.caps["selftime"] {
		r := oneShot(f, "a", bench, size, want, "-time=us")
		r.Kind = "selftime"
		record(r)
	}
}

func median(values []int64) int64 {
	sorted := append([]int64{}, values...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
	if len(sorted) == 0 {
		return 0
	}
	if len(sorted)%2 == 1 {
		return sorted[len(sorted)/2]
	}
	return (sorted[len(sorted)/2-1] + sorted[len(sorted)/2]) / 2
}

func runsLoad(path string) ([]run, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var runs []run
	for _, line := range strings.Split(string(data), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var r run
		if err := json.Unmarshal([]byte(line), &r); err != nil {
			return nil, fmt.Errorf("%s: %v", path, err)
		}
		runs = append(runs, r)
	}
	return runs, nil
}

func formsSkip(forms []form, list string) []form {
	if list == "" {
		return forms
	}
	skipped := map[string]bool{}
	for _, label := range strings.Split(list, ",") {
		skipped[label] = true
	}
	var kept []form
	for _, f := range forms {
		if !skipped[f.label()] {
			kept = append(kept, f)
		}
	}
	return kept
}
