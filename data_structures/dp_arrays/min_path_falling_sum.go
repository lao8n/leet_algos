package data_structures

import "math"

// matrix pathing problems with constraints on movement (1 way) -> dp
// dynamic programming = interdependent problems combined for a solution

// 2 approaches to dp
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation -> good for pathing

// 3 components to dp
// 1. base cases = top row
// 2. recurrence = choose minimum of left, mid and right paths
// 3. state = min cost to get to this square, i, j

// time complexity O(mn)
// space complexity O(n) -> go straight to optimised version as just need prev
// could optimise even further and do in place
func minFallingPathSum(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	prev := make([]int, n)
	for j := 0; j < n; j++ {
		prev[j] = matrix[0][j]
	}
	current := make([]int, n)

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			// choose left, middle or right path
			left, middle, right := math.MaxInt, math.MaxInt, math.MaxInt
			if j > 0 {
				left = prev[j-1]
			}
			if j < n-1 {
				right = prev[j+1]
			}
			middle = prev[j]
			best := middle
			if left < best {
				best = left
			}
			if right < best {
				best = right
			}
			current[j] = best + matrix[i][j]
		}
		prev = current
		current = make([]int, n)
	}
	// find minimum path
	min := math.MaxInt
	for j := 0; j < n; j++ {
		if prev[j] < min {
			min = prev[j]
		}
	}
	return min
}
