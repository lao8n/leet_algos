package data_structures

import "math"

// interdependent subproblems combined for a solution -> dynamic programming
// 2 approaches to dynamic programming:
// 1. top-down recursion with memoization
// 2. bottom-up iteration with tabulation

// go recursion because typically easier
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

// complexity
// time input parameters for recursive function are two strings with M possible text1 and N possible text2 which means MN
// solving each problems takes O(1) where we no longer need to search the full string
// therefore time complexity is O(MN)
// space complexity = need to store solution for O(MN) problems
func longestCommonSubsequence(text1 string, text2 string) int {
	memoized := make(map[string]map[string]int)
	return longestCommonSubsequenceRecursive(text1, text2, memoized)
}

func longestCommonSubsequenceRecursive(text1 string, text2 string, memoized map[string]map[string]int) int {
	// base case
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	// memoized case
	if memoizedLongest, ok := memoized[text1][text2]; ok {
		return memoizedLongest
	}

	// recursive case
	longest := 0
	if text1[0] == text2[0] {
		longest = longestCommonSubsequenceRecursive(text1[1:], text2[1:], memoized) + 1
	} else {
		longest = int(math.Max(
			float64(longestCommonSubsequenceRecursive(text1[1:], text2, memoized)),
			float64(longestCommonSubsequenceRecursive(text1, text2[1:], memoized)),
		))
	}
	if _, ok := memoized[text1]; ok {
		memoized[text1][text2] = longest
	} else {
		memoized[text1] = make(map[string]int)
		memoized[text1][text2] = longest
	}
	return longest
}
