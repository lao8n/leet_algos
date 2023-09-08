package data_structures

// dfs by 1. recursion or 2. stack
func findTargetSumWays(nums []int, target int) int {
	// base case
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}

	// recursive case
	positiveWays := findTargetSumWays(nums[1:], target-nums[0]) // adding reduces target
	negativeWays := findTargetSumWays(nums[1:], target+nums[0]) // subtracting increases target
	return positiveWays + negativeWays
}
