package data_structures

// clarifying questions
// * order - sorted
// * if has b should all bs be put there? probably yes
// approaches
// * create map of s, loop over order, remove from map
func customSortString(order string, s string) string {
	sMap := make(map[rune]int)
	for _, c := range s {
		sMap[c]++
	}

	result := make([]rune, 0)
	for _, c := range order {
		if count, ok := sMap[c]; ok {
			for i := 0; i < count; i++ {
				result = append(result, c)
			}
			delete(sMap, c) // order is unique
		}
	}
	// any remaining in map
	for c, count := range sMap {
		for i := 0; i < count; i++ {
			result = append(result, c)
		}
	}
	return string(result)
}
