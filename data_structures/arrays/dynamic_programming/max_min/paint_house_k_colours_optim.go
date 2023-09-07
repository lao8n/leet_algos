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

// time complexity O(nk)
// space complexity O(1)
func minCostII(costs [][]int) int {
	n := len(costs)
	k := len(costs[0])
	prevBestHouse, prevSecondBestHouse, prevBestHouseColour := 0, 0, -1
	for i := 0; i < n; i++ {
		currBestHouse, currSecondBestHouse, currBestHouseColour := math.MaxInt, math.MaxInt, -1
		for c := 0; c < k; c++ {
			var cHouse int
			if c != prevBestHouseColour {
				cHouse = costs[i][c] + prevBestHouse
			} else {
				cHouse = costs[i][c] + prevSecondBestHouse
			}
			if cHouse < currBestHouse {
				currSecondBestHouse = currBestHouse
				currBestHouse = cHouse
				currBestHouseColour = c
			} else if cHouse < currSecondBestHouse {
				currSecondBestHouse = cHouse
			}
		}
		prevBestHouse, prevSecondBestHouse, prevBestHouseColour = currBestHouse, currSecondBestHouse, currBestHouseColour
	}
	return prevBestHouse
}
