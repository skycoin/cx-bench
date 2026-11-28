#include "bench.h"

static int32_t fib_loop(int32_t n) {
	int32_t a = 0, b = 1;
	for (int32_t k = 0; k < n; k++) {
		int32_t next = a + b;
		a = b;
		b = next;
	}
	return a;
}

/* fib(25) with a loop, 25 steps, once per round. The loop variant of fib. */
uint64_t bench_fibiter(int64_t size, int *pass) {
	uint64_t hash = FNV_OFFSET;
	int ok = 1;
	for (int64_t round = 0; round < size; round++) {
		int32_t result = fib_loop(25 + bench_zero);
		hash = fnv(hash, (uint64_t)result);
		ok = ok && result == 75025;
	}
	*pass = ok;
	return hash;
}
