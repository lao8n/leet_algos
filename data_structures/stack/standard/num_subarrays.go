package data_structures

import "math"

// approach
// * monotonically decreasing stack
// https://leetcode.com/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/
// specifics
// * key formula is number of subarrays defined by c (c+1) / 2
// this is the formula for c natural numbers because between i and j,
// for each starting index i (which there are c options), there are
// c - i subarrays starting at index i
func numberOfSubarrays(nums []int) int64 {
	var numWays int64

	stack := make([]int, 0)
	nums = append(nums, math.MaxInt) // handle end condition

	for _, num := range nums {
		// if new num is bigger than stack value
		for len(stack) > 0 && num > peek(stack) {
			p := peek(stack)
			var c int64
			for len(stack) > 0 && peek(stack) == p {
				c++
				stack = stack[:len(stack)-1] // pop
			}
			numWays += c * (c + 1) / 2
		}
		// if new num is less add value to stack
		stack = append(stack, num)
	}
	return numWays
}

func peek(stack []int) int {
	return stack[len(stack)-1]
}
