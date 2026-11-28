/* bench runs the cx-bench programs in C. See ../../README.md for the
   protocol, which every language follows.

	bench <name|number> <size> [-time=ms|us]   one run of two passes
	bench empty                                 start and exit only
	bench -server                               benchmarks from stdin */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include "bench.h"

#define PROGRAM_COUNT 14

typedef uint64_t (*bench_function)(int64_t size, int *pass);

static const struct {
	const char *name;
	bench_function run;
} programs[PROGRAM_COUNT] = {
	{"count", bench_count},
	{"fib", bench_fib},
	{"array", bench_array},
	{"struct", bench_struct},
	{"memtest", bench_memtest},
	{"primes", bench_primes},
	{"ackermann", bench_ackermann},
	{"digitalroot", bench_digitalroot},
	{"factorialiter", bench_factorialiter},
	{"factorialrec", bench_factorialrec},
	{"structs", bench_structs},
	{"structserialize", bench_structserialize},
	{"fibiter", bench_fibiter},
	{"ackermanniter", bench_ackermanniter},
};

/* Read again before each pass, so the compiler cannot reuse the first pass
   for the second. */
static volatile int64_t bench_size;

volatile int32_t bench_zero = 0;

static void marker(char phase) {
	fputc(phase, stderr);
	fflush(stderr);
}

static int64_t now_ns(void) {
	struct timespec t;
	clock_gettime(CLOCK_MONOTONIC, &t);
	return (int64_t)t.tv_sec * 1000000000 + t.tv_nsec;
}

static long long elapsed_in(int64_t ns, const char *unit) {
	return (long long)(strcmp(unit, "ms") == 0 ? ns / 1000000 : ns / 1000);
}

static int program_find(const char *key) {
	char *end;
	long number = strtol(key, &end, 10);
	if (*key && !*end)
		return number >= 1 && number <= PROGRAM_COUNT ? (int)number : 0;
	for (int i = 0; i < PROGRAM_COUNT; i++)
		if (strcmp(programs[i].name, key) == 0)
			return i + 1;
	return 0;
}

static void usage(void) {
	fprintf(stderr, "usage: bench <name|number> <size> [-time=ms|us] | bench empty | bench -server\n");
	for (int i = 0; i < PROGRAM_COUNT; i++)
		fprintf(stderr, "%d %s\n", i + 1, programs[i].name);
	exit(2);
}

static void server_run(void) {
	char line[256], unit[8];
	long long number, size;
	while (fgets(line, sizeof line, stdin)) {
		int fields = sscanf(line, "%lld %lld %7s", &number, &size, unit);
		if (fields < 1)
			continue;
		if (number == 0)
			break;
		if (number < 1 || number > PROGRAM_COUNT || fields != 3) {
			printf("%lld error\n", number);
			fflush(stdout);
			continue;
		}
		int pass;
		bench_size = size;
		int64_t start = now_ns();
		uint64_t hash = programs[number - 1].run(bench_size, &pass);
		long long elapsed = elapsed_in(now_ns() - start, unit);
		printf("%lld %llu %lld%s\n", number, (unsigned long long)hash, elapsed, pass ? "" : " FAIL");
		fflush(stdout);
	}
	marker('F');
}

int main(int argc, char **argv) {
	marker('S');
	if (argc == 2 && strcmp(argv[1], "empty") == 0) {
		marker('F');
		return 0;
	}
	if (argc == 2 && strcmp(argv[1], "-server") == 0) {
		server_run();
		return 0;
	}
	if (argc < 3 || argc > 4)
		usage();
	int number = program_find(argv[1]);
	char *end;
	long long size = strtoll(argv[2], &end, 10);
	if (number == 0 || *end || size < 0)
		usage();
	const char *unit = NULL;
	if (argc == 4) {
		if (strncmp(argv[3], "-time=", 6) != 0)
			usage();
		unit = argv[3] + 6;
		if (strcmp(unit, "ms") != 0 && strcmp(unit, "us") != 0)
			usage();
	}
	int pass1, pass2;
	bench_size = size;
	int64_t start = now_ns();
	uint64_t hash1 = programs[number - 1].run(bench_size, &pass1);
	marker('1');
	int64_t middle = now_ns();
	uint64_t hash2 = programs[number - 1].run(bench_size, &pass2);
	marker('2');
	int64_t end_ns = now_ns();
	int pass = pass1 && pass2 && hash1 == hash2;
	printf("%llu\n%s %s\n", (unsigned long long)hash1, programs[number - 1].name, pass ? "pass" : "FAIL");
	if (unit)
		printf("time %s loop1=%lld loop2=%lld\n", unit, elapsed_in(middle - start, unit), elapsed_in(end_ns - middle, unit));
	fflush(stdout);
	marker('F');
	return pass ? 0 : 1;
}
