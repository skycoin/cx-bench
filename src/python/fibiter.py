"""fib(25) with a loop, 25 steps, once per round. The loop variant of fib."""

# zero is the 0 the compiled forms read so no compiler precomputes a round, see
# README.md. It costs the same read here.
zero = 0


def fib(n):
    a, b = 0, 1
    for _ in range(n):
        a, b = b, a + b
    return a


def make(fnv, offset):
    def run(size):
        h, ok = offset, True
        for _ in range(size):
            r = fib(25 + zero)
            h = fnv(h, r)
            ok = ok and r == 75025
        return h, ok
    return run
