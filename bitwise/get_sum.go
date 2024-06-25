package data_structures

import "fmt"

// approaches
// * byte manipulation << left and right shift. also have xor, and not

// specifics
// * reduce use-cases to sum of two positive integers and difference
// * xor is sum but without carry, carry = x & y << 1
func getSum(a int, b int) int {
	x, y := abs(a), abs(b)
	if x < y {
		return getSum(b, a)
	}
	// a determines the sign
	sign := 1
	if a <= 0 {
		sign = -1
	}
	// two cases
	if a*b >= 0 {
		// two pos integers
		for y != 0 {
			x, y = x^y, x&y<<1
			fmt.Println(x, y)
		}
	} else {
		// diff two positive integers x - y
		for y != 0 {
			x, y = x^y, ^x&y<<1
		}
	}
	return x * sign
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
