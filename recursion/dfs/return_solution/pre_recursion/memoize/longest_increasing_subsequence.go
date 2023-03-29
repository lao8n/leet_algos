package data_structures

// interdependent subproblems combined for a solution -> dynamic programming
// 2 approaches ot dynamic programming
// 1. top-down recursion with memoization
// 2. bottom-up iteration with tabulation

// recurrence state has 3 components
// 1. base cases
// 2. recurrence step
// 3. state variables

// complexity
// time = O(n^2)
// space = O(n)
func lengthOfLIS(nums []int) int {
	// longest seq ending at index j
	memoized := make([]int, len(nums))
	return recursiveLengthOfLIS(nums, len(nums)-1, 0, memoized)
}

func recursiveLengthOfLIS(nums []int, endsAt int, longest int, memoized []int) int {
	// base case
	if endsAt == 0 {
		memoized[0] = 1
		return 1
	}
	// recurrent case - given an ending index j try every i
	memoized[endsAt] = 1 // can always have length 1
	longest = recursiveLengthOfLIS(nums, endsAt-1, longest, memoized)
	for i := endsAt; i >= 0; i-- {
		if nums[endsAt] > nums[i] {
			if memoized[i]+1 > memoized[endsAt] {
				memoized[endsAt] = memoized[i] + 1
			}
		}
	}
	if memoized[endsAt] > longest {
		longest = memoized[endsAt]
	}
	return longest
}
