package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// benchStats is what the report shows of one benchmark on one form.
type benchStats struct {
	form, bench string
	size        int
	nsRound     []float64 // per timed run, layout a
	median      float64
	spread      float64 // (max - min) / median, percent
	placement   float64 // |median b - median a| / median a, percent, -1 without layout b
	estimated   bool    // no markers: phases from the empty program
	phases      [4]float64
	total       float64
	sweep       [][2]float64 // size, total ns per round with start up
	server      []float64    // ns per round of each server mode run
	selfDiff    float64      // self timed loops over harness loops - 1, percent, NaN without
	loadMin     float64
	loadMax     float64
	note        string
	precomputed bool // under a tenth of C: the compiler removed the work
}

// reportGroups are the four columns of the best of table. A form counts in
// the first group whose prefix list holds its name.
var reportGroups = []struct {
	name  string
	forms []string
}{
	{"CX", []string{"cx-llvm", "cx-vm", "cx-image", "cx-llvm-wasm", "cx-wasm"}},
	{"Lua", []string{"lua", "luajit", "lua-wasm"}},
	{"Go", []string{"go", "tinygo", "go-wasm", "tinygo-wasm"}},
	{"C", []string{"c", "c-wasm"}},
}

func formBase(label string) string { return strings.SplitN(label, "/", 2)[0] }

func statsCompute(runs []run) (map[string]map[string]*benchStats, []string, map[string]float64) {
	emptyTotals := map[string][]int64{}
	for _, r := range runs {
		if r.Kind == "empty" && r.Pass {
			emptyTotals[r.Form+" "+r.Layout] = append(emptyTotals[r.Form+" "+r.Layout], r.TotalNs)
		}
	}
	empty := map[string]float64{}
	for key, totals := range emptyTotals {
		empty[key] = float64(median(totals))
	}
	stats := map[string]map[string]*benchStats{}
	var forms []string
	layoutB := map[string][]float64{}
	get := func(r run) *benchStats {
		if stats[r.Form] == nil {
			stats[r.Form] = map[string]*benchStats{}
			forms = append(forms, r.Form)
		}
		s := stats[r.Form][r.Bench]
		if s == nil {
			s = &benchStats{form: r.Form, bench: r.Bench, placement: -1, selfDiff: math.NaN(), loadMin: math.Inf(1)}
			stats[r.Form][r.Bench] = s
		}
		return s
	}
	perRound := func(r run) (float64, bool) {
		if s, ok1 := r.MarksNs["S"]; ok1 {
			if two, ok2 := r.MarksNs["2"]; ok2 {
				return float64(two-s) / float64(2*r.Size), false
			}
		}
		return (float64(r.TotalNs) - empty[r.Form+" "+r.Layout]) / float64(2*r.Size), true
	}
	for _, r := range runs {
		if r.Kind == "note" && r.Bench != "" {
			get(r).note = r.Note
			continue
		}
		if r.Bench == "" || !r.Pass {
			continue
		}
		s := get(r)
		switch {
		case r.Kind == "run" && r.Layout == "a":
			ns, estimated := perRound(r)
			s.size = r.Size
			s.nsRound = append(s.nsRound, ns)
			s.estimated = estimated
			s.loadMin = math.Min(s.loadMin, r.LoadAvg)
			s.loadMax = math.Max(s.loadMax, r.LoadAvg)
			s.total += float64(r.TotalNs)
			if estimated {
				start := empty[r.Form+" a"]
				loop := (float64(r.TotalNs) - start) / 2
				s.phases[0] += start
				s.phases[1] += loop
				s.phases[2] += loop
			} else {
				s.phases[0] += float64(r.MarksNs["S"])
				s.phases[1] += float64(r.MarksNs["1"] - r.MarksNs["S"])
				s.phases[2] += float64(r.MarksNs["2"] - r.MarksNs["1"])
				s.phases[3] += float64(r.TotalNs - r.MarksNs["2"])
			}
		case r.Kind == "run" && r.Layout == "b":
			ns, _ := perRound(r)
			layoutB[r.Form+" "+r.Bench] = append(layoutB[r.Form+" "+r.Bench], ns)
		case r.Kind == "sweep":
			s.sweep = append(s.sweep, [2]float64{float64(r.Size), float64(r.TotalNs) / float64(2*r.Size)})
		case r.Kind == "server":
			for _, us := range r.Server {
				s.server = append(s.server, float64(us)*1000/float64(r.Size))
			}
		case r.Kind == "selftime" && len(r.SelfUs) == 2:
			if two, ok := r.MarksNs["2"]; ok {
				harness := float64(two-r.MarksNs["S"]) / 1000
				s.selfDiff = (float64(r.SelfUs[0]+r.SelfUs[1])/harness - 1) * 100
			}
		}
	}
	for _, byBench := range stats {
		for _, s := range byBench {
			n := float64(len(s.nsRound))
			if n == 0 {
				continue
			}
			for i := range s.phases {
				s.phases[i] /= n
			}
			s.total /= n
			sorted := append([]float64{}, s.nsRound...)
			sort.Float64s(sorted)
			s.median = floatMedian(sorted)
			s.spread = (sorted[len(sorted)-1] - sorted[0]) / s.median * 100
			s.sweep = append(s.sweep, [2]float64{float64(s.size), s.total / float64(2*s.size)})
			if b := layoutB[s.form+" "+s.bench]; len(b) > 0 {
				sort.Float64s(b)
				s.placement = math.Abs(floatMedian(b)-s.median) / s.median * 100
			}
		}
	}
	return stats, forms, empty
}

