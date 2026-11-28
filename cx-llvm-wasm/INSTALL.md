# cx-llvm-wasm

`cx build --target=llvm --llvm-linker=tinygo --llvm-tinygo-target=wasip1`, the LLVM target linked by TinyGo for wasm32-wasip1, run under node:wasi. VARIANT=turbofan adds --no-liftoff.

## Install

Arch Linux:

```
sudo pacman -S --needed go git tinygo nodejs
```

Debian or Ubuntu:

```
sudo apt install golang-go git nodejs   # and TinyGo from its releases
```

macOS (Homebrew, not tested):

```
brew install go git node tinygo-org/tools/tinygo
```

Then `./precheck.sh cx-llvm-wasm` from the repository root lists what is still missing.

## Versions of the first results

skycoin-cx at the commit in out/cx-tools/rev.txt (the first results: see results/), built with go 1.27.3, TinyGo 0.44.0-dev (LLVM 23.1.3), node 26.12.1.

## Use

```
cx-llvm-wasm/build.sh
cx-llvm-wasm/run.sh fib 100               # one run of two passes
cx-llvm-wasm/profile.sh fib 100           # profile into results/profiles
```

cx is built from source by src/cx-tools.sh: CX_SRC=<skycoin-cx checkout> uses a local tree, otherwise it clones github.com/skycoin/skycoin-cx at CX_REF (main) into out/cx-src.

The vm subset of CX (the code cx-vm, cx-image and cx-llvm run) reads its arguments, so run.sh passes the size and the zero as the two arguments, but it has no clock and no stderr, so the program writes no phase markers, and the harness reports the phases as estimates: Load is the empty program's whole run, and the rest is split over the two passes. There is no server mode and no -time flag.

profile.sh runs the program under `node --cpu-prof`.
