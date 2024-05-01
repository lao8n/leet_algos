package data_structures

// approaches
// * could just multiple n times... O(n)
// * but we can do better if need to multiply n times can make jumps - we can instead take powers of 2
func myPow(x float64, n int) float64 {
	// base cases
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1.0 / myPow(x, -1*n)
	}

	// odd
	if n%2 == 1 {
		return x * myPow(x*x, (n-1)/2)
	} else {
		// even
		return myPow(x*x, n/2)
	}
}
