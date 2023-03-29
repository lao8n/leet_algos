package data_structures

import "sort"

// 1. create map O(m) with number of repetition
//    then go over second string O(n) updating map, and return false if any go negative.
// optimizations
// * min is O(m + n) time, could use array rather than map for memory
// 2. sort costing O(n logn)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// sort - 2 * O(n logn)
	// string -> []rune -> sort (no need -> string)
	sSlice, tSlice := []rune(s), []rune(t)
	sort.Slice(sSlice, func(i, j int) bool { return sSlice[i] < sSlice[j] })
	sort.Slice(tSlice, func(i, j int) bool { return tSlice[i] < tSlice[j] })

	// compare - O(n)
	for i := range sSlice {
		if sSlice[i] != tSlice[i] {
			return false
		}
	}
	return true
}
