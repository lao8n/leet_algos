package data_structures

import "math"

// clarifying questions

// approaches
// * multiply up from 3 (or down from n) to see if hit n or hit 1 which will cost O(x)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	quo, _ := math.Modf(math.Log(float64(n)) / math.Log(3))
	return n == int(math.Pow(3, quo))
}
