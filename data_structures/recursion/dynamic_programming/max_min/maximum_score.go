package data_structures

// interdependent subproblems combined for a solution -> combined for a solution
// two approaches: 1. [choose] top down recursion with memoization 2. bottom up with tabulation

// need 3 things for dp
// 1. base cases
// 2. recurrence relation
// 3. state definition = left index, right index, multiplier index
// recursion two choices for solution 1. accumulator argument 2. [choose] return
// how to memoize? need something immutable and reduce dimensions from 3 -> 2
func maximumScore(nums []int, multipliers []int) int {
	// m x l array - default to 0
	memoized := make([][]int, len(multipliers))
	for i := range memoized {
		memoized[i] = make([]int, len(nums))
	}
	return maximumScoreDp(nums, multipliers, memoized, 0, 0)
}

func maximumScoreDp(nums []int, multipliers []int, memoized [][]int, m int, l int) int {
	// don't need 3 arguments can calculate r from l & m
	r := len(nums) - 1 - m + l

	// base cases
	if m == len(multipliers) {
		return 0
	}
	// memoized case
	if memoized[m][l] != 0 {
		return memoized[m][l]
	}

	// recurrence cases
	// choose left
	left := maximumScoreDp(nums, multipliers, memoized, m+1, l+1) + nums[l]*multipliers[m]
	// choose right
	right := maximumScoreDp(nums, multipliers, memoized, m+1, l) + nums[r]*multipliers[m]
	if left > right {
		memoized[m][l] = left
		return left
	}
	memoized[m][l] = right
	return right
}
