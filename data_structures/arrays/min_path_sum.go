package data_structures

import "math"

// matrix pathing problems with constraints on movement (1 way) typically dp
// interdependent problems combined for a solution -> dynamic programming

// 2 main approaches to dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 components to dp approach
// base cases
// recurrence = minimum of top and left + value
// state = i, j -> min sum

// time complexity O(mn)
// space complexity O(n)
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	prev := make([]int, n)
	current := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			from_above, from_left := math.MaxInt, math.MaxInt
			if i > 0 {
				from_above = prev[j]
			}
			if j > 0 {
				from_left = current[j-1]
			}
			if i == 0 && j == 0 {
				current[j] = grid[0][0]
			} else if from_above < from_left {
				current[j] = from_above + grid[i][j]
			} else {
				current[j] = from_left + grid[i][j]
			}
		}
		prev = current
		current = make([]int, n)
	}
	return prev[n-1]
}
