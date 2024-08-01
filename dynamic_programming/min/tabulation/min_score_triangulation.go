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
	if j-i < 2 {
		return 0
	}
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	if j-i == 2 {
		res := values[i] * values[i+1] * values[i+2]
		dp[i][j] = res
		return res
	}

	lowestScore := math.MaxInt
	for k := i + 1; k < j; k++ {
		ll := dfs(values, i, k, dp)
		rr := dfs(values, k, j, dp)
		lowestScore = min(lowestScore, ll+rr+values[i]*values[j]*values[k])
	}
	dp[i][j] = lowestScore
	return lowestScore
}
