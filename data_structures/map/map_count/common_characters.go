package data_structures

// approaches
// * map with counts per word - recursive intersection approach otherwise O(n)
func commonChars(words []string) []string {
	acc := make(map[rune]int)
	for _, c := range words[0] {
		acc[c]++
	}
	for i := 1; i < len(words); i++ {
		intersection(acc, words[i])
	}
	result := []string{}
	for c, count := range acc {
		for i := 0; i < count; i++ {
			result = append(result, string(c))
		}
	}
	return result
}

func intersection(acc map[rune]int, word2 string) {
	w2Map := make(map[rune]int)
	for _, c := range word2 {
		w2Map[c]++
	}
	for c, count := range acc {
		newCount := min(count, w2Map[c])
		if newCount == 0 {
			delete(acc, c)
		}
		if newCount != count {
			acc[c] = newCount
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
