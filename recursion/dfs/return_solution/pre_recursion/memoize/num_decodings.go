package recursion

// interdependent subproblems combined for a solution -> dynamic programming
// other hint is it is num of ways

// dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 components
// 1. base cases
// 2. recurrence
// 3. state = string, 2 digit key flag -> number of ways - recursing forward but return solution - state is more complicated as need to check previous number as 1 can have up to 9 but 2 only up to 6
import "strconv"

func numDecodings(s string) int {
	memoized := make([]int, len(s))
	return recursiveNumDecodings(memoized, s, 0)
}

func recursiveNumDecodings(memoized []int, s string, index int) int {
	// base cases
	if index == len(s) {
		return 1
	}
	// memoized
	if memoized[index] != 0 {
		return memoized[index]
	}

	// recursive case
	// one digit
	d, _ := strconv.Atoi(string(s[index]))
	numWays := 0
	if d > 0 {
		numWays += recursiveNumDecodings(memoized, s, index+1)
	}
	// 2 digit
	if index < len(s)-1 {
		dNext, _ := strconv.Atoi(string(s[index+1]))
		if d == 1 || (d == 2 && dNext <= 6) {
			numWays += recursiveNumDecodings(memoized, s, index+2)
		}
	}
	memoized[index] = numWays

	return numWays
}
