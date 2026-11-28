# cx-bench wasm results under node, 2026-11-27

Machine: the x86-64 desktop of the native results (8 logical CPUs, 16 GB RAM, Linux 7.2, the desktop session running beside). Node 26.12.1, as it ships (default) and with --no-liftoff (turbofan). Three timed runs per benchmark, not five, and runs capped at 10 s, see Notes. Tools: clang 23.1.4, Go 1.27.3, TinyGo 0.44.0-dev, Lua 5.4.9, skycoin-cx 01deb81af6fc.

Harness: linux/amd64, 8 logical CPUs, started 2026-11-27T09:39:42Z.

Harness: linux/amd64, 8 logical CPUs, started 2026-11-27T10:35:54Z.

Time per round is the time of the two passes over two times the size, from the phase markers. A form without markers (the CX forms) has it from the whole run less its empty program, marked est. A difference to C counts (sig yes) only when it is larger than the run to run spread and the placement delta of both forms. A form under a tenth of C on a benchmark is marked precomputed: its compiler worked out the result at build time and removed the work, and the best of table leaves it out.

![best of chart](chart.png)

## Best of each language, per benchmark

| benchmark | best CX | x C | best Lua | x C | best Go | x C | C | sig CX vs C |
|---|---|---:|---|---:|---|---:|---:|---|
| count | cx-wasm/turbofan 308.1 ms est | 758.61 | lua-wasm/turbofan 19.7 ms | 48.58 | tinygo-wasm/turbofan 667.6 us | 1.64 | 406.2 us | yes |
| fib | cx-llvm-wasm/turbofan 1.0 ms est | 2.97 | lua-wasm/turbofan 12.0 ms | 34.80 | tinygo-wasm/turbofan 345.4 us | 1.00 | 344.3 us | yes |
| array | cx-llvm-wasm/turbofan 15.0 us est | 1.00 | lua-wasm/turbofan 839.1 us | 55.95 | tinygo-wasm/turbofan 15.2 us | 1.01 | 15.0 us | no |
| struct | cx-llvm-wasm/turbofan 81 ns est | 1.17 | lua-wasm/turbofan 6.8 us | 97.88 | tinygo-wasm/turbofan 65 ns | 0.94 | 69 ns | yes |
| memtest | cx-llvm-wasm/turbofan 19.9 ms est | 1.07 | lua-wasm/turbofan 778.5 ms | 41.84 | tinygo-wasm/turbofan 19.3 ms | 1.04 | 18.6 ms | yes |
| primes | cx-llvm-wasm/turbofan 62.9 ms est | 0.99 | lua-wasm/turbofan 595.4 ms | 9.38 | tinygo-wasm/turbofan 64.0 ms | 1.01 | 63.5 ms | no |
| ackermann | cx-llvm-wasm/turbofan 196.3 us est | 2.78 | lua-wasm/turbofan 1.8 ms | 25.67 | tinygo-wasm/turbofan 71.6 us | 1.01 | 70.7 us | yes |
| digitalroot | cx-llvm-wasm/turbofan 2.0 ms est | 1.06 | lua-wasm/turbofan 71.2 ms | 38.21 | tinygo-wasm/turbofan 1.9 ms | 1.02 | 1.9 ms | yes |
| factorialiter | cx-llvm-wasm/turbofan 136 ns est | 0.99 | lua-wasm/turbofan 5.7 us | 41.92 | tinygo-wasm/turbofan 105 ns | 0.76 | 137 ns | no |
| factorialrec | cx-llvm-wasm/turbofan 433 ns est | 3.26 | lua-wasm/turbofan 12.7 us | 95.49 | tinygo-wasm/turbofan 123 ns | 0.93 | 133 ns | yes |
| structs | cx-llvm-wasm/turbofan 229.7 us est | 1.02 | lua-wasm/turbofan 10.3 ms | 45.82 | tinygo-wasm/turbofan 217.6 us | 0.96 | 225.8 us | yes |
| structserialize | cx-llvm-wasm/turbofan 21 ns est | 3.86 | lua-wasm/turbofan 8.2 us | 1520.54 | tinygo-wasm/turbofan 25 ns | 4.70 | 5.37 ns | yes |

## Phases of the best forms

Load is process start to S, Loop 1 and Loop 2 the two identical passes, Finish from 2 to exit. Percent of the whole run.

