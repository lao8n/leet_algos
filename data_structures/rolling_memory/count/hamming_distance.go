package data_structures

import "math"

func hammingDistance(x int, y int) int {
	diff := 0
	pow2 := 31.0
	for pow2 >= 0 {
		cur := int(math.Pow(2.0, pow2))
		// match
		if x >= cur && y >= cur {
			x -= cur
			y -= cur
		} else if x < cur && y < cur {
			// do nothing
		} else if x >= cur {
			x -= cur
			diff++
		} else if y >= cur {
			y -= cur
			diff++
		}
		pow2--
	}
	return diff
}
