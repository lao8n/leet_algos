package data_structures

import "math"

// trick is to do two ways 1. down and right 2. up and left
func updateMatrix(mat [][]int) [][]int {
	// setup dp
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// down and right
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			minDist := math.MaxInt - 1
			if i > 0 && dp[i-1][j] < minDist {
				minDist = dp[i-1][j]
			}
			if j > 0 && dp[i][j-1] < minDist {
				minDist = dp[i][j-1]
			}
			dp[i][j] = minDist + 1
		}
	}

	// up and left
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			minDist := math.MaxInt - 1 // safely add 1 to minDist
			if i < m-1 && dp[i+1][j] < minDist {
				minDist = dp[i+1][j]
			}
			if j < n-1 && dp[i][j+1] < minDist {
				minDist = dp[i][j+1]
			}
			if minDist+1 < dp[i][j] {
				dp[i][j] = minDist + 1
			}

		}
	}
	return dp
}
