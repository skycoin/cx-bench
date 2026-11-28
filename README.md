# cx-bench

Cross language benchmarks for CX: the same twelve programs in CX, Go, C, Lua
and Python, native and in wasm, with one harness that times them all on one
clock and checks every result.

The question it answers: per benchmark, how does the best CX form compare
with the best Lua, the best Go and C, and is the difference larger than the
noise of the measurement.

## Forms

A form is one way to run the programs. Each has a folder with `build.sh`,
`run.sh`, `profile.sh` and `INSTALL.md` (exact install steps and the tool
versions of the first results).

| form | what runs | phase markers | server mode |
|---|---|---|---|
| c | src/c, clang -O2 | yes | yes |
| c-wasm | src/c, clang -O2 for wasm32-wasip1, node:wasi | yes | yes |
| go | src/go, gc | yes | yes |
| go-wasm | src/go, GOOS=js, node with wasm_exec_node.js | yes | yes |
| tinygo | src/go, TinyGo -opt=2 | yes | yes |
| tinygo-wasm | src/go, TinyGo wasip1, node:wasi | yes | yes |
| lua | src/lua54 on lua5.4 and lua5.5 (variants 5.4, 5.5) | yes | yes |
| luajit | src/luajit on LuaJIT 2.1 | yes | yes |
| lua-wasm | src/lua54 on Lua 5.4.9 built for wasip1, node:wasi | yes | yes |
| gopher-lua | src/lua51 on gopher-lua (Lua 5.1 in Go), 9 of 12 | yes | yes |
| python | src/python on CPython 3 | yes | yes |
| cx-vm | src/cx on `cx --executor=vm` | no | no |
| cx-image | src/cx built with `--target=image`, run by cx-core-image | no | no |
| cx-llvm | src/cx built with `--target=llvm`, native through clang | no | no |
| cx-llvm-wasm | src/cx, `--target=llvm` linked by TinyGo for wasip1, node:wasi | no | no |
| cx-wasm | cx itself built for GOOS=js (cx.wasm), its vm running src/cx, node | no | no |

The node forms have the variants `default` (node as it ships) and
`turbofan` (`--no-liftoff`, every function compiled by TurboFan).

The CX old loop (`--executor=old`) is not a form. Bash versions are left
out: only seven of the twelve existed, they ran at a smaller size with an
estimated scale factor, and Bash cannot reach a two second run of most of
them without minutes per run.

### Why the sources live in src/

Forms that run the same program share one copy: `src/go` serves go,
go-wasm, tinygo and tinygo-wasm, `src/c` serves c and c-wasm, `src/lua54`
serves lua and lua-wasm, and `src/cx` serves the five cx forms. A copy per
form would let the copies drift apart, and a shared copy makes it plain that
two forms differ only in how the same source is compiled and run. Each form
folder holds what differs: how to build, how to start, how to profile and
what to install. LuaJIT and Lua 5.1 need their own sources (no 64 bit
integers), and Python has its own.

## The programs

Each benchmark runs a fixed unit of work, a round, `size` times. The size
comes from the command line, so no compiler sees it at build time. Each
round mixes its result into an FNV-1a hash, and the program checks every
round against the math and prints `<name> pass`.

A round of fib, ackermann, factorialiter, factorialrec and struct is the same
pure computation every time, so clang and TinyGo moved it out of the loop of
rounds and ran it once, and LLVM turned the count loop into its closed form.
These rounds add a zero the compiler cannot see as a constant: `fib(25 +
zero)`, and `sum += i + zero` in count. It is a volatile in C, an atomic
load in Go, a volatile load in TinyGo and a plain variable in the
interpreters, which precompute nothing. CX reads it as the second argument,
see below. The report marks a form that is under a tenth of
C on a benchmark as precomputed (the compiler worked out the result at build
time and removed the loop) and leaves it out of the best of table.

| # | name | one round |
|---:|---|---|
| 1 | count | sum 0 to 999,999 into an int64 |
| 2 | fib | recursive fib(25), 242,785 calls |
| 3 | array | bubble sort of a reversed array of 150, 11,175 swaps |
| 4 | struct | sum of 100 two field structs through a pointer |
| 5 | memtest | memtest86 style passes over a 4 MiB uint32 array |
| 6 | primes | sieve below 2,000,000, Miller-Rabin below 300,000, Lucas-Lehmer to 2^61-1 twenty times |
| 7 | ackermann | Ackermann(3, 5), 42,438 calls |
| 8 | digitalroot | digital root and persistence of 1 to 99,999 |
| 9 | factorialiter | n! for n 0 to 20 with a loop |
| 10 | factorialrec | n! for n 0 to 20 by recursion |
| 11 | structs | fill and sum a struct of two 100,000 element arrays |
| 12 | structserialize | encode and decode a point and 16 int32 in a 72 byte buffer |

The numbers are fixed and used by server mode. `expected.txt` holds the hash
of one pass at each size the results use. `harness/hashmath.go` computes it
from the math, closed forms and an independent digit sum, without running a
benchmark, so any size is checkable, and its test checks it against the
hashes the Go form printed. Every language is written independently and
agrees with it at every size checked.

## The protocol

Every program, in every language:

```
run.sh <name|number> <size> [-time=ms|us]
```

writes one character to stderr at each phase: `S` when main starts, `1`
after the first pass, `2` after an identical second pass, `F` just before it
exits, after it printed the hash and the pass line on stdout. The harness
timestamps each byte as it arrives on the pipe and reports

