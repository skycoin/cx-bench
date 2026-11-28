package main

import "testing"

// hashesOfGoForm were printed by the Go form, src/go, at sizes 1 to 3.
var hashesOfGoForm = []struct {
	bench string
	size  int
	hash  uint64
}{
	{"count", 1, 17727803403943808447},
	{"count", 2, 5677473476889670253},
	{"count", 3, 16217943649578642327},
	{"fib", 1, 12702064428108188252},
	{"fib", 2, 791634122422345687},
	{"fib", 3, 13398842690933890674},
	{"array", 1, 7111567473559590409},
	{"array", 2, 15731497248241456804},
	{"array", 3, 5464382542771218395},
	{"struct", 1, 12635340564950203717},
	{"struct", 2, 16575311544170438053},
	{"struct", 3, 10691217752633942213},
	{"memtest", 1, 874058405928975501},
	{"memtest", 2, 16767549263620469557},
	{"memtest", 3, 10437026336130220317},
	{"primes", 1, 13592104564658177911},
	{"primes", 2, 8648444381812236705},
	{"primes", 3, 7227343155601327507},
	{"ackermann", 1, 12638349928276617224},
	{"ackermann", 2, 762171597903747151},
	{"ackermann", 3, 16930126232538003062},
	{"digitalroot", 1, 11654704861874033752},
	{"digitalroot", 2, 7624075781987266231},
	{"digitalroot", 3, 17787510751109177378},
	{"factorialiter", 1, 16397588511422435985},
	{"factorialiter", 2, 7655353161881067309},
	{"factorialiter", 3, 8361780343919622633},
	{"factorialrec", 1, 16397588511422435985},
	{"factorialrec", 2, 7655353161881067309},
	{"factorialrec", 3, 8361780343919622633},
	{"structs", 1, 14112722035046580255},
	{"structs", 2, 11382818394927254221},
	{"structs", 3, 16734513219187032663},
	{"structserialize", 1, 12638214688346347271},
	{"structserialize", 2, 644099442219955380},
	{"structserialize", 3, 13585581175438645566},
	{"fibiter", 1, 12702064428108188252},
	{"fibiter", 2, 791634122422345687},
	{"fibiter", 3, 13398842690933890674},
	{"ackermanniter", 1, 12638349928276617224},
	{"ackermanniter", 2, 762171597903747151},
	{"ackermanniter", 3, 16930126232538003062},
}

func TestHashComputeMatchesTheGoForm(t *testing.T) {
	for _, want := range hashesOfGoForm {
		if got := hashCompute(want.bench, want.size); got != want.hash {
			t.Errorf("%s %d: math gives %d, the Go form printed %d", want.bench, want.size, got, want.hash)
		}
	}
}

func TestBenchRowsHoldEveryBenchmarkOnce(t *testing.T) {
	seen := map[string]int{}
	for _, bench := range benchRows {
		seen[bench]++
	}
	for _, bench := range benchNames {
		if seen[bench] != 1 {
			t.Errorf("%s is %d times in benchRows, want once", bench, seen[bench])
		}
	}
	if len(benchRows) != len(benchNames) {
		t.Errorf("benchRows has %d names, benchNames %d", len(benchRows), len(benchNames))
	}
}
