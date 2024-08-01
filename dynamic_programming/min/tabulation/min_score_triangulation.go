package data_structures

import (
	"fmt"
	"math"
)

func minScoreTriangulation(values []int) int {
	dp := make([][]int, len(values))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(values))
	}
	x := dfs(values, 0, len(values)-1, dp) // i & j are two neighbours
	fmt.Println(dp)
	return x
}

func dfs(values []int, i, j int, dp [][]int) int {
	// base case
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
	// recursive case
	lowestScore := math.MaxInt
	for k := i + 1; k < j; k++ {
		// try having a triangle with k
		ll := dfs(values, i, k, dp)
		rr := dfs(values, k, j, dp)
		lowestScore = min(lowestScore, ll+rr+values[i]*values[j]*values[k])
	}
	dp[i][j] = lowestScore
	return lowestScore
}
