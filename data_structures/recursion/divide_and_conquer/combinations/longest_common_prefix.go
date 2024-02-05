package data_structures

// clarifying questions?
// * prefix needs to be true for all strs?
// * it isn't a substring but a prefix? yes
// * how should i handle return "" in valid case of len(0) string - return 1 first
// approaches
// * just compare all strs character by character - 1st character loop of n until O(n x p) where p length of longest common prefix
// * instead we can simplify with take the intersection of each pair of strings which is O(n)
// * could have also done divide and conquer
func longestCommonPrefix(strs []string) string {
	// divide
	// base cases
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 0 {
		return ""
	}
	// recursive case
	left := longestCommonPrefix(strs[:len(strs)/2])  // strs[:4] exclusive
	right := longestCommonPrefix(strs[len(strs)/2:]) // strs[4:] inclusive
	// conquer - build
	intersection := ""
	for i := 0; i < len(left) && i < len(right); i++ {
		if left[i] == right[i] {
			intersection += string(left[i])
		} else {
			break
		}
	}
	return intersection
}
