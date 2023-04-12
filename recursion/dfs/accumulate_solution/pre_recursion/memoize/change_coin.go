package recursion

import "fmt"

// interdependent subproblems combined for a solution -> dynamic programming
// dynamic programming has 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// try top down - 3 components
// 1. base cases
// 2. recurrence relation - can add c coin amount to x amounts
// 3. state = amount

// how to handle duplicates - if use smaller coin cannot use larger coins anymore
func change(amount int, coins []int) int {
	s := state{memoized: make(map[string]int)}
	return s.changeRecursive(amount, coins)
}

type state struct {
	memoized map[string]int
}

func (s *state) changeRecursive(amount int, coins []int) int {
	// base cases
	if amount < 0 {
		return 0
	} else if amount == 0 {
		return 1
	}
	// memoized
	key := fmt.Sprintf("%d %d", len(coins), amount)
	if memAmount, ok := s.memoized[key]; ok {
		return memAmount
	}

	// recurrence case
	numWays := 0
	for i, c := range coins {
		numWays += s.changeRecursive(amount-c, coins[i:])
	}

	s.memoized[key] = numWays

	return numWays
}
