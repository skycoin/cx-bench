package main

import (
	"flag"
	"fmt"
)

// verifyCommand runs every benchmark of every form at a small size and
// checks the hash, the pass line, the marker order and server mode.
func verifyCommand(root string, args []string) error {
	flags := flag.NewFlagSet("verify", flag.ExitOnError)
	formList := flags.String("forms", "native", "form folders, comma separated, or native, wasm, all")
	size := flags.Int("size", 2, "rounds per pass")
	flags.Parse(args)
	forms, err := formsLoad(root, formNamesParse(*formList))
	if err != nil {
		return err
	}
	expected, err := expectedLoad(root)
	if err != nil {
		return err
	}
	failures := 0
	for _, f := range forms {
		layouts := []string{"a"}
		if f.caps["layout"] {
			layouts = append(layouts, "b")
		}
		if r := emptyRun(f, "a"); !r.Pass {
			fmt.Printf("FAIL %-18s empty: %s\n", f.label(), r.Error)
			failures++
		}
		for _, bench := range benchRows {
			skipped := false
			for _, layout := range layouts {
				r := oneShot(f, layout, bench, *size, expected.want)
				status := "ok  "
				switch {
				case r.Pass:
				case notRun(r):
					status = "skip"
					skipped = true
				default:
					status = "FAIL"
					failures++
				}
				fmt.Printf("%s %-18s %-15s layout %s size %d %s %s\n", status, f.label(), bench, layout, *size, r.Hash, r.Error)
			}
			if f.caps["server"] && !skipped && expected.want(bench, *size) != "" {
				r := serverRun(f, bench, *size, 2)
				if !r.Pass || r.Hash != expected.want(bench, *size) {
					fmt.Printf("FAIL %-18s %-15s server %s %s\n", f.label(), bench, r.Hash, r.Error)
					failures++
				} else {
					fmt.Printf("ok   %-18s %-15s server %s\n", f.label(), bench, r.Hash)
				}
			}
		}
	}
	if expected.added {
		if err := expected.save(); err != nil {
			return err
		}
	}
	if failures > 0 {
		return fmt.Errorf("%d checks failed", failures)
	}
	return nil
}
