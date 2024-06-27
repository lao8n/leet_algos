package data_structures

import "math"

// can probably do it in O(n) time..
// approach = loop through list backwards element by element - key thing is multiple rolling sum by 10 for each element

func reverse(x int) int {
	var reversed int
	neg := 1
	if x < 0 {
		neg = -1
		x *= -1
	}
	for x > 0 {
		remainder := x % 10 // last digit
		reversed = reversed*10 + remainder
		x = (x - remainder) / 10
	}
	result := neg * reversed
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}

// with overflow checks
func reverseOverflow(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32 || (rev == math.MaxInt32 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32 || (rev == math.MinInt32 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
