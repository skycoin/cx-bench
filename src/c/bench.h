#ifndef CXBENCH_BENCH_H
#define CXBENCH_BENCH_H
#include <stdint.h>

#define FNV_OFFSET 14695981039346656037ULL
#define FNV_PRIME 1099511628211ULL

/* bench_zero is 0, but volatile, so a compiler cannot hoist the work of a
   round out of the loop of rounds or replace a loop with its closed form. */
extern volatile int32_t bench_zero;

static inline uint64_t fnv(uint64_t h, uint64_t v) { return (h ^ v) * FNV_PRIME; }

uint64_t bench_count(int64_t size, int *pass);
uint64_t bench_fib(int64_t size, int *pass);
uint64_t bench_array(int64_t size, int *pass);
uint64_t bench_struct(int64_t size, int *pass);
uint64_t bench_memtest(int64_t size, int *pass);
uint64_t bench_primes(int64_t size, int *pass);
uint64_t bench_ackermann(int64_t size, int *pass);
uint64_t bench_digitalroot(int64_t size, int *pass);
uint64_t bench_factorialiter(int64_t size, int *pass);
uint64_t bench_factorialrec(int64_t size, int *pass);
uint64_t bench_structs(int64_t size, int *pass);
uint64_t bench_structserialize(int64_t size, int *pass);
uint64_t bench_fibiter(int64_t size, int *pass);
uint64_t bench_ackermanniter(int64_t size, int *pass);
#endif
