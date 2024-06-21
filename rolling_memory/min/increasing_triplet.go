package data_structures

import "math"

// clarifying questions
// * can you have neg numbers? yes

// approaches
// * brute force try every combination of 3 indices i.e. O(n^3)
// * min, max and middle number. feels like a dynamic programming question where what you choose for first index affects later choices. top-down recursion with memoization or bottom iteration with tabulation - not going to try to optimize immediately
// * maybe we can do these with incrementing/decremneting pointers
// * maybe try somethign similar to palindrom where pick a midpoint and then look for min and max -> could have a rolling min but would need a map for max
// * maybe should use a stack? only append to stack if increasing

// problems
// * how maintain both say 3 - 5 as combo and 2... if find a single 8 then want 3-5-8, but if find two lower say 3-4 then want 2.. cannot forget either...
func increasingTriplet(nums []int) bool {
	rollingMin := math.MaxInt // best single number
	rollingMid := math.MaxInt // best mid number - can only add if rolling min is there
	for _, num := range nums {
		// valid combination
		if num > rollingMin && num > rollingMid {
			return true
		}
		// update rollingMid
		if num > rollingMin && num < rollingMid {
			rollingMid = num
		}
		// update rolling min
		if num < rollingMin {
			rollingMin = num
		}
	}
	return false
}
