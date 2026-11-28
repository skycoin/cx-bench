package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// benchNames are the benchmarks in their fixed order, number 1 first.
var benchNames = []string{"count", "fib", "array", "struct", "memtest", "primes", "ackermann",
	"digitalroot", "factorialiter", "factorialrec", "structs", "structserialize", "fibiter", "ackermanniter"}

// benchRows are the benchmarks in the order of the runs, the report and the
// chart, each loop variant next to its recursive variant (decision D145).
var benchRows = []string{"count", "fib", "fibiter", "array", "struct", "memtest", "primes", "ackermann",
	"ackermanniter", "digitalroot", "factorialiter", "factorialrec", "structs", "structserialize"}

// notRun tells whether a form left a benchmark out on purpose: its run.sh
// printed "not run: <reason>" (FORM_NOT_RUN in src/form.sh).
func notRun(r run) bool { return strings.Contains(r.Error, "not run: ") }

// formsNative and formsWasm are every form folder, native first.
var (
	formsNative = []string{"c", "go", "tinygo", "lua", "luajit", "gopher-lua", "python", "cx-vm", "cx-image", "cx-llvm"}
	formsWasm   = []string{"c-wasm", "go-wasm", "tinygo-wasm", "lua-wasm", "cx-llvm-wasm", "cx-wasm"}
)

// form is one form folder with one variant, lua with 5.4 for example.
type form struct {
	root    string
	name    string
	variant string
	caps    map[string]bool
}

func (f form) label() string {
	if f.variant == "" || f.variant == "default" {
		return f.name
	}
	return f.name + "/" + f.variant
}

func (f form) environment(layout string) []string {
	env := append(os.Environ(), "CXBENCH_PREOPEN="+f.root)
	if layout != "" {
		env = append(env, "LAYOUT="+layout)
	}
	if f.variant != "" {
		env = append(env, "VARIANT="+f.variant)
	}
	return env
}

// script runs the form's run.sh with the arguments and returns its stdout
// lines. Its stderr, the messages of a build, is returned with an error.
func (f form) script(layout string, args ...string) ([]string, error) {
	cmd := exec.Command("bash", append([]string{filepath.Join(f.root, f.name, "run.sh")}, args...)...)
	cmd.Env = f.environment(layout)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s run.sh %s: %v: %s", f.label(), strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	var lines []string
	for _, line := range strings.Split(strings.TrimRight(stdout.String(), "\n"), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

// formsLoad expands form folder names into one form per variant.
func formsLoad(root string, names []string) ([]form, error) {
	var forms []form
	for _, name := range names {
		base := form{root: root, name: name}
		capsLine, err := base.script("", "--caps")
		if err != nil {
			return nil, err
		}
		variants, err := base.script("", "--variants")
		if err != nil {
			return nil, err
		}
		caps := map[string]bool{}
		for _, line := range capsLine {
			for _, word := range strings.Fields(line) {
				caps[word] = true
			}
		}
		for _, line := range variants {
			for _, variant := range strings.Fields(line) {
				forms = append(forms, form{root: root, name: name, variant: variant, caps: caps})
			}
		}
	}
	return forms, nil
}

func formNamesParse(list string) []string {
	switch list {
	case "", "native":
		return formsNative
	case "wasm":
		return formsWasm
	case "all":
		return append(append([]string{}, formsNative...), formsWasm...)
	}
	return strings.Split(list, ",")
}
