# cx-image

`cx build --target=image`, run by cx-core-image on the vm. The image is built before the timed run, so Load is reading the image. LAYOUT=b uses the cx built with the vm loop moved.

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

Then `./precheck.sh cx-image` from the repository root lists what is still missing.

## Versions of the first results

skycoin-cx at the commit in out/cx-tools/rev.txt (the first results: see results/), built with go 1.27.3.

## Use

```
cx-image/build.sh
cx-image/run.sh fib 100               # one run of two passes
cx-image/profile.sh fib 100           # profile into results/profiles
```

cx is built from source by src/cx-tools.sh: CX_SRC=<skycoin-cx checkout> uses a local tree, otherwise it clones github.com/skycoin/skycoin-cx at CX_REF (main) into out/cx-src.

The vm subset of CX (the code cx-vm, cx-image and cx-llvm run) reads its arguments, so run.sh passes the size and the zero as the two arguments, but it has no clock and no stderr, so the program writes no phase markers, and the harness reports the phases as estimates: Load is the empty program's whole run, and the rest is split over the two passes. There is no server mode and no -time flag.

profile.sh runs the image under perf on cx-core-image.
