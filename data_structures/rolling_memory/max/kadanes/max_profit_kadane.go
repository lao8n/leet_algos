package data_structures

func maxProfit(prices []int) int {
	max_profit := 0
	rolling_profit := 0
	for i := 1; i < len(prices); i++ {
		marginal_profit := prices[i] - prices[i-1]
		rolling_profit = rolling_profit + marginal_profit
		if rolling_profit < marginal_profit {
			rolling_profit = marginal_profit
		}
		if rolling_profit > max_profit {
			max_profit = rolling_profit
		}
	}
	return max_profit
}
