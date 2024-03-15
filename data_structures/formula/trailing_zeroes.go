package data_structures

// approaches
// * manually calculate the factorial -> % 10 to get number of trailing zeros O(n)
// * binary search for whether n is one of the key amounts... to have a trailing zero need to be a multiple of 10 = 2 * 5.. there are loads of 2s so basically counting trailing zeros is about counting multiples of 5?
func trailingZeroes(n int) int {
	// how many times does 5 go into n?
	numFives := 0
	divisor5 := 5
	for n >= divisor5 {
		// process each multiple of 5
		numFives += n / divisor5
		divisor5 *= 5
	}
	return numFives
}
