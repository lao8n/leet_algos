package data_structures

import "math"

// interconnected problems combined for a solution -> dynamic programming
// dp implementation choice
// * top down recursion with memoization -> maybe some sort of connected components approach?
// * bottom up iteration with tabulation

// there are 3 components to a dp approach
// 1. base case
// 2. recursive step
// 3. state variable

// 4 options for a recursive step - consider an element does it make the area bigger?
// if it a zero no, it is a one yes - maybe
// if row of 1s then yes will add
// if it is a column of 1s below will add

func maximalSquare(matrix [][]byte) int {
	// initialize tabulation - time O(mn) space O(mn)
	tabulation := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix)+1; i++ {
		tabulation[i] = make([]int, len(matrix[0])+1)
	}

	// recursive step
	maxSquareLength := 0
	for i := 1; i < len(tabulation); i++ {
		for j := 1; j < len(tabulation[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				tabulation[i][j] = int(math.Min(float64(tabulation[i-1][j-1]), math.Min(float64(tabulation[i-1][j]), float64(tabulation[i][j-1])))) + 1
				if tabulation[i][j] > maxSquareLength {
					maxSquareLength = tabulation[i][j]
				}
			}
		}
	}

	return maxSquareLength * maxSquareLength
}