| benchmark | language | form | size | Load | Loop 1 | Loop 2 | Finish | Load % | Loop 1 % | Loop 2 % | Finish % | spread % | placement % | load avg |
|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|
| count | CX | cx-wasm/turbofan | 7 | 5978.0 ms est | 2161.9 ms | 2161.9 ms | in Load | 58.0 | 21.0 | 21.0 | - | 7.7 | - | 2.37-2.46 |
| count | Lua | lua-wasm/turbofan | 278 | 461.1 ms | 5546.3 ms | 5529.4 ms | 45.9 ms | 4.0 | 47.9 | 47.7 | 0.4 | 5.1 | - | 1.52-1.99 |
| count | Go | tinygo-wasm/turbofan | 8351 | 226.2 ms | 5575.6 ms | 5583.7 ms | 70.9 ms | 2.0 | 48.7 | 48.7 | 0.6 | 0.5 | - | 2.02-2.03 |
| count | C | c-wasm/turbofan | 15108 | 98.1 ms | 6095.8 ms | 6139.7 ms | 27.6 ms | 0.8 | 49.3 | 49.7 | 0.2 | 2.7 | - | 2.54-3.46 |
| fib | CX | cx-llvm-wasm/turbofan | 5398 | 140.5 ms est | 5498.2 ms | 5498.2 ms | in Load | 1.3 | 49.4 | 49.4 | - | 1.7 | - | 2.23-2.70 |
| fib | Lua | lua-wasm/turbofan | 458 | 463.5 ms | 5486.9 ms | 5408.9 ms | 44.1 ms | 4.1 | 48.1 | 47.4 | 0.4 | 6.4 | - | 3.95-4.07 |
| fib | Go | tinygo-wasm/turbofan | 15799 | 238.0 ms | 5478.7 ms | 5513.4 ms | 71.2 ms | 2.1 | 48.5 | 48.8 | 0.6 | 3.0 | - | 1.57-1.67 |
| fib | C | c-wasm/turbofan | 17321 | 100.2 ms | 5997.9 ms | 6033.0 ms | 27.0 ms | 0.8 | 49.3 | 49.6 | 0.2 | 3.3 | - | 2.02-2.44 |
| array | CX | cx-llvm-wasm/turbofan | 333334 | 140.5 ms est | 5016.3 ms | 5016.3 ms | in Load | 1.4 | 49.3 | 49.3 | - | 1.4 | - | 1.53-1.74 |
| array | Lua | lua-wasm/turbofan | 6432 | 443.8 ms | 5448.2 ms | 5378.9 ms | 41.1 ms | 3.9 | 48.2 | 47.6 | 0.4 | 2.1 | - | 2.60-2.90 |
| array | Go | tinygo-wasm/turbofan | 333334 | 226.5 ms | 5077.8 ms | 5063.0 ms | 70.9 ms | 2.2 | 48.6 | 48.5 | 0.7 | 0.5 | - | 2.04-2.24 |
| array | C | c-wasm/turbofan | 400001 | 110.6 ms | 6023.3 ms | 6031.3 ms | 29.7 ms | 0.9 | 49.4 | 49.5 | 0.2 | 1.5 | - | 1.40-1.61 |
| struct | CX | cx-llvm-wasm/turbofan | 71678894 | 140.5 ms est | 5773.0 ms | 5773.0 ms | in Load | 1.2 | 49.4 | 49.4 | - | 2.5 | - | 1.46-1.71 |
| struct | Lua | lua-wasm/turbofan | 805061 | 468.6 ms | 5500.4 ms | 5479.3 ms | 43.4 ms | 4.1 | 47.9 | 47.7 | 0.4 | 1.2 | - | 3.16-3.72 |
| struct | Go | tinygo-wasm/turbofan | 86348592 | 235.7 ms | 5613.3 ms | 5607.8 ms | 76.0 ms | 2.0 | 48.7 | 48.6 | 0.7 | 1.4 | - | 1.50-1.64 |
| struct | C | c-wasm/turbofan | 87917134 | 101.3 ms | 6136.7 ms | 6121.6 ms | 32.6 ms | 0.8 | 49.5 | 49.4 | 0.3 | 2.2 | - | 1.24-1.76 |
| memtest | CX | cx-llvm-wasm/turbofan | 282 | 140.5 ms est | 5614.9 ms | 5614.9 ms | in Load | 1.2 | 49.4 | 49.4 | - | 0.5 | - | 1.89-2.17 |
| memtest | Lua | lua-wasm/turbofan | 7 | 447.8 ms | 5531.6 ms | 5472.9 ms | 44.8 ms | 3.9 | 48.1 | 47.6 | 0.4 | 3.9 | - | 3.48-3.63 |
| memtest | Go | tinygo-wasm/turbofan | 291 | 230.2 ms | 5578.1 ms | 5627.5 ms | 73.5 ms | 2.0 | 48.5 | 48.9 | 0.6 | 1.2 | - | 1.57-1.60 |
| memtest | C | c-wasm/turbofan | 320 | 101.5 ms | 5999.0 ms | 5954.5 ms | 27.2 ms | 0.8 | 49.7 | 49.3 | 0.2 | 3.2 | - | 1.91-2.00 |
| primes | CX | cx-llvm-wasm/turbofan | 90 | 140.5 ms est | 5657.2 ms | 5657.2 ms | in Load | 1.2 | 49.4 | 49.4 | - | 1.6 | - | 2.77-3.00 |
| primes | Lua | lua-wasm/turbofan | 9 | 452.3 ms | 5394.6 ms | 5302.8 ms | 49.9 ms | 4.0 | 48.2 | 47.3 | 0.4 | 2.9 | - | 2.95-3.41 |
| primes | Go | tinygo-wasm/turbofan | 86 | 233.5 ms | 5545.0 ms | 5479.9 ms | 71.5 ms | 2.1 | 48.9 | 48.4 | 0.6 | 1.5 | - | 1.79-1.96 |
| primes | C | c-wasm/turbofan | 93 | 98.1 ms | 5897.9 ms | 5899.2 ms | 34.8 ms | 0.8 | 49.4 | 49.4 | 0.3 | 1.0 | - | 1.61-1.72 |
| ackermann | CX | cx-llvm-wasm/turbofan | 29027 | 140.5 ms est | 5698.2 ms | 5698.2 ms | in Load | 1.2 | 49.4 | 49.4 | - | 0.5 | - | 2.34-2.77 |
| ackermann | Lua | lua-wasm/turbofan | 3063 | 447.0 ms | 5626.3 ms | 5509.6 ms | 43.6 ms | 3.8 | 48.4 | 47.4 | 0.4 | 2.4 | - | 2.20-2.35 |
| ackermann | Go | tinygo-wasm/turbofan | 78566 | 231.2 ms | 5709.7 ms | 5888.0 ms | 76.8 ms | 1.9 | 48.0 | 49.5 | 0.6 | 9.1 | - | 1.75-1.85 |
| ackermann | C | c-wasm/turbofan | 85047 | 98.8 ms | 6047.4 ms | 6026.2 ms | 28.5 ms | 0.8 | 49.6 | 49.4 | 0.2 | 1.2 | - | 2.03-2.36 |
| digitalroot | CX | cx-llvm-wasm/turbofan | 2899 | 140.5 ms est | 5743.3 ms | 5743.3 ms | in Load | 1.2 | 49.4 | 49.4 | - | 0.5 | - | 2.13-2.27 |
| digitalroot | Lua | lua-wasm/turbofan | 83 | 446.6 ms | 5781.9 ms | 5804.8 ms | 42.6 ms | 3.7 | 47.9 | 48.1 | 0.4 | 7.0 | - | 1.50-1.67 |
| digitalroot | Go | tinygo-wasm/turbofan | 2921 | 230.3 ms | 5525.1 ms | 5520.1 ms | 74.2 ms | 2.0 | 48.7 | 48.6 | 0.7 | 3.0 | - | 1.92-2.01 |
| digitalroot | C | c-wasm/turbofan | 3234 | 103.1 ms | 5979.6 ms | 6011.5 ms | 28.8 ms | 0.9 | 49.3 | 49.6 | 0.2 | 2.8 | - | 1.89-1.95 |
| factorialiter | CX | cx-llvm-wasm/turbofan | 41423945 | 140.5 ms est | 5650.0 ms | 5650.0 ms | in Load | 1.2 | 49.4 | 49.4 | - | 0.9 | - | 2.20-3.76 |
| factorialiter | Lua | lua-wasm/turbofan | 821018 | 486.8 ms | 4837.6 ms | 4747.2 ms | 41.7 ms | 4.8 | 47.8 | 46.9 | 0.4 | 9.7 | - | 1.68-2.17 |
| factorialiter | Go | tinygo-wasm/turbofan | 53083622 | 227.2 ms | 5603.0 ms | 5538.9 ms | 73.8 ms | 2.0 | 49.0 | 48.4 | 0.6 | 1.6 | - | 1.66-2.01 |
| factorialiter | C | c-wasm/turbofan | 43977796 | 101.3 ms | 6026.1 ms | 6008.8 ms | 28.2 ms | 0.8 | 49.5 | 49.4 | 0.2 | 1.3 | - | 1.29-1.41 |
| factorialrec | CX | cx-llvm-wasm/turbofan | 13595644 | 140.5 ms est | 6094.1 ms | 6094.1 ms | in Load | 1.1 | 49.4 | 49.4 | - | 13.3 | - | 3.80-4.34 |
| factorialrec | Lua | lua-wasm/turbofan | 420043 | 443.3 ms | 5292.1 ms | 5326.7 ms | 41.3 ms | 4.0 | 47.7 | 48.0 | 0.4 | 2.7 | - | 2.01-2.24 |
| factorialrec | Go | tinygo-wasm/turbofan | 45905215 | 236.6 ms | 5715.1 ms | 5609.8 ms | 72.2 ms | 2.0 | 49.1 | 48.2 | 0.6 | 1.6 | - | 1.55-1.86 |
| factorialrec | C | c-wasm/turbofan | 46183860 | 106.4 ms | 6137.6 ms | 6164.6 ms | 29.7 ms | 0.9 | 49.3 | 49.6 | 0.2 | 1.8 | - | 2.49-2.75 |
| structs | CX | cx-llvm-wasm/turbofan | 24464 | 140.5 ms est | 5622.1 ms | 5622.1 ms | in Load | 1.2 | 49.4 | 49.4 | - | 0.3 | - | 2.91-3.30 |
| structs | Lua | lua-wasm/turbofan | 534 | 460.0 ms | 5511.7 ms | 5549.5 ms | 44.7 ms | 4.0 | 47.7 | 48.0 | 0.4 | 1.2 | - | 2.13-2.39 |
| structs | Go | tinygo-wasm/turbofan | 25406 | 238.6 ms | 5543.3 ms | 5533.4 ms | 69.5 ms | 2.1 | 48.7 | 48.6 | 0.6 | 0.8 | - | 1.94-2.16 |
| structs | C | c-wasm/turbofan | 26508 | 101.9 ms | 6012.3 ms | 5989.6 ms | 26.7 ms | 0.8 | 49.6 | 49.4 | 0.2 | 1.0 | - | 3.16-3.45 |
| structserialize | CX | cx-llvm-wasm/turbofan | 273115630 | 140.5 ms est | 5668.2 ms | 5668.2 ms | in Load | 1.2 | 49.4 | 49.4 | - | 1.3 | - | 1.88-2.10 |
| structserialize | Lua | lua-wasm/turbofan | 652637 | 464.0 ms | 5328.5 ms | 5288.6 ms | 41.0 ms | 4.2 | 47.9 | 47.6 | 0.4 | 1.5 | - | 2.72-3.09 |
| structserialize | Go | tinygo-wasm/turbofan | 222266634 | 236.2 ms | 5604.8 ms | 5627.4 ms | 76.1 ms | 2.0 | 48.5 | 48.7 | 0.7 | 1.0 | - | 1.98-2.14 |
| structserialize | C | c-wasm/turbofan | 1000000003 | 103.2 ms | 5418.3 ms | 5353.8 ms | 26.7 ms | 0.9 | 49.7 | 49.1 | 0.2 | 2.2 | - | 3.08-3.74 |

