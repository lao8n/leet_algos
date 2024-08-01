package data_structures

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	list1Words := make(map[string]int)
	for i, word := range list1 {
		list1Words[word] = i
	}

	bestWords, bestSum := []string{}, math.MaxInt
	for j, word := range list2 {
		if i, ok := list1Words[word]; ok {
			if i+j < bestSum {
				bestWords = []string{word}
				bestSum = i + j
			} else if i+j == bestSum {
				bestWords = append(bestWords, word)
			}
		}
	}
	return bestWords
}
