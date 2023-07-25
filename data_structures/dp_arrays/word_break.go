package data_structures

// dp - two pointers i = len of substring, j = partition index
// complexity - O(n^3) two for loops O(n^2) and substring computation is O(n), space complexity O(n) length of dp array
func wordBreak(s string, wordDict []string) bool {
	// create dictionary
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}

	// dp array
	memoized := make([]bool, len(s)+1)
	memoized[0] = true
	for lenI := 1; lenI <= len(s); lenI++ {
		for partI := 0; partI < lenI; partI++ {
			if memoized[partI] && wordMap[s[partI:lenI]] {
				memoized[lenI] = true
				break // to next i loop
			}
		}
	}
	return memoized[len(s)]
}
