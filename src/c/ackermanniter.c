#include "bench.h"

/* A(3, 5) holds at most 252 pending m values. */
#define ACKERMANN_STACK 1024

/* Each step of the loop is one call of the recursive form: pop m, then
   push what the call would recurse into. */
static int32_t ackermann_loop(int32_t m, int32_t n) {
	int32_t stack[ACKERMANN_STACK];
	int32_t top = 0;
	stack[top++] = m;
	while (top > 0) {
		m = stack[--top];
		if (m == 0) {
			n = n + 1;
		} else if (n == 0) {
			stack[top++] = m - 1;
			n = 1;
		} else {
			stack[top++] = m - 1;
			stack[top++] = m;
			n = n - 1;
		}
	}
	return n;
}

/* Ackermann(3, 5) with an explicit stack, 42,438 steps, once per round. The
   loop variant of ackermann. */
uint64_t bench_ackermanniter(int64_t size, int *pass) {
	uint64_t hash = FNV_OFFSET;
	int ok = 1;
	for (int64_t round = 0; round < size; round++) {
		int32_t result = ackermann_loop(3, 5 + bench_zero);
		hash = fnv(hash, (uint64_t)result);
		ok = ok && result == 253;
	}
	*pass = ok;
	return hash;
}
