package data_structures

import "strings"

// approaches
// * have a map with counts then loop over map
func uncommonFromSentences(s1 string, s2 string) []string {
	// count occurrences
	m := make(map[string]int)
	for _, word := range strings.Split(s1, " ") {
		m[word]++
	}
	for _, word := range strings.Split(s2, " ") {
		m[word]++
	}
	// find single occurrences
	result := []string{}
	for word, count := range m {
		if count == 1 {
			result = append(result, word)
		}
	}
	return result
}
