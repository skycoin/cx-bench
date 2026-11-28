"""Ackermann(3, 5) with an explicit stack, 42,438 steps, once per round. The
loop variant of ackermann. Each step is one call of the recursive form: pop m,
then push what the call would recurse into."""

# zero is the 0 the compiled forms read so no compiler precomputes a round, see
# README.md. It costs the same read here.
zero = 0


def ackermann(m, n):
    stack = [m]
    while stack:
        m = stack.pop()
        if m == 0:
            n += 1
        elif n == 0:
            stack.append(m - 1)
            n = 1
        else:
            stack.append(m - 1)
            stack.append(m)
            n -= 1
    return n


def make(fnv, offset):
    def run(size):
        h, ok = offset, True
        for _ in range(size):
            r = ackermann(3, 5 + zero)
            h = fnv(h, r)
            ok = ok and r == 253
        return h, ok
    return run
