package data_structures

import "math"

// interconnected subproblems combined for a solution -> dynamic programming
// either top down recursion with memoization or bottom up tabulation

// 3 things to define
// 1. base cases
// 2. recurrence relation
// 3. state variables

// greedy algorithm where choose largest coin first? doesn't work because might not be able to equal amount

// complexity
// time O(n x a) - a = amount and n is number of coins
func coinChange(coins []int, amount int) int {
	memoized := make([]int, amount+1)
	return coinChangeMemoize(coins, amount, memoized)
}

func coinChangeMemoize(coins []int, amount int, memoized []int) int {
	// base cases
	if amount == 0 {
		return 0
	}
	minCoin := math.MaxInt32
	for _, coin := range coins {
		if coin < minCoin {
			minCoin = coin
		}
	}
	if amount < minCoin {
		return -1
	}
	// memoized
	if memoized[amount] != 0 {
		return memoized[amount]
	}

	// state = remaining amount, coins -> number of coins used
	// recurrence relation
	leastNumCoins := math.MaxInt32
	for _, coin := range coins {
		numCoins := coinChangeMemoize(coins, amount-coin, memoized) + 1
		if numCoins != 0 && numCoins < leastNumCoins {
			leastNumCoins = numCoins
		}
	}
	if leastNumCoins == math.MaxInt32 {
		leastNumCoins = -1
	}
	memoized[amount] = leastNumCoins

	return leastNumCoins
}