* Load, process start to `S`: exec, the runtime and, for an interpreter,
  reading the program
* Loop 1, `S` to `1`: the first pass, setup included
* Loop 2, `1` to `2`: the same pass again, warm
* Finish, `2` to exit

`-time=us` (or `ms`) makes the program also print `time us loop1=.. loop2=..`
from its own monotonic clock, as a cross check (Lua 5.4 and gopher-lua use
`os.clock`, CPU time, as standard Lua has no monotonic wall clock).

Server mode, `run.sh -server`: the program writes `S` and reads lines
`<number> <size> <unit>` from stdin, for example `3 400000 us`. It runs that
benchmark for one pass and writes `<number> <hash> <elapsed>` with the
elapsed time from its own clock, then waits for the next line. The same line
again shows warm up. `0` exits. Start up is paid once.

`run.sh empty` starts and exits, for the start up time.

### CX

The vm subset of CX, the code cx-vm, cx-image and cx-llvm run, reads its
arguments and stdin through the terminal library of skycoin-cx
(`os.ArgRead` and `os.InputLineRead`, decision D142), but has no clock
(`time.UnixNano`), no `os.Exit` and no stderr. So for the CX forms:

* The size and the zero are the two arguments, `cx program.cx ++<size> ++0`
  for cx-vm and cx-wasm and `program <size> 0` for the built forms, which
  `benchArg` of each program reads as decimal numbers. No compiler sees
  either at build time, as in the other languages.
* There are no phase markers. The harness takes Load as the whole run of the
  empty CX program and splits the rest over the two passes. These phases are
  marked `est`.
* There is no server mode and no `-time` flag.

Until skycoin-cx#611 the size was written into a copy of the program and
the zero was a global, so LLVM computed the count loop at build time and
the first results mark cx-llvm count precomputed.

A clock and a byte written to stderr from the vm would give CX the same
protocol as the other languages.

## Running

```
./precheck.sh                  # every tool of every form, found or missing
./precheck.sh c go cx-llvm     # exits 1 if one of these misses a tool
./run-all.sh native            # build, verify, time and report into results/<date>
./run-all.sh wasm results/<date>-wasm
./run-all.sh c,cx-llvm         # some forms
```

Step by step:

```
go/build.sh && cx-llvm/build.sh
(cd harness && go build -o ../out/harness .)
out/harness verify -forms go,cx-llvm            # hashes, markers, server mode
out/harness measure -forms go,cx-llvm -out results/x
out/harness report -in results/x                # results.md and chart.png
cx-llvm/run.sh fib 50                           # one run by hand
c/profile.sh fib 2000                           # a profile into results/profiles
```

Builds go to `out/`, which is not committed. The cx forms build cx from
skycoin-cx: `CX_SRC=<checkout>` uses a local tree, otherwise
`src/cx-tools.sh` clones github.com/skycoin/skycoin-cx at `CX_REF` (main).

## How the harness measures

For each form and benchmark:

1. The empty program runs ten times. Its median whole run is the start up.
2. Sizing: the size grows from 1 until one run takes at least 2 s and at
   least 100 times the start up, but at most 10 s (`-max`): node with
   `--no-liftoff` compiles a whole Go module before it starts, so 100 times
   its start up would be minutes a run. A slow form stops at a small size,
   a fast one at a large size, so each runs about the same time.
3. Five timed runs at that size.
4. The size sweep: one run each at size/8, size/4 and size/2.
5. Code placement, for the forms with two layouts (c, go, tinygo, cx-vm,
   cx-image, cx-llvm): five runs of a second build whose code sits at
   another address (src/layout). The interpreters are system binaries and
   have no second layout.
6. Server mode, three runs of one process, for the forms that have it.
7. One run with `-time=us`, the self timing cross check.

Every run records the one minute load average when it starts. Runs are
sequential, never in parallel.

## Reading the results

`results/<date>/results.md` is generated by `harness report` from
`raw.jsonl` (every run, kept).

* Time per round is the time of the two passes over twice the size. It is
  what compares across forms that ran at different sizes.
* Spread is (max - min) / median of the five runs.
* Placement is the change of the median in the second layout. Linker
  placement alone can move a hot loop by tens of percent.
* The sweep is the whole run, start up included, per round. Falling then
  flat means the start up is amortized at the size used. A per round time
  far below the other compiled forms means the compiler removed the work.
* A difference counts only when it is larger than the spread and the
  placement delta of both forms. The best of table marks it `sig yes`.

The first results are in
[results/2026-11-27](results/2026-11-27/results.md), with a summary in
[notes.md](results/2026-11-27/notes.md). The wasm forms under node are in
[results/2026-11-27-wasm](results/2026-11-27-wasm/results.md) with their
[notes.md](results/2026-11-27-wasm/notes.md). The count benchmark of the CX
forms with the size and the zero read from the command line is in
[results/2026-11-27-cx-count](results/2026-11-27-cx-count/results.md), with
its [notes.md](results/2026-11-27-cx-count/notes.md).

## CI

The workflow builds the forms whose tools install from apt (c, go, lua,
luajit, gopher-lua, python and the native cx forms, cx built from
skycoin-cx main) and runs `harness verify`. It times nothing.

## License

BSD 3-clause, the license of skycoin-cx, see LICENSE.