func floatMedian(sorted []float64) float64 {
	n := len(sorted)
	if n%2 == 1 {
		return sorted[n/2]
	}
	return (sorted[n/2-1] + sorted[n/2]) / 2
}

// best returns the fastest form of a group for a benchmark.
func best(stats map[string]map[string]*benchStats, forms []string, group []string, bench string) *benchStats {
	var found *benchStats
	for _, label := range forms {
		member := false
		for _, name := range group {
			member = member || formBase(label) == name
		}
		s := stats[label][bench]
		if !member || s == nil || len(s.nsRound) == 0 || s.precomputed {
			continue
		}
		if found == nil || s.median < found.median {
			found = s
		}
	}
	return found
}

// noise is the larger of the run to run spread and the placement delta.
func noise(s *benchStats) float64 { return math.Max(s.spread, math.Max(s.placement, 0)) }

// significant tells whether a and b differ by more than the noise of both.
func significant(a, b *benchStats) bool {
	difference := math.Abs(a.median-b.median) / math.Min(a.median, b.median) * 100
	return difference > math.Max(noise(a), noise(b))
}

func formatNs(ns float64) string {
	switch {
	case ns >= 1e6:
		return fmt.Sprintf("%.1f ms", ns/1e6)
	case ns >= 1e3:
		return fmt.Sprintf("%.1f us", ns/1e3)
	case ns >= 10:
		return fmt.Sprintf("%.0f ns", ns)
	}
	return fmt.Sprintf("%.2f ns", ns)
}

func formatPercent(value float64) string {
	if value < 0 || math.IsNaN(value) {
		return "-"
	}
	return fmt.Sprintf("%.1f", value)
}

