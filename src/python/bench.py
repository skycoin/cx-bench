"""bench.py runs the cx-bench programs in Python 3. See ../../README.md for
the protocol, which every language follows.

    python3 -I bench.py <name|number> <size> [-time=ms|us]   one run of two passes
    python3 -I bench.py empty                                 start and exit only
    python3 -I bench.py -server                               benchmarks from stdin
"""
import sys

sys.stderr.write("S")
sys.stderr.flush()

import importlib.util
import os
import time

NAMES = ["count", "fib", "array", "struct", "memtest", "primes", "ackermann",
         "digitalroot", "factorialiter", "factorialrec", "structs", "structserialize",
         "fibiter", "ackermanniter"]
FNV_OFFSET = 14695981039346656037
FNV_PRIME = 1099511628211
MASK64 = 0xFFFFFFFFFFFFFFFF


def fnv(h, v):
    return ((h ^ (v & MASK64)) * FNV_PRIME) & MASK64


def marker(phase):
    sys.stderr.write(phase)
    sys.stderr.flush()


def load(number):
    path = os.path.join(os.path.dirname(os.path.abspath(__file__)), NAMES[number - 1] + ".py")
    spec = importlib.util.spec_from_file_location(NAMES[number - 1], path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module.make(fnv, FNV_OFFSET)


def find(key):
    if key.isdigit():
        number = int(key)
        return number if 1 <= number <= len(NAMES) else 0
    return NAMES.index(key) + 1 if key in NAMES else 0


def elapsed(ns, unit):
    return ns // 1000000 if unit == "ms" else ns // 1000


def usage():
    sys.stderr.write("usage: python3 -I bench.py <name|number> <size> [-time=ms|us] | empty | -server\n")
    sys.exit(2)


def server():
    runs = {}
    for line in sys.stdin:
        fields = line.split()
        if not fields:
            continue
        number = int(fields[0]) if fields[0].isdigit() else -1
        if number == 0:
            break
        if 1 <= number <= len(NAMES) and len(fields) == 3:
            if number not in runs:
                runs[number] = load(number)
            start = time.perf_counter_ns()
            h, ok = runs[number](int(fields[1]))
            spent = elapsed(time.perf_counter_ns() - start, fields[2])
            sys.stdout.write("%d %d %d%s\n" % (number, h, spent, "" if ok else " FAIL"))
        else:
            sys.stdout.write("%d error\n" % number)
        sys.stdout.flush()
    marker("F")


def main():
    args = sys.argv[1:]
    if args == ["empty"]:
        marker("F")
        return 0
    if args == ["-server"]:
        server()
        return 0
    if len(args) not in (2, 3) or not args[1].isdigit():
        usage()
    number = find(args[0])
    unit = args[2][len("-time="):] if len(args) == 3 and args[2].startswith("-time=") else None
    if number == 0 or (len(args) == 3 and unit not in ("ms", "us")):
        usage()
    size = int(args[1])
    run = load(number)
    start = time.perf_counter_ns()
    h1, ok1 = run(size)
    marker("1")
    middle = time.perf_counter_ns()
    h2, ok2 = run(size)
    marker("2")
    stop = time.perf_counter_ns()
    ok = ok1 and ok2 and h1 == h2
    print(h1)
    print(NAMES[number - 1], "pass" if ok else "FAIL")
    if unit:
        print("time %s loop1=%d loop2=%d" % (unit, elapsed(middle - start, unit), elapsed(stop - middle, unit)))
    sys.stdout.flush()
    marker("F")
    return 0 if ok else 1


sys.exit(main())
