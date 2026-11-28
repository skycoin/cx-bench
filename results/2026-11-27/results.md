# cx-bench native results, 2026-11-27

Machine: an x86-64 desktop, 8 logical CPUs, 16 GB RAM, Linux 7.2, with the user's desktop session running beside the benchmarks (the one minute load average of each run is in the tables, about 1.5 to 4.7). Tools: clang 23.1.4, Go 1.27.3, TinyGo 0.44.0-dev (LLVM 23.1.3), Lua 5.4.9 and 5.5.1, LuaJIT 2.1.1792588164, gopher-lua 1.1.2, Python 3.14.8, skycoin-cx 01deb81af6fc.

Harness: linux/amd64, 8 logical CPUs, started 2026-11-27T07:38:19Z.

Harness: linux/amd64, 8 logical CPUs, started 2026-11-27T08:47:16Z.

Time per round is the time of the two passes over two times the size, from the phase markers. A form without markers (the CX forms) has it from the whole run less its empty program, marked est. A difference to C counts (sig yes) only when it is larger than the run to run spread and the placement delta of both forms. A form under a tenth of C on a benchmark is marked precomputed: its compiler worked out the result at build time and removed the work, and the best of table leaves it out.

![best of chart](chart.png)

## Best of each language, per benchmark

| benchmark | best CX | x C | best Lua | x C | best Go | x C | C | sig CX vs C |
|---|---|---:|---|---:|---|---:|---:|---|
| count | cx-image 33.8 ms est | 75.67 | luajit 1.8 ms | 3.93 | tinygo 425.6 us | 0.95 | 446.1 us | yes |
| fib | cx-llvm 382.4 us est | 1.19 | luajit 872.9 us | 2.72 | tinygo 307.9 us | 0.96 | 321.2 us | yes |
| array | cx-llvm 15.3 us est | 1.00 | luajit 15.9 us | 1.04 | tinygo 15.7 us | 1.02 | 15.3 us | no |
| struct | cx-llvm 25 ns est | 1.08 | luajit 339 ns | 14.56 | tinygo 24 ns | 1.02 | 23 ns | yes |
| memtest | cx-llvm 9.3 ms est | 1.04 | luajit 32.7 ms | 3.66 | tinygo 9.8 ms | 1.10 | 8.9 ms | no |
| primes | cx-llvm 57.4 ms est | 1.00 | luajit 110.6 ms | 1.94 | tinygo 58.5 ms | 1.02 | 57.1 ms | no |
| ackermann | cx-llvm 86.5 us est | 1.72 | luajit 133.0 us | 2.65 | tinygo 46.3 us | 0.92 | 50.3 us | yes |
| digitalroot | cx-llvm 1.2 ms est | 0.91 | luajit 10.0 ms | 7.33 | tinygo 1.1 ms | 0.83 | 1.4 ms | no |
| factorialiter | cx-llvm 97 ns est | 1.00 | luajit 1.6 us | 16.53 | tinygo 121 ns | 1.25 | 97 ns | no |
| factorialrec | cx-llvm 290 ns est | 2.82 | luajit 2.6 us | 25.22 | tinygo 103 ns | 1.00 | 103 ns | yes |
| structs | cx-llvm 167.6 us est | 1.18 | luajit 347.8 us | 2.44 | tinygo 174.0 us | 1.22 | 142.3 us | yes |
| structserialize | cx-llvm 15 ns est | 4.75 | luajit 336 ns | 105.15 | tinygo 4.96 ns | 1.55 | 3.20 ns | yes |

## Phases of the best forms

Load is process start to S, Loop 1 and Loop 2 the two identical passes, Finish from 2 to exit. Percent of the whole run.

