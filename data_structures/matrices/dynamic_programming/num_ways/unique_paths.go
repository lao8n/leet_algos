package data_structures

// interdependent problems combined for a solution -> dynamic programming
// path constraint -> dp rather than recursion

// dynamic programming has 2 approaches
// 1. top-down recursion with memoization
// 2. bottom up iteration with tabulation -> for pathing typically easier

// time complexity O(mn)
// space complexity O(mn) -> can optimise to O(n)
func uniquePaths(m int, n int) int {
	tabulation := make([][]int, m)
	for i, _ := range tabulation {
		tabulation[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num_ways := 0
			if i == 0 && j == 0 {
				num_ways = 1
			}
			if i > 0 {
				num_ways += tabulation[i-1][j]
			}
			if j > 0 {
				num_ways += tabulation[i][j-1]
			}
			tabulation[i][j] = num_ways
		}
	}
	return tabulation[m-1][n-1]
}
