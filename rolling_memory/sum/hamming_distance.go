package data_structures

import "math"

// approach is ot convert each into a binary number removing powers of 2 until nothing left
// as remove each compare digits
// it is a purely local comparison so don't need to roll anything
// we have a max number
func hammingDistance(x int, y int) int {
	pow2 := 31.0
	countDiff := 0
	for pow2 >= 0 { // go until 2^0 = 1
		xFlag, yFlag := false, false
		//fmt.Println(math.Pow(2, pow2), x, y)
		if x >= int(math.Pow(2, pow2)) {
			xFlag = true
			x -= int(math.Pow(2, pow2))
		}
		if y >= int(math.Pow(2, pow2)) {
			yFlag = true
			y -= int(math.Pow(2, pow2))
		}
		if xFlag != yFlag {
			countDiff++
		}
		pow2 -= 1.0
	}
	return countDiff
}
