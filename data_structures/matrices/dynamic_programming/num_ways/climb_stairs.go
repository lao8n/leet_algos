package data_structures

// clarifying questions
// approaches
// * dynamic programming - num ways - either top down or bottom up

// specifics
// * can either go forward or backwards. let's go backward to 0th step from nth
// * memoize - bottom up with tabulation
func climbStairs(n int) int {
	// base case
	if n == 1 {
		return 1
	}
	// dp
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	// go backwards
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