## Every form

Sweep is the whole run, start up included, per round at size/8, /4, /2 and the size: flat means the start up is amortized. Server is the time per round the program measured itself in server mode, first and third run of one process. Self is how far the program's own -time=us loop times are from the harness markers.

| form | benchmark | size | per round | min | max | spread % | placement % | sweep /8 | /4 | /2 | x1 | server 1st | 3rd | self % | load avg |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|
| c-wasm | count | 5413 | 609.6 us | 603.2 us | 615.4 us | 2.0 | - | 674.8 us | 622.5 us | 632.9 us | 615.3 us | 870.0 us | 398.5 us | -0.0 | 1.82-1.89 |
| c-wasm | fib | 7145 | 478.3 us | 476.1 us | 478.5 us | 0.5 | - | 512.5 us | 488.5 us | 474.5 us | 482.1 us | 482.9 us | 484.6 us | -0.0 | 1.76-1.97 |
| c-wasm | array | 165161 | 21.0 us | 20.2 us | 26.5 us | 30.3 | - | 31.7 us | 26.4 us | 21.6 us | 22.8 us | 25.3 us | 15.0 us | -0.0 | 2.11-2.33 |
| c-wasm | struct | 32953790 | 102 ns | 101 ns | 103 ns | 1.7 | - | 110 ns | 105 ns | 104 ns | 103 ns | 132 ns | 69 ns | -0.0 | 3.02-3.11 |
| c-wasm | memtest | 124 | 27.8 ms | 27.7 ms | 27.8 ms | 0.4 | - | 34.1 ms | 28.1 ms | 27.4 ms | 28.0 ms | 35.5 ms | 18.6 ms | -0.0 | 5.17-5.69 |
| c-wasm | primes | 45 | 76.8 ms | 76.1 ms | 77.4 ms | 1.7 | - | 83.9 ms | 81.5 ms | 79.9 ms | 77.5 ms | 88.6 ms | 62.4 ms | +0.0 | 4.01-4.62 |
| c-wasm | ackermann | 33334 | 90.6 us | 90.0 us | 90.9 us | 1.1 | - | 96.5 us | 94.8 us | 93.6 us | 91.4 us | 92.3 us | 89.9 us | -0.0 | 3.20-3.63 |
| c-wasm | digitalroot | 878 | 3.8 ms | 3.8 ms | 3.8 ms | 0.2 | - | 4.2 ms | 3.9 ms | 3.9 ms | 3.9 ms | 5.8 ms | 1.9 ms | -0.0 | 2.28-2.39 |
| c-wasm | factorialiter | 14491866 | 236 ns | 231 ns | 236 ns | 2.4 | - | 250 ns | 247 ns | 237 ns | 237 ns | 328 ns | 135 ns | -0.0 | 1.63-1.68 |
| c-wasm | factorialrec | 14804437 | 226 ns | 226 ns | 232 ns | 2.6 | - | 239 ns | 236 ns | 230 ns | 230 ns | 321 ns | 132 ns | -0.0 | 1.51-1.55 |
| c-wasm | structs | 12004 | 289.9 us | 289.5 us | 305.8 us | 5.6 | - | 318.4 us | 437.0 us | 330.5 us | 299.0 us | 340.6 us | 228.8 us | -0.0 | 1.95-2.20 |
| c-wasm | structserialize | 566534831 | 6.00 ns | 6.00 ns | 6.14 ns | 2.4 | - | 6.90 ns | 6.31 ns | 6.20 ns | 6.11 ns | 7.18 ns | 6.79 ns | -0.0 | 2.28-2.56 |
| c-wasm/turbofan | count | 15108 | 406.2 us | 398.9 us | 409.7 us | 2.7 | - | 433.5 us | 414.6 us | 409.8 us | 409.1 us | 405.1 us | 400.7 us | -0.0 | 2.54-3.46 |
| c-wasm/turbofan | fib | 17321 | 344.3 us | 343.2 us | 354.5 us | 3.3 | - | 369.7 us | 361.0 us | 356.4 us | 351.0 us | 341.8 us | 339.6 us | -0.0 | 2.02-2.44 |
| c-wasm/turbofan | array | 400001 | 15.0 us | 15.0 us | 15.2 us | 1.5 | - | 16.4 us | 15.7 us | 15.4 us | 15.2 us | 15.0 us | 15.3 us | -0.0 | 1.40-1.61 |
| c-wasm/turbofan | struct | 87917134 | 69 ns | 69 ns | 71 ns | 2.2 | - | 76 ns | 74 ns | 71 ns | 70 ns | 70 ns | 70 ns | -0.0 | 1.24-1.76 |
| c-wasm/turbofan | memtest | 320 | 18.6 ms | 18.4 ms | 19.0 ms | 3.2 | - | 21.0 ms | 19.4 ms | 19.2 ms | 18.9 ms | 18.8 ms | 18.7 ms | -0.0 | 1.91-2.00 |
| c-wasm/turbofan | primes | 93 | 63.5 ms | 63.1 ms | 63.7 ms | 1.0 | - | 71.7 ms | 67.4 ms | 66.2 ms | 64.1 ms | 64.2 ms | 63.7 ms | -0.0 | 1.61-1.72 |
| c-wasm/turbofan | ackermann | 85047 | 70.7 us | 70.7 us | 71.5 us | 1.2 | - | 77.7 us | 74.8 us | 72.5 us | 71.7 us | 71.2 us | 70.9 us | -0.0 | 2.03-2.36 |
| c-wasm/turbofan | digitalroot | 3234 | 1.9 ms | 1.8 ms | 1.9 ms | 2.8 | - | 2.0 ms | 1.9 ms | 1.9 ms | 1.9 ms | 1.8 ms | 1.8 ms | -0.0 | 1.89-1.95 |
| c-wasm/turbofan | factorialiter | 43977796 | 137 ns | 136 ns | 138 ns | 1.3 | - | 149 ns | 142 ns | 153 ns | 138 ns | 144 ns | 136 ns | -0.0 | 1.29-1.41 |
| c-wasm/turbofan | factorialrec | 46183860 | 133 ns | 132 ns | 135 ns | 1.8 | - | 146 ns | 136 ns | 137 ns | 135 ns | 133 ns | 131 ns | -0.0 | 2.49-2.75 |
| c-wasm/turbofan | structs | 26508 | 225.8 us | 225.5 us | 227.8 us | 1.0 | - | 249.8 us | 239.0 us | 233.8 us | 228.8 us | 226.6 us | 230.8 us | -0.0 | 3.16-3.45 |
| c-wasm/turbofan | structserialize | 1000000003 | 5.37 ns | 5.33 ns | 5.45 ns | 2.2 | - | 5.83 ns | 5.68 ns | 5.70 ns | 5.45 ns | 5.35 ns | 5.37 ns | -0.0 | 3.08-3.74 |
| go-wasm | count | 524 | 11.9 ms | 11.8 ms | 11.9 ms | 1.0 | - | 12.8 ms | 12.4 ms | 12.1 ms | 12.0 ms | 17.7 ms | 5.9 ms | -0.0 | 1.96-2.17 |
| go-wasm | fib | 3230 | 1.8 ms | 1.8 ms | 1.8 ms | 3.4 | - | 2.0 ms | 1.9 ms | 1.8 ms | 1.8 ms | 1.9 ms | 1.9 ms | -0.0 | 1.53-1.67 |
| go-wasm | array | 100003 | 58.3 us | 58.2 us | 58.6 us | 0.6 | - | 63.6 us | 60.3 us | 59.3 us | 58.9 us | 74.1 us | 41.1 us | -0.0 | 1.42-1.51 |
| go-wasm | struct | 16682013 | 373 ns | 370 ns | 376 ns | 1.6 | - | 404 ns | 392 ns | 375 ns | 376 ns | 495 ns | 258 ns | -0.0 | 1.61-1.92 |
| go-wasm | memtest | 68 | 91.3 ms | 90.6 ms | 91.4 ms | 1.0 | - | 96.9 ms | 94.7 ms | 93.0 ms | 91.9 ms | 115.8 ms | 66.7 ms | -0.0 | 1.70-1.85 |
| go-wasm | primes | 59 | 105.7 ms | 105.5 ms | 106.5 ms | 0.9 | - | 113.8 ms | 108.3 ms | 107.9 ms | 106.8 ms | 124.6 ms | 113.1 ms | -0.0 | 1.26-1.38 |
| go-wasm | ackermann | 13616 | 455.6 us | 454.0 us | 464.9 us | 2.4 | - | 486.4 us | 470.8 us | 478.2 us | 462.2 us | 461.7 us | 464.7 us | -0.0 | 3.41-3.53 |
| go-wasm | digitalroot | 1003 | 5.7 ms | 5.7 ms | 5.8 ms | 2.0 | - | 6.1 ms | 5.9 ms | 5.8 ms | 5.8 ms | 6.6 ms | 3.7 ms | -0.0 | 2.60-2.82 |
| go-wasm | factorialiter | 8302388 | 757 ns | 748 ns | 764 ns | 2.0 | - | 822 ns | 770 ns | 764 ns | 763 ns | 963 ns | 559 ns | -0.0 | 3.05-3.18 |
| go-wasm | factorialrec | 5604937 | 1.1 us | 1.1 us | 1.1 us | 2.0 | - | 1.2 us | 1.2 us | 1.1 us | 1.1 us | 1.2 us | 1.2 us | -0.0 | 3.31-3.46 |
| go-wasm | structs | 9191 | 648.1 us | 646.8 us | 650.3 us | 0.5 | - | 705.1 us | 666.5 us | 664.7 us | 654.6 us | 873.3 us | 425.5 us | -0.0 | 3.62-3.80 |
| go-wasm | structserialize | 18915359 | 332 ns | 332 ns | 340 ns | 2.3 | - | 351 ns | 345 ns | 334 ns | 338 ns | 415 ns | 268 ns | -0.0 | 2.86-3.15 |
| go-wasm/turbofan | count | 848 | 5.9 ms | 5.8 ms | 5.9 ms | 1.2 | - | 12.4 ms | 9.2 ms | 7.5 ms | 6.6 ms | 5.8 ms | 5.8 ms | -0.8 | 2.79-3.07 |
| go-wasm/turbofan | fib | 2519 | 1.8 ms | 1.8 ms | 1.8 ms | 2.9 | - | 4.0 ms | 2.9 ms | 2.4 ms | 2.1 ms | 1.8 ms | 1.8 ms | -0.7 | 3.88-4.26 |
| go-wasm/turbofan | array | 122208 | 41.0 us | 40.8 us | 41.8 us | 2.3 | - | 85.7 us | 62.9 us | 52.3 us | 46.6 us | 39.7 us | 42.6 us | -0.8 | 2.56-2.57 |
| go-wasm/turbofan | struct | 18026041 | 269 ns | 266 ns | 270 ns | 1.4 | - | 593 ns | 427 ns | 341 ns | 306 ns | 261 ns | 265 ns | -0.9 | 2.68-2.91 |
| go-wasm/turbofan | memtest | 74 | 67.9 ms | 67.5 ms | 68.5 ms | 1.5 | - | 143.1 ms | 106.5 ms | 84.5 ms | 77.2 ms | 66.9 ms | 66.3 ms | -0.8 | 2.83-2.98 |
| go-wasm/turbofan | primes | 54 | 90.4 ms | 90.1 ms | 92.3 ms | 2.5 | - | 213.2 ms | 147.6 ms | 116.3 ms | 103.3 ms | 90.3 ms | 88.9 ms | -0.7 | 1.84-2.17 |
| go-wasm/turbofan | ackermann | 9841 | 483.2 us | 477.6 us | 553.9 us | 15.8 | - | 1.0 ms | 759.1 us | 628.1 us | 574.3 us | 466.1 us | 480.5 us | -0.8 | 1.84-2.45 |
| go-wasm/turbofan | digitalroot | 1345 | 3.7 ms | 3.7 ms | 4.5 ms | 22.1 | - | 8.6 ms | 6.6 ms | 5.6 ms | 4.4 ms | 4.8 ms | 4.6 ms | -0.6 | 2.15-2.41 |
| go-wasm/turbofan | factorialiter | 7174453 | 625 ns | 528 ns | 656 ns | 20.4 | - | 1.4 us | 1.0 us | 745 ns | 695 ns | 568 ns | 550 ns | -1.0 | 1.73-1.86 |
| go-wasm/turbofan | factorialrec | 4748193 | 997 ns | 996 ns | 1.0 us | 0.5 | - | 2.2 us | 1.6 us | 1.3 us | 1.1 us | 980 ns | 991 ns | -0.8 | 1.60-1.84 |
| go-wasm/turbofan | structs | 9841 | 431.3 us | 431.0 us | 441.9 us | 2.5 | - | 1.0 ms | 711.7 us | 565.3 us | 501.1 us | 426.9 us | 433.8 us | -0.9 | 1.41-1.66 |
| go-wasm/turbofan | structserialize | 18031014 | 272 ns | 270 ns | 277 ns | 2.5 | - | 597 ns | 434 ns | 355 ns | 310 ns | 290 ns | 292 ns | -0.8 | 1.68-1.80 |
| tinygo-wasm | count | 2626 | 1.4 ms | 1.4 ms | 1.4 ms | 3.3 | - | 1.6 ms | 1.5 ms | 1.4 ms | 1.4 ms | 2.1 ms | 665.6 us | -0.0 | 1.65-1.68 |
| tinygo-wasm | fib | 7698 | 479.4 us | 474.8 us | 479.6 us | 1.0 | - | 509.4 us | 492.8 us | 483.9 us | 482.4 us | 470.7 us | 483.0 us | -0.0 | 1.65-1.71 |
| tinygo-wasm | array | 145239 | 25.7 us | 25.5 us | 25.7 us | 0.8 | - | 28.0 us | 26.6 us | 25.8 us | 25.9 us | 36.4 us | 15.3 us | -0.0 | 1.84-2.01 |
| tinygo-wasm | struct | 23443213 | 157 ns | 157 ns | 160 ns | 1.7 | - | 168 ns | 164 ns | 161 ns | 160 ns | 256 ns | 67 ns | -0.0 | 2.40-2.60 |
| tinygo-wasm | memtest | 85 | 44.1 ms | 43.9 ms | 46.0 ms | 4.6 | - | 46.6 ms | 45.6 ms | 60.2 ms | 45.1 ms | 76.6 ms | 19.8 ms | -0.0 | 1.99-2.28 |
| tinygo-wasm | primes | 46 | 80.2 ms | 80.1 ms | 80.9 ms | 1.0 | - | 89.3 ms | 82.6 ms | 82.1 ms | 81.1 ms | 99.8 ms | 64.4 ms | -0.0 | 2.19-2.41 |
| tinygo-wasm | ackermann | 41231 | 90.5 us | 90.2 us | 90.5 us | 0.3 | - | 95.6 us | 93.6 us | 91.8 us | 91.2 us | 90.4 us | 90.7 us | -0.0 | 1.86-1.96 |
| tinygo-wasm | digitalroot | 980 | 3.8 ms | 3.8 ms | 3.8 ms | 0.8 | - | 4.1 ms | 3.9 ms | 3.9 ms | 3.9 ms | 5.8 ms | 1.9 ms | -0.0 | 1.66-1.78 |
| tinygo-wasm | factorialiter | 13037938 | 285 ns | 280 ns | 286 ns | 2.1 | - | 312 ns | 296 ns | 291 ns | 286 ns | 506 ns | 105 ns | -0.0 | 2.17-2.22 |
| tinygo-wasm | factorialrec | 12000004 | 314 ns | 310 ns | 316 ns | 1.9 | - | 335 ns | 330 ns | 330 ns | 316 ns | 527 ns | 123 ns | -0.0 | 3.73-4.55 |
| tinygo-wasm | structs | 9107 | 399.3 us | 397.9 us | 402.3 us | 1.1 | - | 435.9 us | 446.0 us | 414.6 us | 403.5 us | 581.2 us | 221.5 us | -0.0 | 3.80-4.22 |
| tinygo-wasm | structserialize | 88213831 | 42 ns | 42 ns | 43 ns | 3.6 | - | 46 ns | 44 ns | 43 ns | 43 ns | 60 ns | 25 ns | -0.0 | 2.60-2.87 |
| tinygo-wasm/turbofan | count | 8351 | 667.6 us | 666.8 us | 670.1 us | 0.5 | - | 822.3 us | 744.8 us | 705.6 us | 685.9 us | 668.1 us | 662.7 us | -0.2 | 2.02-2.03 |
| tinygo-wasm/turbofan | fib | 15799 | 345.4 us | 343.9 us | 354.3 us | 3.0 | - | 430.1 us | 386.4 us | 364.0 us | 357.7 us | 341.0 us | 340.9 us | -0.2 | 1.57-1.67 |
| tinygo-wasm/turbofan | array | 333334 | 15.2 us | 15.2 us | 15.3 us | 0.5 | - | 19.0 us | 17.2 us | 16.1 us | 15.7 us | 15.4 us | 15.1 us | -0.2 | 2.04-2.24 |
| tinygo-wasm/turbofan | struct | 86348592 | 65 ns | 64 ns | 65 ns | 1.4 | - | 80 ns | 72 ns | 69 ns | 67 ns | 67 ns | 65 ns | -0.2 | 1.50-1.64 |
| tinygo-wasm/turbofan | memtest | 291 | 19.3 ms | 19.1 ms | 19.4 ms | 1.2 | - | 23.5 ms | 21.1 ms | 20.0 ms | 19.8 ms | 19.3 ms | 18.9 ms | -0.2 | 1.57-1.60 |
| tinygo-wasm/turbofan | primes | 86 | 64.0 ms | 63.7 ms | 64.6 ms | 1.5 | - | 80.0 ms | 70.9 ms | 66.7 ms | 65.9 ms | 63.2 ms | 62.4 ms | -0.2 | 1.79-1.96 |
| tinygo-wasm/turbofan | ackermann | 78566 | 71.6 us | 71.6 us | 78.2 us | 9.1 | - | 94.1 us | 79.0 us | 74.4 us | 75.8 us | 71.1 us | 70.2 us | -0.2 | 1.75-1.85 |
| tinygo-wasm/turbofan | digitalroot | 2921 | 1.9 ms | 1.9 ms | 1.9 ms | 3.0 | - | 2.4 ms | 2.1 ms | 2.0 ms | 1.9 ms | 1.9 ms | 1.9 ms | -0.2 | 1.92-2.01 |
| tinygo-wasm/turbofan | factorialiter | 53083622 | 105 ns | 104 ns | 106 ns | 1.6 | - | 128 ns | 117 ns | 111 ns | 108 ns | 104 ns | 105 ns | -0.2 | 1.66-2.01 |
| tinygo-wasm/turbofan | factorialrec | 45905215 | 123 ns | 123 ns | 125 ns | 1.6 | - | 148 ns | 138 ns | 128 ns | 127 ns | 120 ns | 120 ns | -0.2 | 1.55-1.86 |
| tinygo-wasm/turbofan | structs | 25406 | 217.6 us | 217.4 us | 219.1 us | 0.8 | - | 269.4 us | 245.7 us | 230.2 us | 224.1 us | 214.9 us | 218.1 us | -0.2 | 1.94-2.16 |
| tinygo-wasm/turbofan | structserialize | 222266634 | 25 ns | 25 ns | 25 ns | 1.0 | - | 31 ns | 29 ns | 26 ns | 26 ns | 25 ns | 25 ns | -0.2 | 1.98-2.14 |
| lua-wasm | count | 131 | 35.3 ms | 35.2 ms | 35.6 ms | 1.2 | - | 37.8 ms | 37.8 ms | 36.1 ms | 35.7 ms | 35.3 ms | 35.3 ms | -0.0 | 1.39-1.55 |
| lua-wasm | fib | 291 | 16.6 ms | 16.4 ms | 16.8 ms | 2.4 | - | 17.7 ms | 17.2 ms | 16.7 ms | 16.8 ms | 16.3 ms | 16.2 ms | -0.0 | 2.70-2.87 |
| lua-wasm | array | 4031 | 1.2 ms | 1.2 ms | 1.4 ms | 20.5 | - | 1.2 ms | 1.2 ms | 1.2 ms | 1.3 ms | 1.2 ms | 1.2 ms | -0.0 | 4.60-5.23 |
| lua-wasm | struct | 643006 | 7.1 us | 7.0 us | 7.1 us | 1.0 | - | 7.4 us | 7.2 us | 7.2 us | 7.1 us | 7.6 us | 7.6 us | -0.0 | 3.49-4.11 |
| lua-wasm | memtest | 4 | 1370.6 ms | 1354.2 ms | 1377.0 ms | 1.7 | - | - | 1445.9 ms | 1406.3 ms | 1376.8 ms | 1346.2 ms | 1378.7 ms | -0.5 | 2.27-2.82 |
| lua-wasm | primes | 6 | 875.9 ms | 871.9 ms | 877.1 ms | 0.6 | - | - | 959.8 ms | 925.2 ms | 881.7 ms | 848.2 ms | 875.3 ms | -1.0 | 1.94-2.13 |
| lua-wasm | ackermann | 1569 | 3.0 ms | 3.0 ms | 3.0 ms | 0.5 | - | 3.2 ms | 3.1 ms | 3.1 ms | 3.0 ms | 2.9 ms | 2.9 ms | -0.0 | 1.60-1.77 |
| lua-wasm | digitalroot | 46 | 100.7 ms | 99.4 ms | 101.1 ms | 1.7 | - | 106.1 ms | 106.2 ms | 100.6 ms | 101.2 ms | 100.7 ms | 99.0 ms | -0.0 | 2.31-2.83 |
| lua-wasm | factorialiter | 616792 | 7.8 us | 7.7 us | 7.9 us | 2.7 | - | 8.8 us | 8.1 us | 7.7 us | 7.9 us | 8.0 us | 8.0 us | -0.0 | 3.24-3.34 |
| lua-wasm | factorialrec | 244696 | 19.0 us | 18.7 us | 19.1 us | 2.3 | - | 20.6 us | 19.4 us | 19.5 us | 19.1 us | 18.8 us | 18.3 us | -0.0 | 2.21-2.40 |
| lua-wasm | structs | 447 | 10.7 ms | 10.4 ms | 11.3 ms | 8.4 | - | 11.9 ms | 10.9 ms | 10.9 ms | 10.9 ms | 10.4 ms | 10.3 ms | -0.2 | 2.44-2.70 |
| lua-wasm | structserialize | 406188 | 11.6 us | 11.5 us | 11.6 us | 1.0 | - | 12.7 us | 11.7 us | 11.7 us | 11.7 us | 11.5 us | 11.4 us | -0.0 | 1.62-1.68 |
| lua-wasm/turbofan | count | 278 | 19.7 ms | 19.5 ms | 20.5 ms | 5.1 | - | 28.9 ms | 22.9 ms | 21.3 ms | 20.8 ms | 19.9 ms | 19.6 ms | -0.2 | 1.52-1.99 |
| lua-wasm/turbofan | fib | 458 | 12.0 ms | 11.5 ms | 12.2 ms | 6.4 | - | 17.3 ms | 14.2 ms | 14.7 ms | 12.4 ms | 15.3 ms | 12.0 ms | -0.2 | 3.95-4.07 |
| lua-wasm/turbofan | array | 6432 | 839.1 us | 834.2 us | 851.7 us | 2.1 | - | 1.1 ms | 989.7 us | 927.2 us | 879.4 us | 856.0 us | 922.2 us | -0.3 | 2.60-2.90 |
| lua-wasm/turbofan | struct | 805061 | 6.8 us | 6.8 us | 6.9 us | 1.2 | - | 9.1 us | 8.1 us | 7.4 us | 7.1 us | 6.8 us | 6.9 us | -0.2 | 3.16-3.72 |
| lua-wasm/turbofan | memtest | 7 | 778.5 ms | 774.7 ms | 804.9 ms | 3.9 | - | - | 1534.4 ms | 1102.2 ms | 821.2 ms | 780.0 ms | 763.0 ms | -0.6 | 3.48-3.63 |
| lua-wasm/turbofan | primes | 9 | 595.4 ms | 585.1 ms | 602.4 ms | 2.9 | - | 872.4 ms | 737.4 ms | 646.4 ms | 622.2 ms | 568.8 ms | 587.2 ms | -1.1 | 2.95-3.41 |
| lua-wasm/turbofan | ackermann | 3063 | 1.8 ms | 1.8 ms | 1.8 ms | 2.4 | - | 2.5 ms | 2.1 ms | 2.0 ms | 1.9 ms | 1.8 ms | 1.8 ms | -0.2 | 2.20-2.35 |
| lua-wasm/turbofan | digitalroot | 83 | 71.2 ms | 66.6 ms | 71.6 ms | 7.0 | - | 93.8 ms | 83.4 ms | 77.4 ms | 72.7 ms | 71.9 ms | 71.3 ms | -0.2 | 1.50-1.67 |
| lua-wasm/turbofan | factorialiter | 821018 | 5.7 us | 5.6 us | 6.2 us | 9.7 | - | 8.6 us | 7.3 us | 6.7 us | 6.2 us | 6.5 us | 6.4 us | -0.2 | 1.68-2.17 |
| lua-wasm/turbofan | factorialrec | 420043 | 12.7 us | 12.5 us | 12.8 us | 2.7 | - | 17.7 us | 15.1 us | 13.9 us | 13.2 us | 12.5 us | 13.5 us | -0.2 | 2.01-2.24 |
| lua-wasm/turbofan | structs | 534 | 10.3 ms | 10.3 ms | 10.4 ms | 1.2 | - | 14.8 ms | 12.9 ms | 11.8 ms | 10.8 ms | 10.4 ms | 10.3 ms | -0.4 | 2.13-2.39 |
| lua-wasm/turbofan | structserialize | 652637 | 8.2 us | 8.1 us | 8.2 us | 1.5 | - | 11.2 us | 9.5 us | 9.0 us | 8.5 us | 8.0 us | 8.0 us | -0.2 | 2.72-3.09 |
| cx-llvm-wasm | count | 1073741824 | 1.36 ns est precomputed | 1.36 ns | 1.37 ns | 0.5 | - | 1.56 ns | 1.47 ns | 1.41 ns | 1.39 ns | - | - | - | 1.82-2.24 |
| cx-llvm-wasm | count | | size limit reached below the target time, the work is precomputed | | | | | | | | | | | | |
| cx-llvm-wasm | fib | 2989 | 1.1 ms est | 1.1 ms | 1.1 ms | 1.5 | - | 1.2 ms | 1.2 ms | 1.1 ms | 1.1 ms | - | - | - | 2.04-2.25 |
| cx-llvm-wasm | array | 155125 | 21.5 us est | 21.0 us | 21.7 us | 3.6 | - | 22.2 us | 21.7 us | 21.8 us | 21.6 us | - | - | - | 3.80-3.89 |
| cx-llvm-wasm | struct | 21524244 | 154 ns est | 154 ns | 156 ns | 1.2 | - | 167 ns | 156 ns | 156 ns | 156 ns | - | - | - | 3.59-3.73 |
| cx-llvm-wasm | memtest | 103 | 30.6 ms est | 30.4 ms | 31.7 ms | 4.4 | - | 35.0 ms | 31.3 ms | 31.2 ms | 31.2 ms | - | - | - | 3.24-3.32 |
| cx-llvm-wasm | primes | 43 | 76.6 ms est | 76.3 ms | 78.1 ms | 2.4 | - | 83.0 ms | 79.2 ms | 78.3 ms | 77.7 ms | - | - | - | 3.77-4.27 |
| cx-llvm-wasm | ackermann | 14725 | 219.3 us est | 218.6 us | 221.5 us | 1.3 | - | 234.1 us | 232.3 us | 225.0 us | 221.8 us | - | - | - | 2.97-3.44 |
| cx-llvm-wasm | digitalroot | 573 | 6.0 ms est | 6.0 ms | 6.0 ms | 0.3 | - | 6.6 ms | 6.2 ms | 6.1 ms | 6.1 ms | - | - | - | 3.95-4.50 |
| cx-llvm-wasm | factorialiter | 10000003 | 334 ns est | 328 ns | 334 ns | 2.0 | - | 362 ns | 350 ns | 334 ns | 335 ns | - | - | - | 3.13-3.73 |
| cx-llvm-wasm | factorialrec | 7544663 | 443 ns est | 442 ns | 444 ns | 0.3 | - | 490 ns | 473 ns | 456 ns | 447 ns | - | - | - | 2.92-3.27 |
| cx-llvm-wasm | structs | 7952 | 410.0 us est | 406.1 us | 431.5 us | 6.2 | - | 454.3 us | 434.8 us | 415.3 us | 419.6 us | - | - | - | 3.00-3.58 |
| cx-llvm-wasm | structserialize | 145881591 | 23 ns est | 23 ns | 23 ns | 2.2 | - | 25 ns | 25 ns | 23 ns | 23 ns | - | - | - | 2.50-2.83 |
| cx-llvm-wasm/turbofan | count | 1073741824 | 1.17 ns est precomputed | 1.17 ns | 1.17 ns | 0.8 | - | 1.70 ns | 1.49 ns | 1.31 ns | 1.23 ns | - | - | - | 2.20-2.34 |
| cx-llvm-wasm/turbofan | count | | size limit reached below the target time, the work is precomputed | | | | | | | | | | | | |
| cx-llvm-wasm/turbofan | fib | 5398 | 1.0 ms est | 1.0 ms | 1.0 ms | 1.7 | - | 1.1 ms | 1.1 ms | 1.0 ms | 1.0 ms | - | - | - | 2.23-2.70 |
| cx-llvm-wasm/turbofan | array | 333334 | 15.0 us est | 14.9 us | 15.2 us | 1.4 | - | 16.7 us | 16.2 us | 16.0 us | 15.3 us | - | - | - | 1.53-1.74 |
| cx-llvm-wasm/turbofan | struct | 71678894 | 81 ns est | 79 ns | 81 ns | 2.5 | - | 87 ns | 84 ns | 82 ns | 82 ns | - | - | - | 1.46-1.71 |
| cx-llvm-wasm/turbofan | memtest | 282 | 19.9 ms est | 19.9 ms | 20.0 ms | 0.5 | - | 21.9 ms | 21.0 ms | 20.6 ms | 20.2 ms | - | - | - | 1.89-2.17 |
| cx-llvm-wasm/turbofan | primes | 90 | 62.9 ms est | 62.4 ms | 63.4 ms | 1.6 | - | 69.4 ms | 66.0 ms | 65.8 ms | 63.6 ms | - | - | - | 2.77-3.00 |
| cx-llvm-wasm/turbofan | ackermann | 29027 | 196.3 us est | 195.9 us | 196.8 us | 0.5 | - | 213.6 us | 205.5 us | 202.4 us | 198.7 us | - | - | - | 2.34-2.77 |
| cx-llvm-wasm/turbofan | digitalroot | 2899 | 2.0 ms est | 2.0 ms | 2.0 ms | 0.5 | - | 2.2 ms | 2.1 ms | 2.0 ms | 2.0 ms | - | - | - | 2.13-2.27 |
| cx-llvm-wasm/turbofan | factorialiter | 41423945 | 136 ns est | 136 ns | 137 ns | 0.9 | - | 149 ns | 141 ns | 142 ns | 138 ns | - | - | - | 2.20-3.76 |
| cx-llvm-wasm/turbofan | factorialrec | 13595644 | 433 ns est | 427 ns | 485 ns | 13.3 | - | 468 ns | 440 ns | 429 ns | 453 ns | - | - | - | 3.80-4.34 |
| cx-llvm-wasm/turbofan | structs | 24464 | 229.7 us est | 229.6 us | 230.2 us | 0.3 | - | 253.2 us | 240.3 us | 236.3 us | 232.7 us | - | - | - | 2.91-3.30 |
| cx-llvm-wasm/turbofan | structserialize | 273115630 | 21 ns est | 21 ns | 21 ns | 1.3 | - | 23 ns | 22 ns | 21 ns | 21 ns | - | - | - | 1.88-2.10 |
| cx-wasm | count | 13 | 425.8 ms est | 417.0 ms | 432.0 ms | 3.5 | - | 671.3 ms | 493.7 ms | 444.8 ms | 444.1 ms | - | - | - | 1.68-1.88 |
| cx-wasm | fib | 41 | 135.5 ms est | 133.0 ms | 136.0 ms | 2.2 | - | 173.3 ms | 160.5 ms | 150.4 ms | 140.9 ms | - | - | - | 1.59-1.80 |
| cx-wasm | array | 334 | 15.0 ms est | 14.9 ms | 15.1 ms | 1.1 | - | 19.7 ms | 18.7 ms | 16.0 ms | 15.7 ms | - | - | - | 1.68-1.81 |
| cx-wasm | struct | 120004 | 49.3 us est | 44.1 us | 49.4 us | 10.8 | - | 60.5 us | 56.6 us | 51.9 us | 49.6 us | - | - | - | 2.02-2.44 |
| cx-wasm | memtest | 1 | 15747.7 ms est | 15498.1 ms | 17333.9 ms | 11.7 | - | - | - | - | 16442.5 ms | - | - | - | 2.12-3.35 |
| cx-wasm | primes | 1 | 4846.2 ms est | 4835.8 ms | 4915.1 ms | 1.6 | - | - | - | - | 5114.9 ms | - | - | - | 2.78-3.10 |
| cx-wasm | ackermann | 198 | 27.4 ms est | 27.0 ms | 27.5 ms | 1.6 | - | 36.0 ms | 31.8 ms | 29.7 ms | 28.6 ms | - | - | - | 2.25-2.39 |
| cx-wasm | digitalroot | 11 | 560.7 ms est | 560.5 ms | 574.3 ms | 2.5 | - | 772.9 ms | 691.1 ms | 608.9 ms | 587.8 ms | - | - | - | 1.97-2.28 |
| cx-wasm | factorialiter | 58130 | 91.5 us est | 90.3 us | 94.5 us | 4.7 | - | 117.5 us | 104.9 us | 96.1 us | 96.4 us | - | - | - | 2.06-2.16 |
| cx-wasm | factorialrec | 33334 | 157.9 us est | 156.7 us | 160.1 us | 2.1 | - | 202.6 us | 178.8 us | 168.1 us | 165.7 us | - | - | - | 1.94-2.11 |
| cx-wasm | structs | 78 | 73.8 ms est | 71.7 ms | 77.8 ms | 8.3 | - | 107.9 ms | 83.8 ms | 73.4 ms | 77.6 ms | - | - | - | 2.38-2.52 |
| cx-wasm | structserialize | 100003 | 49.9 us est | 47.7 us | 54.1 us | 12.9 | - | 69.2 us | 68.6 us | 60.7 us | 53.0 us | - | - | - | 2.66-3.39 |
| cx-wasm/turbofan | count | 7 | 308.1 ms est | 297.3 ms | 321.1 ms | 7.7 | - | - | 4037.6 ms | 1518.0 ms | 735.8 ms | - | - | - | 2.37-2.46 |
| cx-wasm/turbofan | fib | 17 | 129.9 ms est | 116.0 ms | 149.8 ms | 26.0 | - | 2049.3 ms | 1047.2 ms | 548.8 ms | 307.7 ms | - | - | - | 2.00-2.44 |
| cx-wasm/turbofan | array | 187 | 11.7 ms est | 11.3 ms | 11.9 ms | 4.6 | - | 181.0 ms | 91.4 ms | 48.3 ms | 27.6 ms | - | - | - | 2.90-3.79 |
| cx-wasm/turbofan | struct | 25106 | 87.5 us est | 87.0 us | 120.3 us | 38.1 | - | 1.3 ms | 713.6 us | 366.9 us | 217.3 us | - | - | - | 2.89-3.94 |
| cx-wasm/turbofan | memtest | 1 | 7684.4 ms est | 7436.8 ms | 8659.8 ms | 15.9 | - | - | - | - | 10916.0 ms | - | - | - | 2.18-3.62 |
| cx-wasm/turbofan | primes | 1 | 3892.5 ms est | 3810.4 ms | 4105.2 ms | 7.6 | - | - | - | - | 6925.0 ms | - | - | - | 2.03-2.70 |
| cx-wasm/turbofan | ackermann | 103 | 21.9 ms est | 21.6 ms | 31.1 ms | 43.1 | - | 340.8 ms | 169.4 ms | 88.8 ms | 53.9 ms | - | - | - | 2.38-2.51 |
| cx-wasm/turbofan | digitalroot | 5 | 426.3 ms est | 418.9 ms | 436.2 ms | 4.1 | - | - | 4142.8 ms | 2234.1 ms | 1025.0 ms | - | - | - | 2.16-2.33 |
| cx-wasm/turbofan | factorialiter | 27618 | 75.4 us est | 74.3 us | 89.1 us | 19.6 | - | 1.2 ms | 606.9 us | 323.8 us | 187.8 us | - | - | - | 2.07-2.50 |
| cx-wasm/turbofan | factorialrec | 15456 | 136.4 us est | 127.9 us | 139.3 us | 8.3 | - | 2.1 ms | 1.1 ms | 579.4 us | 328.0 us | - | - | - | 1.84-2.84 |
| cx-wasm/turbofan | structs | 16 | 143.4 ms est | 138.3 ms | 144.5 ms | 4.3 | - | 2108.9 ms | 1104.4 ms | 582.6 ms | 328.9 ms | - | - | - | 1.78-2.00 |
| cx-wasm/turbofan | structserialize | 21062 | 102.6 us est | 99.8 us | 105.3 us | 5.5 | - | 1.6 ms | 827.8 us | 437.9 us | 244.5 us | - | - | - | 2.96-3.13 |

