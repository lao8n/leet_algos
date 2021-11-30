func firstUniqChar(s string) int {
	// naive for loop trying every letter
	// map = use a map of all the characters and their count
	// slice = even faster
	charCounts := make([]int, 26)
	for _, v := range s {
		charCounts[v-'a']++
	}
	for i, v := range s {
		if charCounts[v-'a'] == 1 {
			return i
		}
	}
	return -1
}