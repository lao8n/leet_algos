package recursion

// problem : how do we know whether to partition cats as cat or cats?
// brute force recursion : string of length n can be split into two parts n+1 ways, so choice of split or not 2^n, space complexity O(n)
// time complexity O(n^3) - size of recursion tree can go up to n^2,
// space complexity O(n) - depth of recursion tree can go up to n
func wordBreak(s string, wordDict []string) bool {
	memoizedWords := make(map[string]bool)
	return wordBreakRecursive(s, wordDict, memoizedWords)
}

func wordBreakRecursive(s string, wordDict []string, memoizedWords map[string]bool) bool {
	// base case
	if len(s) == 0 {
		return true
	}
	// check if memoizedWords has result already
	if memoizedValue, ok := memoizedWords[s]; ok {
		return memoizedValue
	}
	// recursive case
	for _, word := range wordDict {
		if len(s) >= len(word) && word == s[:len(word)] {
			flag := wordBreakRecursive(s[len(word):], wordDict, memoizedWords)
			if flag {
				memoizedWords[s[len(word):]] = flag
				return true
			}
		}
	}
	memoizedWords[s] = false
	return false
}
