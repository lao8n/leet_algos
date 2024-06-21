package data_structures

// clarifying questions - so characters can be used multiple times across words
// 1. initial approach is to take an anagrams angle either 1. map of counts 2. sort - one costs memory
// the other time
// 2.

func countCharacters(words []string, chars string) int {
	cMap := make([]int, 26)
	for _, c := range chars {
		cMap[c-'a'] += 1
	}
	count := 0
	for _, word := range words {
		copyChars := make([]int, 26)
		copy(copyChars, cMap)
		count += checkWord(word, copyChars)
	}
	return count
}

func checkWord(word string, chars []int) int {
	for _, c := range word {
		if chars[c-'a'] <= 0 {
			return 0
		}
		chars[c-'a']--
	}
	return len(word)
}
