package data_structures

import "math"

// interdependent subproblems combined for a solution -> dynamic programming
// 2 approaches to dynamic programming:
// 1. top-down recursion with memoization
// 2. bottom-up iteration with tabulation

// go bottom up iteration as it is faster
// need 3 components
// 1. base cases
// 2. recursive step
// 3. state variables

// state variable: left seq index, right seq index, don't need length because can return that - or just subsequence strings
// we can do better however
// there are 4 scenarios
// 1. text1 first letter included only
// 2. text2 first letter included only
// 3. text1 and text2 first letters both included
// 4. neither included

// in the previous recursion with memoization approach
// * we handled 2 & 4 together - i.e. text1 not included
// * and handled 1 & 3 together where we loop to find which of text 2 is included

// an alternative solution would be to
// * handle 3
// * and recurse for skipping text1 i.e. 2 & 4 and then recurse for skipping text2 i.e. 1 & 4 taking max

func longestCommonSubsequence(text1 string, text2 string) int {
	// initialize tabulation
	tabulation := make([][]int, len(text1)+1)
	for i := range tabulation {
		tabulation[i] = make([]int, len(text2)+1)
	}

	// build table
	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			// two cases
			// if both the same + 1
			if text1[i] == text2[j] {
				tabulation[i][j] = tabulation[i+1][j+1] + 1
			} else {
				// otherwise choose maximum of below and right
				tabulation[i][j] = int(math.Max(float64(tabulation[i+1][j]), float64(tabulation[i][j+1])))
			}
		}
	}
	return tabulation[0][0]
}
