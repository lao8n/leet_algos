package data_structures

// recurse from beginning or end
// still have double for loop here
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	// base case
	if len(digits) == 1 {
		return digitLetterMap[digits[0]]
	}
	// recursive case
	lastElement := len(digits) - 1
	combinations := letterCombinations(digits[:lastElement])
	var newCombinations []string
	for _, l := range digitLetterMap[digits[lastElement]] { // [a, b, c]
		for _, c := range combinations {
			newCombinations = append(newCombinations, c+l)
		}
	}
	return newCombinations
}

var digitLetterMap = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}
