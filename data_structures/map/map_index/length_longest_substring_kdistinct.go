package data_structures

// approaches
// * use a map from character to index
// * two pointers left and right and map to track number of distinct elements

// specifics
// * what to do when > k distinct characters? have to jump to +1 of latest location of that letter.. for new starting point

// counter-example
// * abaccc - valid answer is accc for length 4 with k = 2 - rather than deleting all elements we should only delete enough to get len(distinctChars) <= k
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	start := 0
	distinctChars := make(map[byte]int, 0)
	maxLength := 0
	for finish := 0; finish < len(s); finish++ {
		c := s[finish]
		// c is old character
		if _, ok := distinctChars[c]; ok {
			// seen character before need to update index
			distinctChars[c] = finish
		} else {
			// c is new character
			distinctChars[c] = finish
			if len(distinctChars) > k {
				// need to increment start
				startChar := s[start]
				startCharLatestIndex := distinctChars[startChar]
				j := start
				for j <= startCharLatestIndex {
					if j == distinctChars[s[j]] { // we have found last example of this element
						delete(distinctChars, s[j])
					}
					if len(distinctChars) <= k {
						break
					}
					j++
				}
				start = j + 1
			}
		}
		if finish+1-start > maxLength {
			maxLength = finish + 1 - start
		}

	}
	return maxLength
}
