package data_structures

import "math"

// clarifying questions
// * cannot use multiplication, division or mod operator? so can only use addition and subtraction?
// approaches
// * addition & subtraction?
// * bitwise operations?
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	quotient := 0
	divid := abs(dividend)
	divis := abs(divisor)
	for divid >= divis {
		sub := divis
		add := 1
		for divid >= sub<<1 {
			sub <<= 1
			add <<= 1
		}
		divid -= sub
		quotient += add
	}

	negative := (dividend < 0) != (divisor < 0)
	if negative {
		return -quotient
	}

	return quotient
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
