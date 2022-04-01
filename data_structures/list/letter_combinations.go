// build a map from number -> letter - no clever rule because cannot just add 3 etc.
// total time = time cost of one path * number of paths
// for each combination we need O(N) time and there are O(4^n) combinations
func letterCombinations(digits string) []string {
	combinations := []string{}
	for i, d := range digits { // d = '2'
		newCombinations := []string{}
		if i == 0 {
			combinations = digitLetterMap[d]
		} else {
			for _, l := range digitLetterMap[d] { // l = "a"
				for _, c := range combinations {
					newCombinations = append(newCombinations, c+l)
				}
			}
			combinations = newCombinations
		}
	}
	return combinations
}

var digitLetterMap = map[rune][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}