func reportCommand(root string, args []string) error {
	flags := flag.NewFlagSet("report", flag.ExitOnError)
	in := flags.String("in", "", "results directory holding raw.jsonl")
	title := flags.String("title", "cx-bench results", "heading of results.md")
	machine := flags.String("machine", "", "generic description of the machine")
	flags.Parse(args)
	if *in == "" {
		return fmt.Errorf("report: -in is required")
	}
	runs, err := runsLoad(filepath.Join(*in, "raw.jsonl"))
	if err != nil {
		return err
	}
	stats, forms, empty := statsCompute(runs)
	precomputedMark(stats, forms)
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n", *title)
	if *machine != "" {
		fmt.Fprintf(&b, "%s\n\n", *machine)
	}
	for _, r := range runs {
		if r.Kind == "machine" {
			fmt.Fprintf(&b, "Harness: %s.\n\n", r.Note)
		}
	}
	b.WriteString("Time per round is the time of the two passes over two times the size, from the phase markers. ")
	b.WriteString("A form without markers (the CX forms) has it from the whole run less its empty program, marked est. ")
	b.WriteString("A difference to C counts (sig yes) only when it is larger than the run to run spread and the placement delta of both forms. ")
	b.WriteString("A form under a tenth of C on a benchmark is marked precomputed: its compiler worked out the result at build time and removed the work, and the best of table leaves it out.\n\n")
	b.WriteString("![best of chart](chart.png)\n\n")

	b.WriteString("## Best of each language, per benchmark\n\n")
	b.WriteString("| benchmark | best CX | x C | best Lua | x C | best Go | x C | C | sig CX vs C |\n|---|---|---:|---|---:|---|---:|---:|---|\n")
	var chartRows []chartRow
	for _, bench := range benchNames {
		row := chartRow{bench: bench}
		c := best(stats, forms, reportGroups[3].forms, bench)
		fmt.Fprintf(&b, "| %s |", bench)
		for gi := 0; gi < 3; gi++ {
			s := best(stats, forms, reportGroups[gi].forms, bench)
			if s == nil {
				b.WriteString(" - | - |")
				continue
			}
			ratio := "-"
			if c != nil {
				ratio = fmt.Sprintf("%.2f", s.median/c.median)
				row.ratios[gi] = s.median / c.median
			}
			est := ""
			if s.estimated {
				est = " est"
			}
			fmt.Fprintf(&b, " %s %s%s | %s |", s.form, formatNs(s.median), est, ratio)
		}
		cx := best(stats, forms, reportGroups[0].forms, bench)
		if c != nil {
			fmt.Fprintf(&b, " %s |", formatNs(c.median))
			row.ratios[3] = 1
		} else {
			b.WriteString(" - |")
		}
		if cx != nil && c != nil {
			fmt.Fprintf(&b, " %s |\n", yesNo(significant(cx, c)))
		} else {
			b.WriteString(" - |\n")
		}
		chartRows = append(chartRows, row)
	}

	b.WriteString("\n## Phases of the best forms\n\n")
	b.WriteString("Load is process start to S, Loop 1 and Loop 2 the two identical passes, Finish from 2 to exit. Percent of the whole run.\n\n")
	b.WriteString("| benchmark | language | form | size | Load | Loop 1 | Loop 2 | Finish | Load % | Loop 1 % | Loop 2 % | Finish % | spread % | placement % | load avg |\n")
	b.WriteString("|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, bench := range benchNames {
		for _, group := range reportGroups {
			s := best(stats, forms, group.forms, bench)
			if s == nil {
				continue
			}
			fmt.Fprintf(&b, "| %s | %s | %s |", bench, group.name, s.form)
			b.WriteString(phaseCells(s))
		}
	}

	b.WriteString("\n## Every form\n\n")
	b.WriteString("Sweep is the whole run, start up included, per round at size/8, /4, /2 and the size: flat means the start up is amortized. ")
	b.WriteString("Server is the time per round the program measured itself in server mode, first and third run of one process. ")
	b.WriteString("Self is how far the program's own -time=us loop times are from the harness markers.\n\n")
	b.WriteString("| form | benchmark | size | per round | min | max | spread % | placement % | sweep /8 | /4 | /2 | x1 | server 1st | 3rd | self % | load avg |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, label := range forms {
		for _, bench := range benchNames {
			s := stats[label][bench]
			if s == nil || len(s.nsRound) == 0 {
				continue
			}
			sorted := append([]float64{}, s.nsRound...)
			sort.Float64s(sorted)
			est := ""
			if s.estimated {
				est = " est"
			}
			if s.precomputed {
				est += " precomputed"
			}
			fmt.Fprintf(&b, "| %s | %s | %d | %s%s | %s | %s | %s | %s |", label, bench, s.size, formatNs(s.median), est,
				formatNs(sorted[0]), formatNs(sorted[len(sorted)-1]), formatPercent(s.spread), formatPercent(s.placement))
			sweep := append([][2]float64{}, s.sweep...)
			sort.Slice(sweep, func(i, j int) bool { return sweep[i][0] < sweep[j][0] })
			for len(sweep) < 4 {
				sweep = append([][2]float64{{0, -1}}, sweep...)
			}
			for _, point := range sweep[len(sweep)-4:] {
				if point[1] < 0 {
					b.WriteString(" - |")
				} else {
					fmt.Fprintf(&b, " %s |", formatNs(point[1]))
				}
			}
			if len(s.server) >= 3 {
				fmt.Fprintf(&b, " %s | %s |", formatNs(s.server[0]), formatNs(s.server[2]))
			} else {
				b.WriteString(" - | - |")
			}
			fmt.Fprintf(&b, " %s | %.2f-%.2f |\n", signedPercent(s.selfDiff), s.loadMin, s.loadMax)
			if s.note != "" {
				fmt.Fprintf(&b, "| %s | %s | | %s | | | | | | | | | | | | |\n", label, bench, s.note)
			}
		}
	}

	var formNotes []string
	for _, r := range runs {
		if r.Kind == "note" && r.Bench == "" {
			formNotes = append(formNotes, fmt.Sprintf("* %s: %s\n", r.Form, r.Note))
		}
	}
	if len(formNotes) > 0 {
		b.WriteString("\n## Notes\n\n")
		for _, note := range formNotes {
			b.WriteString(note)
		}
	}

	b.WriteString("\n## Empty program\n\n| form | layout | start to exit | start to S |\n|---|---|---:|---:|\n")
	keys := make([]string, 0, len(empty))
	for key := range empty {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		var starts []int64
		fields := strings.Fields(key)
		for _, r := range runs {
			if r.Kind == "empty" && r.Pass && r.Form == fields[0] && r.Layout == fields[1] {
				if s, ok := r.MarksNs["S"]; ok {
					starts = append(starts, s)
				}
			}
		}
		toS := "-"
		if len(starts) > 0 {
			toS = formatNs(float64(median(starts)))
		}
		fmt.Fprintf(&b, "| %s | %s | %s | %s |\n", fields[0], fields[1], formatNs(empty[key]), toS)
	}
	if err := os.WriteFile(filepath.Join(*in, "results.md"), []byte(b.String()), 0o644); err != nil {
		return err
	}
	return chartWrite(filepath.Join(*in, "chart.png"), chartRows, *title)
}

