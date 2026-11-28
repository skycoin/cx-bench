# cx-llvm

`cx build --target=llvm`, the program compiled to LLVM IR and by clang to a native program. LAYOUT=b links src/layout/pad.c first through cx-llvm/clang-pad.sh. As the size is written into the source, LLVM sees it as a constant; the size sweep and the comparison with C show whether that changed the work.

## Install

Arch Linux:

```
sudo pacman -S --needed go git clang perf
```

Debian or Ubuntu:

```
sudo apt install golang-go git clang linux-perf
```

macOS (Homebrew, not tested):

```
brew install go git llvm
```

Then `./precheck.sh cx-llvm` from the repository root lists what is still missing.

## Versions of the first results

skycoin-cx at the commit in out/cx-tools/rev.txt (the first results: see results/), built with go 1.27.3, clang 23.1.4.

## Use

```
cx-llvm/build.sh
cx-llvm/run.sh fib 100               # one run of two passes
cx-llvm/profile.sh fib 100           # profile into results/profiles
```

cx is built from source by src/cx-tools.sh: CX_SRC=<skycoin-cx checkout> uses a local tree, otherwise it clones github.com/skycoin/skycoin-cx at CX_REF (main) into out/cx-src.

The vm subset of CX (the code cx-vm, cx-image and cx-llvm run) reads its arguments, so run.sh passes the size and the zero as the two arguments, but it has no clock and no stderr, so the program writes no phase markers, and the harness reports the phases as estimates: Load is the empty program's whole run, and the rest is split over the two passes. There is no server mode and no -time flag.

profile.sh runs the program under perf. Build with --llvm-debug for CX line names, see skycoin-cx docs/LlvmTarget.md.
