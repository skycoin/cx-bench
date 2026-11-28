package main

import "strconv"

const (
	fnvOffset = 14695981039346656037
	fnvPrime  = 1099511628211
)

func fnv(hash, value uint64) uint64 { return (hash ^ value) * fnvPrime }

// hashRound returns the words one round of a benchmark hashes, from the
// math rather than from running it: closed forms, and for digitalroot the
// root and persistence of each n found by summing decimal digits.
func hashRound(bench string, round uint64) []uint64 {
	switch bench {
	case "count":
		return []uint64{999999 * 1000000 / 2}
	case "fib", "fibiter":
		return []uint64{75025}
	case "array":
		return []uint64{150 * 149 / 2}
	case "struct":
		return []uint64{3 * 99 * 100 / 2}
	case "memtest":
		const n, walk, all = 1048576, 65536, 0xFFFFFFFF
		// i and all - i are read back once each, then each pattern twice
		// and its inverse once, with no mismatch.
		words := []uint64{0, n * all}
		for bit := 0; bit < 32; bit++ {
			words = append(words, 0, walk<<bit)
		}
		for _, pattern := range []uint64{0, all, 0x55555555, 0xAAAAAAAA} {
			words = append(words, 0, n*(pattern+all))
		}
		return words
	case "primes":
		// 148933 primes below 2,000,000, 25997 below 300,000, and the
		// Mersenne prime exponents up to 61.
		words := []uint64{148933, 25997}
		mersenne := map[uint64]bool{3: true, 5: true, 7: true, 13: true, 17: true, 19: true, 31: true, 61: true}
		for repeat := 0; repeat < 20; repeat++ {
			for _, exponent := range []uint64{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61} {
				if mersenne[exponent] {
					words = append(words, 1)
				} else {
					words = append(words, 0)
				}
			}
		}
		return words
	case "ackermann", "ackermanniter":
		// A(3, n) = 2^(n+3) - 3.
		return []uint64{1<<8 - 3}
	case "digitalroot":
		return digitalRootWords()
	case "factorialiter", "factorialrec":
		words := []uint64{1}
		for n, factorial := uint64(1), uint64(1); n <= 20; n++ {
			factorial *= n
			words = append(words, factorial)
		}
		return words
	case "structs":
		// sum of c + round and 3c for c below 100,000.
		return []uint64{4*99999*100000/2 + 100000*round}
	case "structserialize":
		// value + 2 value + sum of 16 value + k for k below 16, value the
		// round mod 2^20.
		return []uint64{259*(round&1048575) + 120}
	}
	return nil
}

var digitalRootCache []uint64

func digitalRootWords() []uint64 {
	if digitalRootCache != nil {
		return digitalRootCache
	}
	for n := 1; n < 100000; n++ {
		text, persistence := strconv.Itoa(n), uint64(0)
		for len(text) > 1 {
			sum := 0
			for _, digit := range text {
				sum += int(digit - '0')
			}
			text, persistence = strconv.Itoa(sum), persistence+1
		}
		digitalRootCache = append(digitalRootCache, uint64(text[0]-'0')+10*persistence)
	}
	return digitalRootCache
}

// hashCompute returns the hash of run(size), one pass of size rounds, in
// size times the words of a round hashes and none of the work.
func hashCompute(bench string, size int) uint64 {
	hash := uint64(fnvOffset)
	roundDependent := bench == "structs" || bench == "structserialize"
	words := hashRound(bench, 0)
	for round := 0; round < size; round++ {
		if roundDependent {
			words = hashRound(bench, uint64(round))
		}
		for _, word := range words {
			hash = fnv(hash, word)
		}
	}
	if bench == "array" {
		// The array ends sorted, 1 to 150, or all zero when nothing ran.
		for value := uint64(1); value <= 150; value++ {
			if size == 0 {
				hash = fnv(hash, 0)
			} else {
				hash = fnv(hash, value)
			}
		}
	}
	return hash
}
