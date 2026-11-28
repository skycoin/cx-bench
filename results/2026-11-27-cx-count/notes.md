# Notes on the CX count rerun

Only the count benchmark of the five CX forms was measured again, after the
CX programs started to read the size and the zero from the command line
(skycoin-cx#611). `raw.jsonl` holds those runs and, for the reference and
the precomputed mark, the c and c-wasm count and empty runs copied from
results/2026-11-27 and results/2026-11-27-wasm. skycoin-cx was at commit
322c0376. The machine is the one of the first results, and its one minute
load average was 2.1 to 11 during these runs, so the vm forms carry a wider
spread.

| form | count, time per round | before | C |
|---|---|---|---|
| cx-llvm | 360.1 us est | 1.16 ns est precomputed | 446.1 us |
| cx-llvm-wasm | 703.6 us est | | 609.6 us (c-wasm) |
| cx-llvm-wasm/turbofan | 350.6 us est | | 609.6 us (c-wasm) |
| cx-image | 34.9 ms est | 33.8 ms est | 446.1 us |
| cx-vm | 35.1 ms est | 46.8 ms est | 446.1 us |

cx-llvm no longer computes the count loop at build time, so it is no longer
marked precomputed, and it now runs the million additions a round as C
does. The "before" column is skycoin-cx 01deb81a, so the vm forms differ by
the cx revision as well as by the arguments, which they never folded.
