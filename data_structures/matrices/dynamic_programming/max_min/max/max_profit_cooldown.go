package data_structures

import "math"

// interrelated subproblems combined for a solution -> dynamic programming
// two approaches to dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// for a bottom-up approach 3 key things
// 1. base cases
// 2. recurrence relation - choice between buy, hold, or sell - because we only depend upon
// t-1 don't need full table
// 3. state -> args = day index, hold, cooldown, return = profit

// complexity
// time = O(n)
// space = O(1)
func maxProfit(prices []int) int {
	sold, held, reset := math.MinInt, math.MinInt, 0
	for _, p := range prices {
		sold, held, reset = held+p, max(held, reset-p), max(reset, sold)
	}
	return max(sold, reset)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
