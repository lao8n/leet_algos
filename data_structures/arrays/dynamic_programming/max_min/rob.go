package data_structures

// interconnected subproblems combined for a solution -> dynamic programming
// either bottom up tabulation or top-down memoization -> bottom up tabulation
// we could further optimise this to only store the last two numbers
func rob(nums []int) int {
	maxRob := make([]int, len(nums))
	// what would you pick if len of nums was i
	for i := 0; i < len(nums); i++ {
		// base cases
		if i == 0 {
			maxRob[0] = nums[0]
		} else if i == 1 {
			if nums[1] > nums[0] {
				maxRob[1] = nums[1]
			} else {
				maxRob[1] = nums[0]
			}
			// recursive cases
		} else {
			if maxRob[i-2]+nums[i] > maxRob[i-1] {
				maxRob[i] = maxRob[i-2] + nums[i]
			} else {
				maxRob[i] = maxRob[i-1]
			}
		}
	}
	return maxRob[len(maxRob)-1]
}
