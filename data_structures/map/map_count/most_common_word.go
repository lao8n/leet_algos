package data_structures

import (
	"strings"
	"unicode"
)

// 1. for loop with counts
// 2. however then need to find largest count which isn't banned, could sort or heap
// easier approach
// 1. make banned a set and only push store counts if word not banned
// 2. another for loop to get highest - could do it in a single pass as well
func mostCommonWord(paragraph string, banned []string) string {
	// paragraph preprocessing
	ps := paragraphSlice(paragraph)
	// banned map
	bannedMap := make(map[string]bool)
	for _, b := range banned {
		bannedMap[strings.ToLower(b)] = true
	}

	// freq map
	wordCount := make(map[string]int)
	for _, w := range ps {
		if !bannedMap[w] {
			wordCount[w] += 1
		}
	}

	// get highest
	count, mcWord := 0, ""
	for w, c := range wordCount {
		if c > count {
			mcWord, count = w, c
		}
	}

	return mcWord
}

func paragraphSlice(paragraph string) []string {
	ps := []string{}
	var word strings.Builder
	for _, c := range paragraph {
		if unicode.IsLetter(c) {
			word.WriteRune(unicode.ToLower(c))
		} else if word.Len() > 0 {
			ps = append(ps, word.String())
			word.Reset()
		}
	}
	if word.Len() > 0 {
		ps = append(ps, word.String())
	}
	return ps
}

// func isLetter(c rune) bool {
//     return c != '!' && c != '.' && c != ',' && c != ' ' && c != '?' && c != '\'' && c != ';'
// }
