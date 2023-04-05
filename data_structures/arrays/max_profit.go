package data_structures

func maxProfit(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	if k >= len(prices)/2 {
		// If k is large enough to perform all possible transactions,
		// use the simple solution that buys on every price dip and sells on every price rise.
		maxProfit := 0
		for i := 1; i < len(prices); i++ {
			if prices[i] > prices[i-1] {
				maxProfit += prices[i] - prices[i-1]
			}
		}
		return maxProfit
	}
	// If k is smaller than the number of prices, we use a dynamic programming approach.
	dp := make([][]int, k+1) // k n - state
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, len(prices))
	}
	for i := 1; i < k+1; i++ {
		maxDiff := -prices[0]
		for j := 1; j < len(prices); j++ {
			maxDiff = max(maxDiff, dp[i-1][j-1]-prices[j-1])
			dp[i][j] = max(dp[i][j-1], maxDiff+prices[j])
		}
	}
	return dp[k][len(prices)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
