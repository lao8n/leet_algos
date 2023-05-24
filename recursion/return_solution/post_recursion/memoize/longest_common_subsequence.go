package recursion

// interdependent subproblems combined for a solution -> dynamic programming
// 2 approaches to dynamic programming:
// 1. top-down recursion with memoization
// 2. bottom-up iteration with tabulation

// go recursion because typically easier
// need 3 components
// 1. base cases
// 2. recursive step
// 3. state variables

// recursive step: do we include this letter which increments subsequence by 1 but makes us jump in the
// other sequence
// state variable: left seq index, right seq index, don't need length because can return that - or just subsequence strings
// question: unsure whether we should remove string to the right or not?
// memoize? not sure yet, need something immutable maybe concat of two strings? better to do nested map

// complexity
// time input parameters for recursive function are two strings with M possible text1 and N possible text2 which means MN
// solving each problems takes O(N) where we search for a character in a string of length N
// therefore time complexity is O(MN^2)
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
	// memoize
	if text1Map, ok := memoized[text1]; ok {
		if longest, ok := text1Map[text2]; ok {
			return longest
		}
	}

	// recursive step
	// match left
	longest := 0
	for i, t2 := range text2 {
		if rune(text1[0]) == t2 {
			longest = longestCommonSubsequenceRecursive(text1[1:], text2[i+1:], memoized) + 1
			break
		}
	}
	// do not need to match right as will do so if match not left
	// don't match left
	longestNotLeft := longestCommonSubsequenceRecursive(text1[1:], text2, memoized)
	if longestNotLeft > longest {
		longest = longestNotLeft
	}
	if _, ok := memoized[text1]; ok {
		// text1Map exists just not text2
		memoized[text1][text2] = longest
	} else {
		memoized[text1] = make(map[string]int)
		memoized[text1][text2] = longest
	}
	return longest
}
