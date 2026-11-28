package main

func fibLoop(n int32) int32 {
	a, b := int32(0), int32(1)
	for k := int32(0); k < n; k++ {
		a, b = b, a+b
	}
	return a
}

// benchFibIter runs fib(25) with a loop, 25 steps, once per round. It is the
// loop variant of fib.
func benchFibIter(size int) (uint64, bool) {
	hash := uint64(fnvOffset)
	pass := true
	for round := 0; round < size; round++ {
		result := fibLoop(25 + zero())
		hash = fnv(hash, uint64(result))
		pass = pass && result == 75025
	}
	return hash, pass
}
