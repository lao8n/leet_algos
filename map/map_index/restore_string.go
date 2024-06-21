package data_structures

// clarifying questions
// * indices - where the character went to

// approaches
// * brute force - find the 0, 1, 2, etc. O(n^2)
// * build a map of each index to letter O(n) space O(2n) time
// * in-place swap - but how swap?
func restoreString(s string, indices []int) string {
	m := make(map[int]byte, len(indices))
	for i, idx := range indices {
		m[idx] = s[i]
	}
	solution := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		solution[i] = m[i]
	}
	return string(solution)
}
