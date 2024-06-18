package data_structures

import "strings"

// approaches
// * hold two mappings from letter to word and word to letter - see if it breaks
// * s already all in memory so might as well keep it that way
func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")

	// check lens
	if len(sList) != len(pattern) {
		return false
	}
	// create maps
	pMap, sMap := make(map[rune]string), make(map[string]rune)
	// loop
	for i, c := range pattern {
		if word, ok := pMap[c]; ok {
			// c already exists
			if word != sList[i] {
				return false
			}
		} else {
			pMap[c] = sList[i]
		}

		if letter, ok := sMap[sList[i]]; ok {
			if c != letter {
				return false
			}
		} else {
			sMap[sList[i]] = c
		}
	}
	return true
}
