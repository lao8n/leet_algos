package data_structures

// path constraint -> dp rather than recursion
// interdependent problems combined for a solution -> dynamic programming

// dynamic programming has 2 approaches
// 1. top-down recursion with memoization
// 2. bottom-up iteration with tabulation -> typically easier for pathing problems

// 3 components to dynamic programming
// 1. base cases
// 2. recurrence = if object is there then set to 0
// 3. state = i,j

// time complexity O(mn)
// space complexity O(mn) -> could be optimised for O(n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	tabulation := make([][]int, m)
	for i := 0; i < m; i++ {
		tabulation[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			numWays := 0
			if obstacleGrid[i][j] == 0 {
				if i == 0 && j == 0 {
					numWays = 1
				}
				if i > 0 {
					numWays += tabulation[i-1][j]
				}
				if j > 0 {
					numWays += tabulation[i][j-1]
				}
			}
			tabulation[i][j] = numWays
		}
	}
	return tabulation[m-1][n-1]
}
