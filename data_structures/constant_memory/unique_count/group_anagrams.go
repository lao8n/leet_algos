package data_structures

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		var table [26]int
		for _, char := range s {
			table[char-'a']++
		}
		m[table] = append(m[table], s)
	}

	var res [][]string
	for _, val := range m {
		res = append(res, val)
	}
	return res
}
