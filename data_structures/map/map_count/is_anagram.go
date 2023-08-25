package data_structures

// 1. create map O(m) with number of repetition
//    then go over second string O(n) updating map, and return false if any go negative.
// optimizations
// * min is O(m + n) time, could use array rather than map for memory
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// build map on s
	anagramMap := make(map[rune]int, len(s))
	for _, sc := range s {
		anagramMap[sc]++
	}
	// check other string
	for _, tc := range t {
		if anagramMap[tc] == 0 {
			return false
		}
		anagramMap[tc]--
	}
	return true
}
