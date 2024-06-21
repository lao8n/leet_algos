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
	result := 1.0
	for n > 0 {
		// odd
		if n%2 == 1 {
			result *= x
			n--
		} else {
			x *= x
			n /= 2
		}
	}
	return result
}
