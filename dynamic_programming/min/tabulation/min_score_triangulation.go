package data_structures

import "math"

func minScoreTriangulation(values []int) int {
	dp := make([][]int, len(values)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(values)+1)
	}

	return dfs(values, 0, len(values)-1, dp)
}

func dfs(values []int, i, j int, dp [][]int) int {
	// base case
	if j-1 < 2 {
		return 0
	}
	// memoized case
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	if j-i == 2 {
		res := values[i] * values[i+1] * values[i+2]
		dp[i][j] = res
		return res
	}
	// recursive case
	lowestScore := math.MaxInt
	for k := i + 1; k < j; k++ {
		ll := dfs(values, i, k, dp)
		rr := dfs(values, k, j, dp)
		score := ll + rr + values[i]*values[j]*values[k]
		if score < lowestScore {
			lowestScore = score
		}
	}
	dp[i][j] = lowestScore
	return lowestScore
}
