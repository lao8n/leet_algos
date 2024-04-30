package data_structures

import "fmt"

// interdependent problems combined for a solution -> dynamic programming

// 2 main approaches
// 1. top-down recursion with memoization <-
// 2. bottom-up iteration with tabulation

// key question is what does tabulation/memoization look like?
// depends upon what defines state...
// 3 things
// 1. day
// 2. k transactions
// 3. bought price (can be nil) -> this leads to two many options - better to separate profit into buy & sell
// will try top-down memoization because complicated state easier to manage typically
// however typically with recursion you: recurse forward -> accumulate answer backward -  could accumulate answer
// key trick is separate buy & sell into two separate actions so state doesn't explode because of price
// i.e. use d, k, o and d, k, p

// time complexity = O(nk)
// space complexity = O(nk)
func maxProfit(k int, prices []int) int {
	data := data{prices: prices, memoized: make(map[string]int)} // d, k, o
	return data.maxProfitRecursive(0, k, false)
}

type data struct {
	prices   []int
	memoized map[string]int
}

func (data *data) maxProfitRecursive(d int, k int, o bool) int {
	// memoized case
	key := fmt.Sprintf("%d %d %d", d, k, o)
	if profit, ok := data.memoized[key]; ok {
		return profit
	}

	n := len(data.prices) - 1
	// base case
	if n == d {
		if o == true && k > 0 { // we have something to sell
			return data.prices[d]
		}
		return 0
	}

	// recurrent case
	// we add each previous day - either price is lower so we should buy
	notSellNotBuy, notSellBuy, sellNotBuy := 0, 0, 0
	// not sell + not buy
	notSellNotBuy = data.maxProfitRecursive(d+1, k, o)
	// not sell + buy
	if o == false {
		notSellBuy = data.maxProfitRecursive(d+1, k, true) - data.prices[d]
	}
	// sell + not buy
	if o == true && k > 0 {
		sellNotBuy = data.prices[d] + data.maxProfitRecursive(d+1, k-1, false)
	}
	// sell + buy (doesn't ever make sense - should just hold?)

	maxProfit := notSellNotBuy
	if notSellBuy > maxProfit {
		maxProfit = notSellBuy
	}
	if sellNotBuy > maxProfit {
		maxProfit = sellNotBuy
	}

	// memoized
	data.memoized[key] = maxProfit

	return maxProfit
}
