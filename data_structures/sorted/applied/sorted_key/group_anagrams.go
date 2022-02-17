import "sort"

func groupAnagrams(strs []string) [][]string {
	// 1. construct a map perhaps with a unique prime decomposition of int -> list of indices
	// 2. [we do this] use a map with sortedString -> index slowly building up slice

	//sortedStrings := string.Map(func(s string) string { return sort.Sort(s) }, strs)
	solution := make([][]string, 0, len(strs)) // this is only the right capacity if all str are unique
	sortedStringSliceIndex := make(map[string]int)

	for _, str := range strs {
		sortedString := sortString(str)
		i, ok := sortedStringSliceIndex[sortedString]
		if !ok { // new sortedString
			j := len(solution)                         // j is index of new slice of anagrams
			solution = append(solution, []string{str}) // add to solution
			sortedStringSliceIndex[sortedString] = j   // add to map
		} else { // sortedString exists
			solution[i] = append(solution[i], str) // add to solution
		}
	}
	return solution
}

func sortString(str string) string {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}