| benchmark | language | form | size | Load | Loop 1 | Loop 2 | Finish | Load % | Loop 1 % | Loop 2 % | Finish % | spread % | placement % | load avg |
|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|
| count | CX | cx-image | 35 | 2.1 ms est | 1176.0 ms | 1176.0 ms | in Load | 0.1 | 50.0 | 50.0 | - | 5.5 | 0.9 | 2.14-2.35 |
| count | Lua | luajit | 662 | 1.3 ms | 1166.9 ms | 1163.2 ms | 411.0 us | 0.1 | 50.0 | 49.9 | 0.0 | 1.9 | - | 1.48-1.57 |
| count | Go | tinygo | 2599 | 480.7 us | 1145.9 ms | 1107.4 ms | 206.0 us | 0.0 | 50.8 | 49.1 | 0.0 | 7.1 | 0.4 | 1.93-2.26 |
| count | C | c | 2548 | 661.6 us | 1139.1 ms | 1127.0 ms | 262.6 us | 0.0 | 50.2 | 49.7 | 0.0 | 1.9 | 2.1 | 2.28-2.33 |
| fib | CX | cx-llvm | 2980 | 667.4 us est | 1146.9 ms | 1146.9 ms | in Load | 0.0 | 50.0 | 50.0 | - | 4.0 | 1.6 | 2.54-2.82 |
| fib | Lua | luajit | 1290 | 1.4 ms | 1120.4 ms | 1131.2 ms | 444.1 us | 0.1 | 49.7 | 50.2 | 0.0 | 1.4 | - | 1.71-1.78 |
| fib | Go | tinygo | 3687 | 456.9 us | 1134.9 ms | 1134.1 ms | 228.3 us | 0.0 | 50.0 | 50.0 | 0.0 | 1.5 | 0.0 | 1.87-2.03 |
| fib | C | c | 3570 | 831.4 us | 1181.3 ms | 1168.1 ms | 226.9 us | 0.0 | 50.3 | 49.7 | 0.0 | 15.6 | 2.2 | 2.37-3.09 |
| array | CX | cx-llvm | 72532 | 667.4 us est | 1110.4 ms | 1110.4 ms | in Load | 0.0 | 50.0 | 50.0 | - | 1.5 | 0.1 | 2.22-2.34 |
| array | Lua | luajit | 65785 | 1.1 ms | 1036.5 ms | 1056.1 ms | 444.7 us | 0.1 | 49.5 | 50.4 | 0.0 | 2.3 | - | 1.58-1.68 |
| array | Go | tinygo | 67824 | 457.0 us | 1051.1 ms | 1070.2 ms | 195.0 us | 0.0 | 49.5 | 50.4 | 0.0 | 2.9 | 1.0 | 2.03-2.35 |
| array | C | c | 66370 | 721.5 us | 1024.1 ms | 1008.0 ms | 282.7 us | 0.0 | 50.4 | 49.6 | 0.0 | 2.6 | 0.1 | 2.96-2.96 |
| struct | CX | cx-llvm | 46199270 | 667.4 us est | 1150.8 ms | 1150.8 ms | in Load | 0.0 | 50.0 | 50.0 | - | 3.2 | 5.1 | 2.01-2.09 |
| struct | Lua | luajit | 3404213 | 1.1 ms | 1133.8 ms | 1149.3 ms | 433.8 us | 0.0 | 49.6 | 50.3 | 0.0 | 3.3 | - | 1.41-1.49 |
| struct | Go | tinygo | 48524515 | 492.1 us | 1166.1 ms | 1148.0 ms | 203.8 us | 0.0 | 50.4 | 49.6 | 0.0 | 1.8 | 0.3 | 1.86-1.93 |
| struct | C | c | 52202742 | 556.5 us | 1209.8 ms | 1201.7 ms | 279.0 us | 0.0 | 50.2 | 49.8 | 0.0 | 3.3 | 3.0 | 2.83-3.01 |
| memtest | CX | cx-llvm | 145 | 667.4 us est | 1346.0 ms | 1346.0 ms | in Load | 0.0 | 50.0 | 50.0 | - | 2.9 | 4.8 | 1.90-2.06 |
| memtest | Lua | luajit | 32 | 1.2 ms | 1055.7 ms | 1041.2 ms | 1.1 ms | 0.1 | 50.3 | 49.6 | 0.1 | 1.5 | - | 1.37-1.40 |
| memtest | Go | tinygo | 105 | 450.2 us | 1037.2 ms | 1044.8 ms | 570.2 us | 0.0 | 49.8 | 50.2 | 0.0 | 6.1 | 0.1 | 1.72-1.79 |
| memtest | C | c | 131 | 719.6 us | 1182.0 ms | 1163.6 ms | 377.8 us | 0.0 | 50.4 | 49.6 | 0.0 | 4.1 | 2.4 | 2.94-3.35 |
| primes | CX | cx-llvm | 17 | 667.4 us est | 987.7 ms | 987.7 ms | in Load | 0.0 | 50.0 | 50.0 | - | 9.7 | 0.4 | 3.04-3.33 |
| primes | Lua | luajit | 9 | 1.0 ms | 1016.3 ms | 982.5 ms | 15.4 ms | 0.1 | 50.4 | 48.8 | 0.8 | 2.7 | - | 1.59-1.64 |
| primes | Go | tinygo | 20 | 459.9 us | 1159.4 ms | 1156.1 ms | 797.0 us | 0.0 | 50.0 | 49.9 | 0.0 | 3.6 | 1.5 | 1.61-1.66 |
| primes | C | c | 19 | 558.0 us | 1086.8 ms | 1084.9 ms | 892.9 us | 0.0 | 50.0 | 49.9 | 0.0 | 1.8 | 1.4 | 3.09-3.17 |
| ackermann | CX | cx-llvm | 11579 | 667.4 us est | 1007.3 ms | 1007.3 ms | in Load | 0.0 | 50.0 | 50.0 | - | 4.4 | 1.1 | 3.01-3.11 |
| ackermann | Lua | luajit | 9028 | 1.3 ms | 1208.3 ms | 1209.1 ms | 412.8 us | 0.1 | 49.9 | 50.0 | 0.0 | 6.3 | - | 1.73-1.86 |
| ackermann | Go | tinygo | 23871 | 432.6 us | 1099.5 ms | 1103.7 ms | 200.2 us | 0.0 | 49.9 | 50.1 | 0.0 | 1.9 | 0.1 | 1.51-1.71 |
| ackermann | C | c | 23108 | 1.0 ms | 1196.2 ms | 1219.6 ms | 292.4 us | 0.0 | 49.5 | 50.5 | 0.0 | 30.5 | 7.4 | 3.20-3.24 |
| digitalroot | CX | cx-llvm | 952 | 667.4 us est | 1180.9 ms | 1180.9 ms | in Load | 0.0 | 50.0 | 50.0 | - | 2.8 | 0.1 | 3.61-3.88 |
| digitalroot | Lua | luajit | 137 | 1.5 ms | 1376.9 ms | 1383.4 ms | 413.6 us | 0.1 | 49.8 | 50.1 | 0.0 | 4.3 | - | 1.52-1.57 |
| digitalroot | Go | tinygo | 1007 | 477.9 us | 1154.1 ms | 1134.5 ms | 210.6 us | 0.0 | 50.4 | 49.6 | 0.0 | 2.1 | 0.7 | 1.88-1.89 |
| digitalroot | C | c | 825 | 563.4 us | 1128.4 ms | 1122.7 ms | 276.9 us | 0.0 | 50.1 | 49.9 | 0.0 | 1.6 | 13.0 | 3.20-3.60 |
| factorialiter | CX | cx-llvm | 11660165 | 667.4 us est | 1131.9 ms | 1131.9 ms | in Load | 0.0 | 50.0 | 50.0 | - | 3.9 | 0.1 | 3.19-3.41 |
| factorialiter | Lua | luajit | 734573 | 1.2 ms | 1183.2 ms | 1188.3 ms | 415.5 us | 0.1 | 49.9 | 50.1 | 0.0 | 4.0 | - | 1.43-1.52 |
| factorialiter | Go | tinygo | 9196078 | 446.5 us | 1112.5 ms | 1112.7 ms | 193.9 us | 0.0 | 50.0 | 50.0 | 0.0 | 1.3 | 0.4 | 2.03-2.04 |
| factorialiter | C | c | 11905544 | 659.2 us | 1153.0 ms | 1158.0 ms | 261.2 us | 0.0 | 49.9 | 50.1 | 0.0 | 2.9 | 2.1 | 2.56-2.84 |
| factorialrec | CX | cx-llvm | 3932777 | 667.4 us est | 1156.5 ms | 1156.5 ms | in Load | 0.0 | 50.0 | 50.0 | - | 6.7 | 0.3 | 2.86-2.95 |
| factorialrec | Lua | luajit | 538201 | 1.3 ms | 1407.4 ms | 1391.7 ms | 427.4 us | 0.0 | 50.3 | 49.7 | 0.0 | 1.7 | - | 1.41-1.57 |
| factorialrec | Go | tinygo | 10839104 | 511.1 us | 1114.1 ms | 1122.3 ms | 200.0 us | 0.0 | 49.8 | 50.2 | 0.0 | 1.2 | 0.8 | 2.31-2.61 |
| factorialrec | C | c | 10852096 | 619.4 us | 1115.5 ms | 1106.6 ms | 273.7 us | 0.0 | 50.2 | 49.8 | 0.0 | 3.3 | 1.8 | 2.36-2.53 |
| structs | CX | cx-llvm | 6918 | 667.4 us est | 1161.8 ms | 1161.8 ms | in Load | 0.0 | 50.0 | 50.0 | - | 1.0 | 0.5 | 3.76-4.32 |
| structs | Lua | luajit | 3319 | 1.3 ms | 1152.4 ms | 1153.7 ms | 571.2 us | 0.1 | 49.9 | 50.0 | 0.0 | 3.9 | - | 1.64-1.76 |
| structs | Go | tinygo | 6628 | 462.6 us | 1155.0 ms | 1147.0 ms | 275.4 us | 0.0 | 50.2 | 49.8 | 0.0 | 2.0 | 7.3 | 2.66-2.72 |
| structs | C | c | 7937 | 653.8 us | 1125.2 ms | 1123.5 ms | 338.9 us | 0.0 | 50.0 | 49.9 | 0.0 | 2.4 | 1.8 | 2.43-2.47 |
| structserialize | CX | cx-llvm | 74224869 | 667.4 us est | 1129.6 ms | 1129.6 ms | in Load | 0.0 | 50.0 | 50.0 | - | 1.4 | 0.7 | 2.54-2.72 |
| structserialize | Lua | luajit | 3347232 | 1.2 ms | 1124.9 ms | 1112.9 ms | 454.4 us | 0.1 | 50.2 | 49.7 | 0.0 | 5.0 | - | 1.69-1.79 |
| structserialize | Go | tinygo | 235338488 | 490.7 us | 1175.4 ms | 1155.6 ms | 183.5 us | 0.0 | 50.4 | 49.6 | 0.0 | 1.6 | 0.2 | 1.83-1.92 |
| structserialize | C | c | 360613129 | 616.1 us | 1159.1 ms | 1155.8 ms | 271.1 us | 0.0 | 50.1 | 49.9 | 0.0 | 2.9 | 0.6 | 2.08-2.19 |

