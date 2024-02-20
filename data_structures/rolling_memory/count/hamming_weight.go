package data_structures

import "math"

// approaches
// * convert to binary string
// * loop through num until 0 - don't know how big it is so go from smallest to largest O(log2 n) - actually we do know and we should start from the largest
func hammingWeight(num uint32) int {
	count := 0
	pow2 := 31.0
	for num > 0 {
		cur := uint32(math.Pow(2.0, pow2))
		if num >= cur {
			count++
			num -= cur
		}
		pow2--
	}
	return count
}
