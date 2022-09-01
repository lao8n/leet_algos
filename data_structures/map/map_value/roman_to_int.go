package data_structures

func romanToInt(s string) int {
	int_to_roman_map := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	sum, last := 0, 0
	for _, r := range s {
		i := int_to_roman_map[r]
		if last < i {
			sum -= last * 2
		}
		sum += i
		last = i
	}
	return sum
}
