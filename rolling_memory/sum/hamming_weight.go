package data_structures

import (
	"fmt"
	"math"
)

// we do not have the input binary string itself
// 1. recursively reduce to 0 - but the problem is don't know which power of 2 to remove
// 2. convert into binary string
func hammingWeight(num uint32) int {
	bit := 31.0
	count := 0
	for num >= 0 && bit >= 0 {
		pow2 := uint32(math.Pow(2, bit))
		if num >= pow2 {
			num -= pow2
			count++
			fmt.Println(num, count, pow2, bit)
		}
		bit--
	}
	return count
}
