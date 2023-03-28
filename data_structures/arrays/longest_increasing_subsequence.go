package data_structures

// interdependent subproblems combined for a solution -> dynamic programming
// 2 approaches ot dynamic programming
// 1. top-down recursion with memoization
// 2. bottom-up iteration with tabulation

// recurrence state has 3 components
// 1. base cases
// 2. recurrence step
// 3. state variables -> index of start of sequence (go from back) which returns max length

// complexity
// time = O(n^2)
// space = O(n)
func lengthOfLIS(nums []int) int {
	// length of longest sequence ending with index i
	tabulation := make([]int, len(nums))

	longest := 0
	for j := 0; j < len(nums); j++ {
		tabulation[j] = 1 // can always have seq of at least 1
		for i := 0; i < j; i++ {
			if nums[j] > nums[i] {
				if tabulation[i]+1 > tabulation[j] {
					tabulation[j] = tabulation[i] + 1
				}
			}
		}
		if tabulation[j] > longest {
			longest = tabulation[j]
		}
	}

	return longest
}
