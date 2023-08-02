package data_structures

import "fmt"

// interrelated subproblems combined for a solution -> dynamic programming
// 2 approaches to dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 things need to consider
// 1. state variables <- num of posts with last colour, num of colours (doesn't change), num posts remaining
// 2. recurrence relation
// 3. base cases

// because we are using recursion we will return the output (not accumulate), and build from end
type nwState struct {
	k        int
	memoized map[string]int
}

func numWays(n int, k int) int {
	s := nwState{
		k:        k,
		memoized: map[string]int{},
	}
	return s.recursiveNumWays(n, 1)
}

func (s *nwState) recursiveNumWays(nRemain, nSameColour int) int {
	// base cases
	if nRemain == 0 {
		return 1
	}
	if nSameColour == 3 {
		return 0
	}

	// memoized
	key := fmt.Sprintf("%d %d", nRemain, nSameColour)
	if numWays, ok := s.memoized[key]; ok {
		return numWays
	}

	// recursive case - choice between same colour & different
	diffColourNumWays := (s.k - 1) * s.recursiveNumWays(nRemain-1, 1)
	sameColourNumWays := s.recursiveNumWays(nRemain-1, nSameColour+1)
	totalNumWays := diffColourNumWays + sameColourNumWays

	s.memoized[key] = totalNumWays

	return totalNumWays
}
