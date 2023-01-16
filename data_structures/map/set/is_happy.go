package data_structures

// naive: recurse/for loop through each calculation - with some recursion limit until give up
// clever: is there some clever way to tell in adavnce.
// perhaps we can go backwards? the only numbers that have equal 1 are powers of ten 1, 10, 100 etc.
// perhaps we can build a table of all those options.. memoize everything.. but n can be very large 2^31
// detect cycle with hashset
func isHappy(n int) bool {
	seenNums := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}
		if _, ok := seenNums[n]; ok {
			// detected cycle
			return false
		}
		// new number
		seenNums[n] = true
		n = nextNumber(n)
	}
	return false
}

func nextNumber(n int) int {
	nextNumber := 0
	for n > 0 {
		remainder := n % 10 // 253 % 10 = 3
		nextNumber = nextNumber + remainder*remainder
		n = n / 10 // no need to minus remainder first
	}
	return nextNumber
}
