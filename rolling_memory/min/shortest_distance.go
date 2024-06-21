package data_structures

// clarifying questions
// * distance is symmetric doesn't matter whether word1 or word2 is first?
// * guaranteed to be in the list?
// approaches
// * greedily store index1 index2, update with shorter if newer - should be fine as going left to right so can only get better?
// * could go through with map and indices and check every combination
func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	i1, i2 := -1, -1
	shortestDistance := len(wordsDict)
	// loop over words dict
	for i, word := range wordsDict {
		if word == word1 {
			i1 = i
			if i1 != -1 && i2 != -1 && distance(i, i2) < shortestDistance {
				shortestDistance = distance(i, i2)
			}
		}
		if word == word2 {
			i2 = i
			if i1 != -1 && i2 != -1 && distance(i1, i) < shortestDistance {
				shortestDistance = distance(i1, i)
			}
		}
	}
	return shortestDistance
}

func distance(i, j int) int {
	if i > j {
		return i - j
	}
	return j - i
}
