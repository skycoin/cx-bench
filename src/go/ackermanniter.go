package main

// ackermannStack is the explicit stack size. A(3, 5) holds
// at most 252 pending m values.
const ackermannStack = 1024

// ackermannLoop takes one step per call of the recursive form: pop m, then
// push what the call would recurse into.
func ackermannLoop(m, n int32) int32 {
	var stack [ackermannStack]int32
	top := 0
	stack[top] = m
	top++
	for top > 0 {
		top--
		m = stack[top]
		switch {
		case m == 0:
			n++
		case n == 0:
			stack[top] = m - 1
			top++
			n = 1
		default:
			stack[top] = m - 1
			stack[top+1] = m
			top += 2
			n--
		}
	}
	return n
}

// benchAckermannIter runs Ackermann(3, 5) with an explicit stack, 42,438
// steps, once per round. It is the loop variant of ackermann.
func benchAckermannIter(size int) (uint64, bool) {
	hash := uint64(fnvOffset)
	pass := true
	for round := 0; round < size; round++ {
		result := ackermannLoop(3, 5+zero())
		hash = fnv(hash, uint64(result))
		pass = pass && result == 253
	}
	return hash, pass
}
