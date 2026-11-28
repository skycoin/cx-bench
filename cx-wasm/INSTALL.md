# cx-wasm

cx itself built for GOOS=js GOARCH=wasm (cx.wasm), running the .cx source on its vm under node with Go's wasm_exec_node.js. It parses and lowers in the Load phase. VARIANT=turbofan adds --no-liftoff.

## Install

Arch Linux:

```
sudo pacman -S --needed go git nodejs
```

Debian or Ubuntu:

```
sudo apt install golang-go git nodejs
```

macOS (Homebrew, not tested):

```
brew install go git node
```

Then `./precheck.sh cx-wasm` from the repository root lists what is still missing.

## Versions of the first results

skycoin-cx at the commit in out/cx-tools/rev.txt (the first results: see results/), built with go 1.27.3, node 26.12.1.

## Use

```
cx-wasm/build.sh
cx-wasm/run.sh fib 100               # one run of two passes
cx-wasm/profile.sh fib 100           # profile into results/profiles
```

cx is built from source by src/cx-tools.sh: CX_SRC=<skycoin-cx checkout> uses a local tree, otherwise it clones github.com/skycoin/skycoin-cx at CX_REF (main) into out/cx-src.

The vm subset of CX (the code cx-vm, cx-image and cx-llvm run) reads its arguments, so run.sh passes the size and the zero as the two arguments, but it has no clock and no stderr, so the program writes no phase markers, and the harness reports the phases as estimates: Load is the empty program's whole run, and the rest is split over the two passes. There is no server mode and no -time flag.

profile.sh runs the program under `node --cpu-prof`.
