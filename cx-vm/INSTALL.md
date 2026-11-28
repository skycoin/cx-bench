# cx-vm

`cx --executor=vm`, the flat vm of skycoin-cx, running the .cx source of src/cx. The program is parsed and lowered in the Load phase. LAYOUT=b runs a cx built with the layout pads of src/layout, which move the vm loop.

## Install

Arch Linux:

```
sudo pacman -S --needed go git perf
```

Debian or Ubuntu:

```
sudo apt install golang-go git linux-perf
```

macOS (Homebrew, not tested):

```
brew install go git 
```

Then `./precheck.sh cx-vm` from the repository root lists what is still missing.

## Versions of the first results

skycoin-cx at the commit in out/cx-tools/rev.txt (the first results: see results/), built with go 1.27.3.

## Use

```
cx-vm/build.sh
cx-vm/run.sh fib 100               # one run of two passes
cx-vm/profile.sh fib 100           # profile into results/profiles
```

cx is built from source by src/cx-tools.sh: CX_SRC=<skycoin-cx checkout> uses a local tree, otherwise it clones github.com/skycoin/skycoin-cx at CX_REF (main) into out/cx-src.

The vm subset of CX (the code cx-vm, cx-image and cx-llvm run) reads its arguments, so run.sh passes the size and the zero as the two arguments, but it has no clock and no stderr, so the program writes no phase markers, and the harness reports the phases as estimates: Load is the empty program's whole run, and the rest is split over the two passes. There is no server mode and no -time flag.

profile.sh runs the program under perf on the cx binary, then with `cx -debug-profile`, which times each compile phase and writes a pprof of parsing only. A pprof of the vm run loop would be a small addition to cx.
