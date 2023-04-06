package recursion

import "fmt"

// interrelated subproblems combined for a solution -> dynamic programming
// two approaches to dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// for a top-down approach 3 key things
// 1. base cases
// 2. recurrence relation - choice between buy, hold, or sell
// 3. state -> args = day index, hold, cooldown, return = profit

// complexity
// time O(n^2) = every time period recurse to every time period
// space = memoized O(n) - number of days * 3 (hold and cooldown)
func maxProfit(prices []int) int {
	s := state{prices: prices, memoized: map[string]int{}}
	return s.maxProfitRecursive(0, false, false)
}

type state struct {
	prices   []int
	memoized map[string]int // represents maximum amount can make from day d (i.e. going backwards)
}

func (s *state) maxProfitRecursive(d int, hold bool, cooldown bool) int {
	// base cases
	if d == len(s.prices) {
		return 0
	}
	// memoized
	key := fmt.Sprintf("%d %d %d", d, hold, cooldown) // cast bool to int
	if p, ok := s.memoized[key]; ok {
		return p
	}

	// recurrent step
	buyProfit, sellProfit, cooldownProfit, holdProfit := 0, 0, 0, 0
	if hold == false && cooldown == false {
		buyProfit = s.maxProfitRecursive(d+1, true, false) - s.prices[d]
	}

	if hold == true && cooldown == false {
		sellProfit = s.maxProfitRecursive(d+1, false, true) + s.prices[d]
	}

	if cooldown == true {
		cooldownProfit = s.maxProfitRecursive(d+1, hold, false)
	}

	// could be holding nothing
	if cooldown == false {
		holdProfit = s.maxProfitRecursive(d+1, hold, false)
	}

	maxProfit := buyProfit
	if sellProfit > maxProfit {
		maxProfit = sellProfit
	}
	if cooldownProfit > maxProfit {
		maxProfit = cooldownProfit
	}
	if holdProfit > maxProfit {
		maxProfit = holdProfit
	}
	// add to memoized
	s.memoized[key] = maxProfit

	return maxProfit
}
