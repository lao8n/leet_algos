package data_structures

import "fmt"

// interconnected problems combined for a solution -> dynamic programming
// dynamic programming has 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// try top-down where 3 considerations
// 1. base cases
// 2. recurrence - choice between buy, hold, or sell (add fee to sell)
// 3. state -> p, whether have something to buy
func maxProfit(prices []int, fee int) int {
	memoized := map[string]int{}
	return maxProfitRecursive(memoized, prices, fee, false)
}

func maxProfitRecursive(memoized map[string]int, prices []int, fee int, hold bool) int {
	// base cases
	if len(prices) == 1 {
		if hold == true {
			return prices[0] - fee
		} else {
			return 0
		}
	}

	// memoized
	key := fmt.Sprintf("%d %d", len(prices), hold)
	if memoizedProfit, ok := memoized[key]; ok {
		return memoizedProfit
	}

	// recurrence
	buyProfit, sellProfit, holdProfit := 0, 0, 0
	if hold == false {
		buyProfit = -prices[0] + maxProfitRecursive(memoized, prices[1:], fee, true)
	}
	if hold == true {
		sellProfit = prices[0] - fee + maxProfitRecursive(memoized, prices[1:], fee, false)
	}
	holdProfit = maxProfitRecursive(memoized, prices[1:], fee, hold)

	// pick maximum
	maxProfit := buyProfit
	if sellProfit > maxProfit {
		maxProfit = sellProfit
	}
	if holdProfit > maxProfit {
		maxProfit = holdProfit
	}

	memoized[key] = maxProfit

	return maxProfit
}
