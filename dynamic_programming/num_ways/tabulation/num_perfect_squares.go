package data_structures

// specifics
// * go forward or go backwards
// * optimize by starting with larger numbers? i.e. only consider a subset of options?
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	// sqrt n
	for s := 1; s*s <= n; s++ {
		for i := 0; i+s*s <= n; i++ {
			if i != 0 && dp[i] == 0 {
				continue
			}
			// is there a way to get there?
			newIndex := i + s*s
			if dp[newIndex] == 0 || dp[i]+1 < dp[newIndex] {
				dp[newIndex] = dp[i] + 1
			}
		}
	}
	return dp[n]
}
