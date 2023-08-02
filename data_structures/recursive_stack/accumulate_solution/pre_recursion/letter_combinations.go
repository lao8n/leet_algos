package data_structures

// two problems 1. how acc string 2. how acc solutions
// solution 1: recurse forward, acc string and append to pointer
// solution 2: recurse backwards, acc string
func letterCombinations(digits string) []string {
	combinations := []string{}
	if len(digits) == 0 {
		return combinations
	}
	letterCombinationsAcc(digits, "", &combinations)
	return combinations
}

func letterCombinationsAcc(digits string, accString string, combinations *[]string) {
	// base case
	if digits == "" {
		*combinations = append(*combinations, accString)
		return
	}
	for _, c := range digitLetterMap[digits[0]] {
		letterCombinationsAcc(digits[1:], accString+string(c), combinations)
	}
}

func letterCombinations(digits string) []string {
	return letterCombinationsAcc(digits, "")
}

func letterCombinationsAcc(digits string, accString string) []string {
	// base case
	if digits == "" {
		if accString == "" {
			return []string{}
		} else {
			return []string{accString}
		}
	}
	combinations := []string{}
	for _, c := range digitLetterMap[digits[0]] {
		combinations = append(combinations,
			letterCombinationsAcc(digits[1:], accString+string(c))...)
	}
	return combinations
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
