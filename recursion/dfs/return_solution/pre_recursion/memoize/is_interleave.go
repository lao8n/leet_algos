package recursion

import "fmt"

// interconnected problems -> combined for a solution -> dynamic programming

// 2 approaches to dp
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 components to dp
// 1. base cases
// 2. recurrence
// 3. state -> index in s1 and index in s2

// time complexity O(m x n)
// space complexity O(m x n)
func isInterleave(s1 string, s2 string, s3 string) bool {
	// base cases
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	memoized := make(map[string]bool, len(s1)*len(s2))
	result := isInterleaveRecursive(memoized, s1, s2, s3)
	fmt.Println(memoized)
	return result
}

func isInterleaveRecursive(memoized map[string]bool, s1 string, s2 string, s3 string) bool {
	// memoized
	key := s1 + " " + s2
	if interleaved, ok := memoized[key]; ok {
		return interleaved
	}

	// base cases
	if len(s1) == 0 {
		memoized[key] = s2 == s3
		return s2 == s3
	} else if len(s2) == 0 {
		memoized[key] = s1 == s3
		return s1 == s3
	}

	// recursive
	chooseLeft, chooseRight := false, false
	if s1[0] == s3[0] {
		chooseLeft = isInterleaveRecursive(memoized, s1[1:], s2, s3[1:])
	}
	if s2[0] == s3[0] {
		chooseRight = isInterleaveRecursive(memoized, s1, s2[1:], s3[1:])
	}

	memoized[key] = chooseLeft || chooseRight

	return chooseLeft || chooseRight
}
