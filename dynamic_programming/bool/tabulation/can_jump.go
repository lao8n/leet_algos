package data_structures

// approaches
// * dynamic programming - interconnected problems - either top down with memoization or bottom up with memoization.. - tabulation is probably easier here

// specifics
// * can go forward or backwards.. given nature of data probably easier to go forward
func canJump(nums []int) bool {
	// edge case
	if len(nums) == 0 {
		return false
	}

	dp := make([]bool, len(nums))
	dp[0] = true // can always get to first step
	for i, num := range nums {
		if dp[i] { // we've been able to get to this step
			for j := 1; j <= num; j++ {
				if i+j < len(nums) {
					dp[i+j] = true
				}
			}
		}
	}

	return dp[len(nums)-1]
}
