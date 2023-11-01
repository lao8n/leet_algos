package data_structures

// approaches
// * could brute force recursion but O(2^n)
// * instead can use memoization
func fib(n int) int {
	// setup up mem slice
	mem := make([]int, n+1)

	// recursive call
	return recursiveFib(n, mem)
}

func recursiveFib(n int, mem []int) int {
	// base case
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	// mem
	if mem[n] != 0 {
		return mem[n]
	}

	// recursive
	fib := recursiveFib(n-1, mem) + recursiveFib(n-2, mem)
	mem[n] = fib
	// return
	return fib
}
