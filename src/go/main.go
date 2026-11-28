// Command bench runs the cx-bench programs in Go. See ../../README.md for
// the protocol, which every language follows.
//
//	bench <name|number> <size> [-time=ms|us]   one run of two passes
//	bench empty                                 start and exit only
//	bench -server                               benchmarks from stdin
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	fnvOffset = 14695981039346656037
	fnvPrime  = 1099511628211
)

func fnv(hash uint64, value uint64) uint64 {
	return (hash ^ value) * fnvPrime
}

type benchProgram struct {
	name string
	run  func(size int) (uint64, bool)
}

// benchPrograms holds the benchmarks in their fixed order, number 1 first.
var benchPrograms = []benchProgram{
	{"count", benchCount},
	{"fib", benchFib},
	{"array", benchArray},
	{"struct", benchStruct},
	{"memtest", benchMemtest},
	{"primes", benchPrimes},
	{"ackermann", benchAckermann},
	{"digitalroot", benchDigitalRoot},
	{"factorialiter", benchFactorialIter},
	{"factorialrec", benchFactorialRec},
	{"structs", benchStructs},
	{"structserialize", benchStructSerialize},
	{"fibiter", benchFibIter},
	{"ackermanniter", benchAckermannIter},
}

// benchSize is read again before each pass, so no compiler can reuse the
// first pass for the second.
var benchSize int

func marker(phase byte) {
	os.Stderr.Write([]byte{phase})
}

func programFind(key string) int {
	if number, err := strconv.Atoi(key); err == nil && number >= 1 && number <= len(benchPrograms) {
		return number
	}
	for i, program := range benchPrograms {
		if program.name == key {
			return i + 1
		}
	}
	return 0
}

func elapsedIn(duration time.Duration, unit string) int64 {
	if unit == "ms" {
		return duration.Milliseconds()
	}
	return duration.Microseconds()
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: bench <name|number> <size> [-time=ms|us] | bench empty | bench -server")
	for i, program := range benchPrograms {
		fmt.Fprintf(os.Stderr, "%d %s\n", i+1, program.name)
	}
	os.Exit(2)
}

func serverRun() {
	input := bufio.NewScanner(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	for input.Scan() {
		fields := strings.Fields(input.Text())
		if len(fields) == 0 {
			continue
		}
		number, _ := strconv.Atoi(fields[0])
		if number == 0 {
			break
		}
		if number < 1 || number > len(benchPrograms) || len(fields) != 3 {
			fmt.Fprintf(output, "%d error\n", number)
			output.Flush()
			continue
		}
		size, _ := strconv.Atoi(fields[1])
		benchSize = size
		start := time.Now()
		hash, pass := benchPrograms[number-1].run(benchSize)
		elapsed := elapsedIn(time.Since(start), fields[2])
		if pass {
			fmt.Fprintf(output, "%d %d %d\n", number, hash, elapsed)
		} else {
			fmt.Fprintf(output, "%d %d %d FAIL\n", number, hash, elapsed)
		}
		output.Flush()
	}
	marker('F')
}

func main() {
	marker('S')
	if len(os.Args) == 2 && os.Args[1] == "empty" {
		marker('F')
		return
	}
	if len(os.Args) == 2 && os.Args[1] == "-server" {
		stop := profileStart()
		serverRun()
		stop()
		return
	}
	if len(os.Args) < 3 || len(os.Args) > 4 {
		usage()
	}
	number := programFind(os.Args[1])
	size, err := strconv.Atoi(os.Args[2])
	if number == 0 || err != nil || size < 0 {
		usage()
	}
	unit := ""
	if len(os.Args) == 4 {
		unit = strings.TrimPrefix(os.Args[3], "-time=")
		if unit != "ms" && unit != "us" {
			usage()
		}
	}
	program := benchPrograms[number-1]
	benchSize = size
	start := time.Now()
	hash1, pass1 := program.run(benchSize)
	marker('1')
	middle := time.Now()
	hash2, pass2 := program.run(benchSize)
	marker('2')
	end := time.Now()
	pass := pass1 && pass2 && hash1 == hash2
	fmt.Println(hash1)
	if pass {
		fmt.Println(program.name, "pass")
	} else {
		fmt.Println(program.name, "FAIL")
	}
	if unit != "" {
		fmt.Printf("time %s loop1=%d loop2=%d\n", unit, elapsedIn(middle.Sub(start), unit), elapsedIn(end.Sub(middle), unit))
	}
	marker('F')
	if !pass {
		os.Exit(1)
	}
}