## Notes

* go-wasm/turbofan: start up 1.325717869s, runs capped at 10s, under 100x start up
* tinygo-wasm/turbofan: start up 237.664331ms, runs capped at 10s, under 100x start up
* lua-wasm/turbofan: start up 477.823019ms, runs capped at 10s, under 100x start up
* cx-llvm-wasm/turbofan: start up 140.464005ms, runs capped at 10s, under 100x start up
* cx-wasm: start up 498.514746ms, runs capped at 10s, under 100x start up
* cx-wasm/turbofan: start up 5.977993498s, runs capped at 10s, under 100x start up

## Empty program

| form | layout | start to exit | start to S |
|---|---|---:|---:|
| c-wasm | a | 59.5 ms | 55.6 ms |
| c-wasm/turbofan | a | 106.7 ms | 100.4 ms |
| cx-llvm-wasm | a | 58.7 ms | - |
| cx-llvm-wasm/turbofan | a | 140.5 ms | - |
| cx-wasm | a | 498.5 ms | - |
| cx-wasm/turbofan | a | 5978.0 ms | - |
| go-wasm | a | 109.1 ms | 102.2 ms |
| go-wasm/turbofan | a | 1325.7 ms | 1250.7 ms |
| lua-wasm | a | 82.8 ms | 69.5 ms |
| lua-wasm/turbofan | a | 477.8 ms | 455.9 ms |
| tinygo-wasm | a | 65.6 ms | 61.1 ms |
| tinygo-wasm/turbofan | a | 237.7 ms | 225.6 ms |
