package data_structures

// brute force - try all possible buy prices with all possible sell prices O(n^2)
// rolling min and max -> for each partition what is the max and min?
func maxProfit(prices []int) int {
	rollingMin := make([]int, len(prices))
	rollingMin[0] = prices[0] // from the left
	rollingMax := make([]int, len(prices))
	rollingMax[len(prices)-1] = prices[len(prices)-1] // from the right

	// find rolling min and max O(n)
	for i := 1; i < len(prices); i++ {
		if prices[i] < rollingMin[i-1] {
			rollingMin[i] = prices[i]
		} else {
			rollingMin[i] = rollingMin[i-1]
		}
		j := len(prices) - i - 1
		if prices[j] > rollingMax[j] {
			rollingMax[j] = prices[j]
		} else {
			rollingMax[j] = rollingMax[j+1]
		}
	}

	// then just try all possible partitions O(n)
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if rollingMax[i]-rollingMin[i-1] > maxProfit {
			maxProfit = rollingMax[i] - rollingMin[i-1]
		}
	}
	return maxProfit
}
