package data_structures

func letterCombinations(digits string) []string {
	// base case
	if len(digits) == 0 {
		return []string{}
	}

	digitsMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	d := data{
		digits:       digits,
		digitsMap:    digitsMap,
		combinations: []string{},
	}
	d.backtrack(0, "")
	return d.combinations
}

type data struct {
	digits       string
	digitsMap    map[string]string
	combinations []string
}

func (d *data) backtrack(i int, path string) {
	// base case
	if len(path) == len(d.digits) {
		d.combinations = append(d.combinations, path)
		return
	}
	// recursive case
	digit := d.digits[i]
	for _, char := range d.digitsMap[string(digit)] {
		path += string(char)
		d.backtrack(i+1, path)
		path = path[:len(path)-1]
	}
}