func phaseCells(s *benchStats) string {
	total := s.phases[0] + s.phases[1] + s.phases[2] + s.phases[3]
	finish := formatNs(s.phases[3])
	finishPercent := fmt.Sprintf("%.1f", s.phases[3]/total*100)
	load := formatNs(s.phases[0])
	if s.estimated {
		load += " est"
		finish, finishPercent = "in Load", "-"
	}
	return fmt.Sprintf(" %d | %s | %s | %s | %s | %.1f | %.1f | %.1f | %s | %s | %s | %.2f-%.2f |\n", s.size, load,
		formatNs(s.phases[1]), formatNs(s.phases[2]), finish, s.phases[0]/total*100, s.phases[1]/total*100,
		s.phases[2]/total*100, finishPercent, formatPercent(s.spread), formatPercent(s.placement), s.loadMin, s.loadMax)
}

func signedPercent(value float64) string {
	if math.IsNaN(value) {
		return "-"
	}
	return fmt.Sprintf("%+.1f", value)
}

func yesNo(value bool) string {
	if value {
		return "yes"
	}
	return "no"
}

// precomputedMark marks a form whose time per round is under a tenth of C's on
// the same benchmark: no form does the work ten times faster than clang
// -O2, so the compiler removed it (cx-llvm precomputed count while CX had no
// input a compiler cannot see). The best of table leaves such a form out.
func precomputedMark(stats map[string]map[string]*benchStats, forms []string) {
	for _, label := range forms {
		reference := "c"
		if strings.HasSuffix(formBase(label), "-wasm") {
			reference = "c-wasm"
		}
		if stats[reference] == nil {
			reference = "c"
		}
		for bench, s := range stats[label] {
			c := stats[reference][bench]
			if c != nil && len(c.nsRound) > 0 && len(s.nsRound) > 0 && label != reference && s.median < c.median/10 {
				s.precomputed = true
			}
		}
	}
}