## Every form

Sweep is the whole run, start up included, per round at size/8, /4, /2 and the size: flat means the start up is amortized. Server is the time per round the program measured itself in server mode, first and third run of one process. Self is how far the program's own -time=us loop times are from the harness markers.

| form | benchmark | size | per round | min | max | spread % | placement % | sweep /8 | /4 | /2 | x1 | server 1st | 3rd | self % | load avg |
|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|
| c | count | 2548 | 446.1 us | 440.3 us | 448.8 us | 1.9 | 2.1 | 440.6 us | 448.3 us | 443.5 us | 444.9 us | 438.2 us | 442.3 us | +0.0 | 2.28-2.33 |
| c | fib | 3570 | 321.2 us | 315.4 us | 365.5 us | 15.6 | 2.2 | 384.9 us | 460.1 us | 519.1 us | 329.2 us | 324.3 us | 314.6 us | -0.0 | 2.37-3.09 |
| c | array | 66370 | 15.3 us | 15.1 us | 15.5 us | 2.6 | 0.1 | 15.5 us | 15.5 us | 15.2 us | 15.3 us | 15.5 us | 15.4 us | -0.0 | 2.96-2.96 |
| c | struct | 52202742 | 23 ns | 23 ns | 23 ns | 3.3 | 3.0 | 23 ns | 23 ns | 23 ns | 23 ns | 23 ns | 23 ns | -0.0 | 2.83-3.01 |
| c | memtest | 131 | 8.9 ms | 8.8 ms | 9.1 ms | 4.1 | 2.4 | 9.6 ms | 9.5 ms | 9.5 ms | 9.0 ms | 9.1 ms | 9.1 ms | -0.0 | 2.94-3.35 |
| c | primes | 19 | 57.1 ms | 56.7 ms | 57.7 ms | 1.8 | 1.4 | 59.8 ms | 58.7 ms | 57.8 ms | 57.2 ms | 58.4 ms | 58.2 ms | -0.0 | 3.09-3.17 |
| c | ackermann | 23108 | 50.3 us | 48.1 us | 63.4 us | 30.5 | 7.4 | 48.7 us | 46.6 us | 48.5 us | 52.3 us | 46.2 us | 46.4 us | -0.0 | 3.20-3.24 |
| c | digitalroot | 825 | 1.4 ms | 1.4 ms | 1.4 ms | 1.6 | 13.0 | 1.4 ms | 1.3 ms | 1.4 ms | 1.4 ms | 1.3 ms | 1.4 ms | -0.0 | 3.20-3.60 |
| c | factorialiter | 11905544 | 97 ns | 96 ns | 98 ns | 2.9 | 2.1 | 97 ns | 107 ns | 96 ns | 97 ns | 95 ns | 93 ns | -0.0 | 2.56-2.84 |
| c | factorialrec | 10852096 | 103 ns | 100 ns | 104 ns | 3.3 | 1.8 | 99 ns | 101 ns | 104 ns | 102 ns | 101 ns | 101 ns | -0.0 | 2.36-2.53 |
| c | structs | 7937 | 142.3 us | 139.7 us | 143.1 us | 2.4 | 1.8 | 144.3 us | 141.1 us | 139.6 us | 141.7 us | 141.1 us | 141.3 us | -0.0 | 2.43-2.47 |
| c | structserialize | 360613129 | 3.20 ns | 3.17 ns | 3.27 ns | 2.9 | 0.6 | 3.33 ns | 3.30 ns | 3.20 ns | 3.21 ns | 3.19 ns | 3.18 ns | -0.0 | 2.08-2.19 |
| go | count | 1782 | 657.5 us | 650.5 us | 685.3 us | 5.3 | 0.1 | 691.8 us | 681.7 us | 654.0 us | 663.7 us | 661.6 us | 659.3 us | -0.0 | 2.17-2.30 |
| go | fib | 2586 | 492.8 us | 489.2 us | 505.2 us | 3.3 | 1.2 | 446.3 us | 469.0 us | 543.8 us | 496.6 us | 512.2 us | 507.1 us | -0.0 | 1.85-1.92 |
| go | array | 71178 | 16.4 us | 16.3 us | 16.9 us | 3.7 | 0.0 | 16.8 us | 17.2 us | 16.5 us | 16.5 us | 17.6 us | 16.7 us | -0.0 | 1.60-1.65 |
| go | struct | 12533634 | 90 ns | 88 ns | 91 ns | 3.1 | 0.1 | 92 ns | 89 ns | 93 ns | 89 ns | 91 ns | 90 ns | -0.0 | 1.54-1.93 |
| go | memtest | 44 | 25.9 ms | 25.3 ms | 26.1 ms | 3.2 | 0.2 | 25.9 ms | 26.1 ms | 25.4 ms | 25.8 ms | 26.2 ms | 25.9 ms | -0.0 | 2.06-2.15 |
| go | primes | 17 | 71.8 ms | 70.2 ms | 76.6 ms | 8.8 | 4.1 | 74.9 ms | 72.7 ms | 72.4 ms | 72.5 ms | 71.7 ms | 72.4 ms | -0.0 | 1.99-2.08 |
| go | ackermann | 6307 | 232.2 us | 227.6 us | 244.7 us | 7.4 | 3.3 | 237.7 us | 220.0 us | 239.4 us | 233.2 us | 237.3 us | 226.8 us | -0.0 | 1.86-1.94 |
| go | digitalroot | 877 | 1.5 ms | 1.5 ms | 1.5 ms | 3.4 | 9.1 | 1.5 ms | 1.5 ms | 1.5 ms | 1.5 ms | 1.4 ms | 1.5 ms | -0.0 | 2.04-2.19 |
| go | factorialiter | 5659898 | 201 ns | 201 ns | 204 ns | 1.6 | 0.1 | 214 ns | 205 ns | 203 ns | 202 ns | 204 ns | 201 ns | -0.0 | 1.94-2.11 |
| go | factorialrec | 3307465 | 314 ns | 295 ns | 331 ns | 11.6 | 3.9 | 413 ns | 328 ns | 309 ns | 314 ns | 314 ns | 309 ns | -0.0 | 1.88-2.04 |
| go | structs | 4978 | 235.0 us | 230.4 us | 237.6 us | 3.0 | 7.4 | 229.9 us | 232.1 us | 245.2 us | 234.7 us | 232.6 us | 241.2 us | -0.0 | 1.59-1.69 |
| go | structserialize | 11266356 | 103 ns | 101 ns | 105 ns | 3.3 | 1.0 | 101 ns | 102 ns | 105 ns | 103 ns | 102 ns | 100 ns | -0.0 | 1.72-1.90 |
| tinygo | count | 2599 | 425.6 us | 420.4 us | 450.8 us | 7.1 | 0.4 | 421.0 us | 441.3 us | 437.4 us | 433.6 us | 426.3 us | 429.8 us | -0.0 | 1.93-2.26 |
| tinygo | fib | 3687 | 307.9 us | 305.2 us | 310.0 us | 1.5 | 0.0 | 301.2 us | 310.5 us | 306.7 us | 307.8 us | 303.7 us | 310.6 us | -0.0 | 1.87-2.03 |
| tinygo | array | 67824 | 15.7 us | 15.4 us | 15.8 us | 2.9 | 1.0 | 15.4 us | 15.4 us | 15.8 us | 15.6 us | 15.2 us | 15.3 us | -0.0 | 2.03-2.35 |
| tinygo | struct | 48524515 | 24 ns | 24 ns | 24 ns | 1.8 | 0.3 | 24 ns | 24 ns | 24 ns | 24 ns | 24 ns | 24 ns | -0.0 | 1.86-1.93 |
| tinygo | memtest | 105 | 9.8 ms | 9.7 ms | 10.3 ms | 6.1 | 0.1 | 10.1 ms | 9.6 ms | 9.9 ms | 9.9 ms | 9.6 ms | 9.8 ms | -0.0 | 1.72-1.79 |
| tinygo | primes | 20 | 58.5 ms | 56.6 ms | 58.7 ms | 3.6 | 1.5 | 62.9 ms | 59.3 ms | 57.3 ms | 57.9 ms | 57.5 ms | 56.0 ms | +0.0 | 1.61-1.66 |
| tinygo | ackermann | 23871 | 46.3 us | 45.7 us | 46.6 us | 1.9 | 0.1 | 49.0 us | 47.3 us | 45.8 us | 46.2 us | 46.2 us | 46.3 us | -0.0 | 1.51-1.71 |
| tinygo | digitalroot | 1007 | 1.1 ms | 1.1 ms | 1.2 ms | 2.1 | 0.7 | 1.1 ms | 1.1 ms | 1.1 ms | 1.1 ms | 1.1 ms | 1.1 ms | -0.0 | 1.88-1.89 |
| tinygo | factorialiter | 9196078 | 121 ns | 120 ns | 122 ns | 1.3 | 0.4 | 126 ns | 123 ns | 120 ns | 121 ns | 121 ns | 123 ns | -0.0 | 2.03-2.04 |
| tinygo | factorialrec | 10839104 | 103 ns | 102 ns | 104 ns | 1.2 | 0.8 | 105 ns | 109 ns | 107 ns | 103 ns | 104 ns | 104 ns | -0.0 | 2.31-2.61 |
| tinygo | structs | 6628 | 174.0 us | 171.9 us | 175.3 us | 2.0 | 7.3 | 173.7 us | 176.7 us | 174.0 us | 173.7 us | 171.9 us | 171.8 us | -0.0 | 2.66-2.72 |
| tinygo | structserialize | 235338488 | 4.96 ns | 4.91 ns | 4.99 ns | 1.6 | 0.2 | 4.85 ns | 4.93 ns | 4.95 ns | 4.95 ns | 4.87 ns | 5.15 ns | -0.0 | 1.83-1.92 |
| lua/5.4 | count | 80 | 13.2 ms | 12.9 ms | 13.4 ms | 3.8 | - | 13.0 ms | 13.1 ms | 13.8 ms | 13.2 ms | 13.0 ms | 12.9 ms | -0.4 | 2.66-2.71 |
| lua/5.4 | fib | 159 | 7.4 ms | 7.2 ms | 7.4 ms | 2.9 | - | 8.0 ms | 7.6 ms | 7.5 ms | 7.4 ms | 7.5 ms | 7.2 ms | -0.6 | 2.41-2.58 |
| lua/5.4 | array | 2839 | 405.2 us | 402.4 us | 407.7 us | 1.3 | - | 401.8 us | 418.4 us | 419.6 us | 405.7 us | 397.1 us | 394.4 us | -1.4 | 2.37-2.54 |
| lua/5.4 | struct | 240145 | 4.2 us | 4.1 us | 4.2 us | 3.3 | - | 4.2 us | 4.1 us | 4.2 us | 4.2 us | 4.2 us | 4.1 us | -1.3 | 2.05-2.24 |
| lua/5.4 | memtest | 3 | 489.6 ms | 485.3 ms | 505.2 ms | 4.1 | - | - | - | 492.2 ms | 494.7 ms | 479.2 ms | 470.6 ms | -1.6 | 1.75-1.89 |
| lua/5.4 | primes | 3 | 437.6 ms | 431.8 ms | 495.9 ms | 14.7 | - | - | - | 494.3 ms | 458.2 ms | 412.8 ms | 426.5 ms | -2.3 | 1.81-1.84 |
| lua/5.4 | ackermann | 1327 | 1.1 ms | 1.1 ms | 1.1 ms | 1.3 | - | 1.1 ms | 1.2 ms | 1.1 ms | 1.1 ms | 1.1 ms | 1.1 ms | -1.1 | 1.55-1.65 |
| lua/5.4 | digitalroot | 22 | 47.5 ms | 44.2 ms | 53.1 ms | 18.7 | - | 50.6 ms | 46.3 ms | 49.5 ms | 48.2 ms | 47.2 ms | 48.4 ms | -0.5 | 1.63-1.68 |
| lua/5.4 | factorialiter | 262394 | 4.5 us | 4.1 us | 5.6 us | 34.1 | - | 4.2 us | 4.2 us | 4.3 us | 4.6 us | 4.4 us | 4.3 us | -0.4 | 2.10-2.21 |
| lua/5.4 | factorialrec | 121138 | 8.8 us | 8.5 us | 9.0 us | 6.2 | - | 8.6 us | 9.0 us | 8.9 us | 8.8 us | 8.8 us | 8.4 us | -0.5 | 2.03-2.22 |
| lua/5.4 | structs | 204 | 5.5 ms | 5.5 ms | 5.5 ms | 1.2 | - | 6.8 ms | 5.8 ms | 5.7 ms | 5.5 ms | 5.7 ms | 5.5 ms | -2.7 | 1.95-2.11 |
| lua/5.4 | structserialize | 268313 | 5.4 us | 5.2 us | 5.5 us | 5.6 | - | 5.6 us | 5.1 us | 5.5 us | 5.3 us | 5.5 us | 5.3 us | -0.5 | 1.54-1.64 |
| lua/5.5 | count | 82 | 13.1 ms | 13.0 ms | 13.2 ms | 1.7 | - | 13.8 ms | 13.3 ms | 13.0 ms | 13.1 ms | 13.5 ms | 13.2 ms | -0.5 | 1.95-2.13 |
| lua/5.5 | fib | 206 | 7.3 ms | 7.3 ms | 7.4 ms | 2.1 | - | 7.2 ms | 7.5 ms | 7.6 ms | 7.4 ms | 7.3 ms | 7.2 ms | -0.4 | 1.63-1.74 |
| lua/5.5 | array | 2854 | 412.5 us | 405.9 us | 415.2 us | 2.2 | - | 403.4 us | 405.2 us | 410.1 us | 411.4 us | 400.9 us | 409.6 us | -0.4 | 1.54-1.64 |
| lua/5.5 | struct | 249965 | 4.5 us | 4.2 us | 4.6 us | 7.6 | - | 4.7 us | 4.2 us | 4.2 us | 4.5 us | 4.6 us | 4.5 us | -0.4 | 1.50-1.58 |
| lua/5.5 | memtest | 3 | 513.2 ms | 508.9 ms | 515.4 ms | 1.3 | - | - | - | 527.0 ms | 512.9 ms | 502.4 ms | 519.8 ms | -1.7 | 1.70-1.80 |
| lua/5.5 | primes | 3 | 392.0 ms | 384.0 ms | 399.7 ms | 4.0 | - | - | - | 441.7 ms | 391.5 ms | 382.3 ms | 380.9 ms | -2.1 | 1.46-1.54 |
| lua/5.5 | ackermann | 1118 | 1.1 ms | 1.1 ms | 1.2 ms | 6.9 | - | 1.1 ms | 1.1 ms | 1.1 ms | 1.1 ms | 1.1 ms | 1.1 ms | -0.7 | 1.45-1.54 |
| lua/5.5 | digitalroot | 26 | 45.1 ms | 42.6 ms | 50.2 ms | 16.8 | - | 47.9 ms | 47.1 ms | 47.4 ms | 45.6 ms | 51.0 ms | 49.6 ms | -0.7 | 1.55-1.65 |
| lua/5.5 | factorialiter | 341845 | 4.1 us | 3.9 us | 4.4 us | 12.9 | - | 4.1 us | 4.2 us | 4.0 us | 4.1 us | 4.4 us | 4.1 us | -0.4 | 1.43-1.51 |
| lua/5.5 | factorialrec | 124934 | 8.7 us | 8.5 us | 8.9 us | 3.8 | - | 8.2 us | 9.6 us | 8.2 us | 8.7 us | 8.6 us | 8.9 us | -0.5 | 1.54-1.58 |
| lua/5.5 | structs | 207 | 5.5 ms | 5.5 ms | 5.6 ms | 2.5 | - | 5.6 ms | 5.6 ms | 5.5 ms | 5.5 ms | 5.4 ms | 5.4 ms | -0.7 | 1.65-1.71 |
| lua/5.5 | structserialize | 281576 | 4.8 us | 4.7 us | 5.2 us | 9.0 | - | 4.8 us | 5.0 us | 5.1 us | 4.9 us | 5.0 us | 5.1 us | -0.4 | 1.69-1.87 |
| luajit | count | 662 | 1.8 ms | 1.8 ms | 1.8 ms | 1.9 | - | 1.7 ms | 1.7 ms | 1.8 ms | 1.8 ms | 1.7 ms | 1.7 ms | -0.0 | 1.48-1.57 |
| luajit | fib | 1290 | 872.9 us | 865.3 us | 877.3 us | 1.4 | - | 895.4 us | 893.5 us | 879.5 us | 873.4 us | 880.5 us | 903.7 us | -0.0 | 1.71-1.78 |
| luajit | array | 65785 | 15.9 us | 15.7 us | 16.1 us | 2.3 | - | 15.9 us | 16.2 us | 16.3 us | 15.9 us | 15.5 us | 15.5 us | -0.0 | 1.58-1.68 |
| luajit | struct | 3404213 | 339 ns | 328 ns | 340 ns | 3.3 | - | 370 ns | 323 ns | 326 ns | 336 ns | 338 ns | 331 ns | -0.0 | 1.41-1.49 |
| luajit | memtest | 32 | 32.7 ms | 32.6 ms | 33.1 ms | 1.5 | - | 34.1 ms | 33.2 ms | 33.5 ms | 32.8 ms | 35.4 ms | 34.0 ms | -0.3 | 1.37-1.40 |
| luajit | primes | 9 | 110.6 ms | 110.0 ms | 113.0 ms | 2.7 | - | 130.2 ms | 121.5 ms | 112.0 ms | 112.0 ms | 113.8 ms | 122.0 ms | -0.6 | 1.59-1.64 |
| luajit | ackermann | 9028 | 133.0 us | 130.6 us | 138.9 us | 6.3 | - | 137.5 us | 132.1 us | 128.0 us | 134.0 us | 131.9 us | 129.3 us | -0.0 | 1.73-1.86 |
| luajit | digitalroot | 137 | 10.0 ms | 9.8 ms | 10.3 ms | 4.3 | - | 10.0 ms | 9.9 ms | 10.1 ms | 10.1 ms | 9.9 ms | 9.9 ms | -0.0 | 1.52-1.57 |
| luajit | factorialiter | 734573 | 1.6 us | 1.6 us | 1.6 us | 4.0 | - | 1.6 us | 1.7 us | 1.7 us | 1.6 us | 1.6 us | 1.6 us | -0.0 | 1.43-1.52 |
| luajit | factorialrec | 538201 | 2.6 us | 2.6 us | 2.6 us | 1.7 | - | 2.6 us | 2.6 us | 2.7 us | 2.6 us | 2.7 us | 2.8 us | -0.0 | 1.41-1.57 |
| luajit | structs | 3319 | 347.8 us | 342.2 us | 355.8 us | 3.9 | - | 385.4 us | 343.0 us | 345.3 us | 347.7 us | 356.7 us | 347.1 us | -0.1 | 1.64-1.76 |
| luajit | structserialize | 3347232 | 336 ns | 324 ns | 341 ns | 5.0 | - | 333 ns | 325 ns | 329 ns | 335 ns | 328 ns | 322 ns | -0.0 | 1.69-1.79 |
| gopher-lua | count | 16 | 162.6 ms | 132.8 ms | 194.2 ms | 37.8 | - | 200.2 ms | 218.6 ms | 172.7 ms | 165.3 ms | 154.4 ms | 196.5 ms | -0.7 | 2.37-2.67 |
| gopher-lua | fib | 92 | 40.4 ms | 38.8 ms | 41.5 ms | 6.9 | - | 42.1 ms | 42.4 ms | 40.3 ms | 40.2 ms | 42.3 ms | 38.8 ms | -0.6 | 2.23-2.55 |
| gopher-lua | array | 1001 | 4.2 ms | 4.1 ms | 4.2 ms | 3.2 | - | 4.4 ms | 4.3 ms | 4.0 ms | 4.2 ms | 4.0 ms | 4.0 ms | -0.5 | 1.75-1.97 |
| gopher-lua | struct | 106710 | 35.1 us | 34.5 us | 36.7 us | 6.2 | - | 32.1 us | 34.9 us | 35.4 us | 35.4 us | 34.4 us | 37.7 us | -0.7 | 1.92-2.16 |
| gopher-lua | memtest | 1 | 4073.6 ms | 4064.6 ms | 4179.6 ms | 2.8 | - | - | - | - | 4115.9 ms | 3887.7 ms | 3993.1 ms | -3.6 | 1.98-2.51 |
| gopher-lua | ackermann | 383 | 9.3 ms | 8.7 ms | 9.4 ms | 7.9 | - | 9.6 ms | 8.7 ms | 8.9 ms | 9.2 ms | 8.5 ms | 8.8 ms | -0.7 | 1.71-1.85 |
| gopher-lua | digitalroot | 3 | 1364.5 ms | 1237.7 ms | 1417.2 ms | 13.2 | - | - | - | 1464.8 ms | 1349.1 ms | 1370.0 ms | 1360.2 ms | -0.5 | 2.18-2.71 |
| gopher-lua | structs | 77 | 45.7 ms | 44.7 ms | 47.7 ms | 6.6 | - | 47.0 ms | 54.4 ms | 49.6 ms | 46.0 ms | 44.4 ms | 43.9 ms | -1.2 | 4.19-4.70 |
| gopher-lua | structserialize | 59093 | 64.4 us | 63.7 us | 67.5 us | 6.0 | - | 56.3 us | 64.3 us | 65.7 us | 65.2 us | 61.1 us | 63.9 us | -0.7 | 1.95-2.05 |
| python | count | 12 | 109.1 ms | 102.3 ms | 112.8 ms | 9.6 | - | 117.6 ms | 109.0 ms | 110.3 ms | 109.0 ms | 105.0 ms | 100.5 ms | -0.0 | 3.54-3.82 |
| python | fib | 70 | 14.7 ms | 14.4 ms | 15.2 ms | 5.7 | - | 15.7 ms | 14.6 ms | 15.2 ms | 14.9 ms | 14.3 ms | 14.7 ms | -0.1 | 3.02-3.11 |
| python | array | 768 | 1.5 ms | 1.5 ms | 1.5 ms | 3.0 | - | 1.6 ms | 1.6 ms | 1.7 ms | 1.5 ms | 1.5 ms | 1.4 ms | -0.1 | 2.55-2.68 |
| python | struct | 174749 | 7.4 us | 7.4 us | 7.7 us | 4.3 | - | 7.9 us | 7.8 us | 7.3 us | 7.5 us | 7.2 us | 7.4 us | -0.0 | 2.29-2.34 |
| python | memtest | 1 | 2184.6 ms | 2153.3 ms | 2314.9 ms | 7.4 | - | - | - | - | 2218.0 ms | 2177.7 ms | 2260.9 ms | -0.1 | 2.14-2.66 |
| python | primes | 1 | 1446.5 ms | 1421.4 ms | 1492.5 ms | 4.9 | - | - | - | - | 1462.2 ms | 1463.3 ms | 1431.2 ms | -0.2 | 2.84-3.08 |
| python | ackermann | 332 | 3.9 ms | 3.9 ms | 4.0 ms | 3.2 | - | 4.1 ms | 4.1 ms | 4.3 ms | 4.0 ms | 3.9 ms | 4.1 ms | -0.0 | 3.03-3.51 |
| python | digitalroot | 12 | 106.6 ms | 104.7 ms | 110.2 ms | 5.2 | - | 116.6 ms | 107.3 ms | 112.1 ms | 108.0 ms | 104.4 ms | 109.7 ms | -0.0 | 2.95-3.20 |
| python | factorialiter | 76348 | 16.0 us | 15.6 us | 16.2 us | 3.4 | - | 16.5 us | 17.9 us | 16.7 us | 16.0 us | 16.5 us | 17.0 us | -0.0 | 2.67-2.73 |
| python | factorialrec | 50398 | 21.5 us | 20.9 us | 21.8 us | 4.2 | - | 22.2 us | 21.6 us | 21.5 us | 21.7 us | 20.3 us | 20.3 us | -0.1 | 2.45-2.90 |
| python | structs | 50 | 22.3 ms | 21.4 ms | 22.6 ms | 5.5 | - | 25.3 ms | 22.2 ms | 22.8 ms | 22.4 ms | 22.3 ms | 22.3 ms | -0.1 | 2.33-2.48 |
| python | structserialize | 74502 | 20.1 us | 16.6 us | 23.7 us | 35.7 | - | 19.0 us | 16.6 us | 17.1 us | 19.7 us | 17.4 us | 17.4 us | -0.1 | 2.33-2.91 |
| cx-vm | count | 31 | 46.8 ms est | 45.9 ms | 48.4 ms | 5.3 | 27.5 | 50.0 ms | 47.0 ms | 49.1 ms | 47.1 ms | - | - | - | 2.89-3.14 |
| cx-vm | fib | 67 | 14.7 ms est | 14.5 ms | 15.0 ms | 3.4 | 0.3 | 15.0 ms | 15.2 ms | 15.4 ms | 14.9 ms | - | - | - | 3.03-3.03 |
| cx-vm | array | 497 | 2.4 ms est | 2.3 ms | 2.4 ms | 4.9 | 1.2 | 2.5 ms | 2.3 ms | 2.4 ms | 2.4 ms | - | - | - | 2.93-3.10 |
| cx-vm | struct | 89711 | 12.2 us est | 12.1 us | 12.4 us | 2.8 | 1.0 | 14.0 us | 12.3 us | 12.6 us | 12.3 us | - | - | - | 2.66-2.86 |
| cx-vm | memtest | 1 | 1620.8 ms est | 1599.0 ms | 1631.3 ms | 2.0 | 14.6 | - | - | - | 1625.4 ms | - | - | - | 2.46-2.67 |
| cx-vm | primes | 2 | 791.5 ms est | 780.8 ms | 819.0 ms | 4.8 | 3.6 | - | - | 811.6 ms | 799.5 ms | - | - | - | 2.20-2.35 |
| cx-vm | ackermann | 503 | 2.8 ms est | 2.7 ms | 2.8 ms | 4.4 | 0.2 | 2.9 ms | 2.8 ms | 2.7 ms | 2.8 ms | - | - | - | 2.29-2.35 |
| cx-vm | digitalroot | 20 | 57.3 ms est | 56.4 ms | 58.0 ms | 2.8 | 2.0 | 66.4 ms | 60.5 ms | 58.6 ms | 57.6 ms | - | - | - | 2.13-2.23 |
| cx-vm | factorialiter | 125707 | 8.3 us est | 8.2 us | 8.4 us | 2.1 | 0.2 | 9.3 us | 8.4 us | 8.5 us | 8.3 us | - | - | - | 2.43-2.56 |
| cx-vm | factorialrec | 75304 | 17.4 us est | 15.3 us | 21.7 us | 36.3 | 10.6 | 16.5 us | 14.8 us | 15.4 us | 18.0 us | - | - | - | 2.71-2.75 |
| cx-vm | structs | 80 | 15.5 ms est | 15.3 ms | 15.9 ms | 4.3 | 9.5 | 17.9 ms | 15.7 ms | 15.8 ms | 15.7 ms | - | - | - | 3.11-3.29 |
| cx-vm | structserialize | 80579 | 13.0 us est | 12.7 us | 13.1 us | 3.0 | 4.0 | 13.6 us | 14.6 us | 13.1 us | 13.0 us | - | - | - | 2.33-2.44 |
| cx-image | count | 35 | 33.8 ms est | 32.7 ms | 34.5 ms | 5.5 | 0.9 | 32.3 ms | 33.4 ms | 34.0 ms | 33.6 ms | - | - | - | 2.14-2.35 |
| cx-image | fib | 77 | 14.8 ms est | 14.7 ms | 15.2 ms | 3.4 | 1.5 | 16.7 ms | 15.0 ms | 14.5 ms | 14.8 ms | - | - | - | 2.02-2.11 |
| cx-image | array | 575 | 2.3 ms est | 2.3 ms | 2.4 ms | 5.4 | 6.5 | 2.3 ms | 2.4 ms | 2.3 ms | 2.3 ms | - | - | - | 2.06-2.21 |
| cx-image | struct | 84470 | 14.2 us est | 14.1 us | 15.0 us | 6.5 | 3.1 | 12.9 us | 15.5 us | 15.0 us | 14.4 us | - | - | - | 2.54-2.64 |
| cx-image | memtest | 1 | 1847.8 ms est | 1841.6 ms | 1872.4 ms | 1.7 | 21.0 | - | - | - | 1857.1 ms | - | - | - | 2.01-2.30 |
| cx-image | primes | 2 | 762.3 ms est | 752.0 ms | 771.1 ms | 2.5 | 4.1 | - | - | 770.1 ms | 762.3 ms | - | - | - | 2.05-2.15 |
| cx-image | ackermann | 414 | 2.7 ms est | 2.6 ms | 2.7 ms | 3.3 | 1.5 | 2.8 ms | 2.6 ms | 2.7 ms | 2.7 ms | - | - | - | 1.69-1.81 |
| cx-image | digitalroot | 18 | 59.7 ms est | 58.6 ms | 64.6 ms | 10.0 | 9.1 | 70.6 ms | 61.8 ms | 60.0 ms | 61.0 ms | - | - | - | 1.49-1.57 |
| cx-image | factorialiter | 109717 | 8.6 us est | 8.4 us | 8.9 us | 5.8 | 11.6 | 8.5 us | 9.2 us | 8.1 us | 8.6 us | - | - | - | 2.30-2.41 |
| cx-image | factorialrec | 73351 | 14.9 us est | 14.7 us | 15.5 us | 4.7 | 4.0 | 16.2 us | 15.0 us | 15.3 us | 15.0 us | - | - | - | 1.93-2.01 |
| cx-image | structs | 70 | 17.7 ms est | 16.8 ms | 18.6 ms | 9.9 | 0.8 | 16.1 ms | 16.7 ms | 17.8 ms | 17.8 ms | - | - | - | 1.88-1.97 |
| cx-image | structserialize | 82115 | 14.6 us est | 14.0 us | 14.9 us | 6.2 | 4.0 | 13.7 us | 16.3 us | 15.4 us | 14.5 us | - | - | - | 2.46-2.51 |
| cx-llvm | count | 1073741824 | 1.16 ns est precomputed | 1.15 ns | 1.16 ns | 0.8 | 0.2 | 1.16 ns | 1.16 ns | 1.17 ns | 1.16 ns | - | - | - | 3.22-3.44 |
| cx-llvm | fib | 2980 | 382.4 us est | 378.5 us | 393.9 us | 4.0 | 1.6 | 372.2 us | 376.2 us | 385.5 us | 385.0 us | - | - | - | 2.54-2.82 |
| cx-llvm | array | 72532 | 15.3 us est | 15.2 us | 15.5 us | 1.5 | 0.1 | 15.7 us | 16.0 us | 15.3 us | 15.3 us | - | - | - | 2.22-2.34 |
| cx-llvm | struct | 46199270 | 25 ns est | 24 ns | 25 ns | 3.2 | 5.1 | 24 ns | 25 ns | 24 ns | 25 ns | - | - | - | 2.01-2.09 |
| cx-llvm | memtest | 145 | 9.3 ms est | 9.2 ms | 9.4 ms | 2.9 | 4.8 | 9.6 ms | 9.2 ms | 9.9 ms | 9.3 ms | - | - | - | 1.90-2.06 |
| cx-llvm | primes | 17 | 57.4 ms est | 56.6 ms | 62.1 ms | 9.7 | 0.4 | 59.6 ms | 60.5 ms | 58.9 ms | 58.1 ms | - | - | - | 3.04-3.33 |
| cx-llvm | ackermann | 11579 | 86.5 us est | 85.1 us | 89.0 us | 4.4 | 1.1 | 86.5 us | 84.2 us | 90.3 us | 87.0 us | - | - | - | 3.01-3.11 |
| cx-llvm | digitalroot | 952 | 1.2 ms est | 1.2 ms | 1.3 ms | 2.8 | 0.1 | 1.3 ms | 1.3 ms | 1.2 ms | 1.2 ms | - | - | - | 3.61-3.88 |
| cx-llvm | factorialiter | 11660165 | 97 ns est | 95 ns | 99 ns | 3.9 | 0.1 | 101 ns | 97 ns | 97 ns | 97 ns | - | - | - | 3.19-3.41 |
| cx-llvm | factorialrec | 3932777 | 290 ns est | 289 ns | 309 ns | 6.7 | 0.3 | 291 ns | 300 ns | 292 ns | 294 ns | - | - | - | 2.86-2.95 |
| cx-llvm | structs | 6918 | 167.6 us est | 167.1 us | 168.8 us | 1.0 | 0.5 | 167.8 us | 172.5 us | 167.7 us | 168.0 us | - | - | - | 3.76-4.32 |
| cx-llvm | structserialize | 74224869 | 15 ns est | 15 ns | 15 ns | 1.4 | 0.7 | 15 ns | 15 ns | 15 ns | 15 ns | - | - | - | 2.54-2.72 |

## Empty program

| form | layout | start to exit | start to S |
|---|---|---:|---:|
| c | a | 574.7 us | 487.0 us |
| c | b | 615.5 us | 502.8 us |
| cx-image | a | 2.1 ms | - |
| cx-image | b | 2.2 ms | - |
| cx-llvm | a | 667.4 us | - |
| cx-llvm | b | 745.0 us | - |
| cx-vm | a | 13.8 ms | - |
| cx-vm | b | 14.0 ms | - |
| go | a | 1.4 ms | 1.2 ms |
| go | b | 1.5 ms | 1.2 ms |
| gopher-lua | a | 51.9 ms | 2.7 ms |
| lua/5.4 | a | 1.6 ms | 1.4 ms |
| lua/5.5 | a | 1.3 ms | 1.1 ms |
| luajit | a | 1.2 ms | 1.0 ms |
| python | a | 21.1 ms | 16.8 ms |
| tinygo | a | 459.6 us | 399.0 us |
| tinygo | b | 493.2 us | 427.0 us |
