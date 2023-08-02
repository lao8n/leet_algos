package data_structures

// these are interconnected subproblems combined for a solution -> dynamic programming
// can either do bottom up tabulation or top down recursion with memoization -> recursion is better if we can skip some (but here we cannot) so tabulation without the overhead of recursion is better.
// however will do recursion as it's harder to code and want the practice
func tribonacci(n int) int {
	tribonacciMap := make(map[int]int, n+1)
	// base cases
	tribonacciMap[0] = 0
	tribonacciMap[1] = 1
	tribonacciMap[2] = 1
	return tribonacciRecursive(n, tribonacciMap)
}

func tribonacciRecursive(n int, tribonacciMap map[int]int) int {
	// recursive cases
	if memoized, ok := tribonacciMap[n]; ok {
		return memoized
	}
	// relies upon map being a pointer which can be updated
	n1 := tribonacciRecursive(n-1, tribonacciMap)
	n2 := tribonacciRecursive(n-2, tribonacciMap)
	n3 := tribonacciRecursive(n-3, tribonacciMap)
	result := n1 + n2 + n3
	tribonacciMap[n] = result
	return result
}
