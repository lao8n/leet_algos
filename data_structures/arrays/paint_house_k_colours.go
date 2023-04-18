package data_structures

import "math"

// interconnected problems combined for a solution -> dynamic programming

// 2 approaches to dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// bottom up iteration with optimised tabulation
// 1. base cases
// 2. recurrence case
// 3. state variables = house number + available colours -> cost

// time complexity O(nk^2)
// space complexity O(k)
func minCostII(costs [][]int) int {
	n := len(costs)
	k := len(costs[0])
	prevHouse := make([]int, k) // 3 colours
	for i := 0; i < n; i++ {
		currHouse := make([]int, k)
		// find two minimum values
		for c := 0; c < k; c++ {
			// find min cost
			minCost := math.MaxInt
			for notc := 0; notc < k; notc++ {
				if notc != c && prevHouse[notc] < minCost {
					minCost = prevHouse[notc]
				}
			}
			// add cost
			currHouse[c] = costs[i][c] + minCost
		}
		prevHouse = currHouse
	}
	min_cost := math.MaxInt
	for i := 0; i < k; i++ {
		if prevHouse[i] < min_cost {
			min_cost = prevHouse[i]
		}
	}
	return min_cost
}
