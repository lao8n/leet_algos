package data_structures

import "math"

// either dynamic programming or kadane's algorithm - because have buy and sell decision
// dp either top down recursion with memoization or tabulation
// do the latter but only three states so just have to maintain all three
// make buying and selling separate decisions
// didn't need dp here, could have just added up all the troughs to peak
func maxProfit(prices []int) int {
	hold, notHold := math.MinInt, 0
	for _, p := range prices {
		hold = max(hold, notHold-p) // bought or held
		notHold = max(notHold, hold+p)
	}
	return max(hold, notHold)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
