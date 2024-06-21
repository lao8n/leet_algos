package data_structures

func titleToNumber(columnTitle string) int {
	// process from right to left multiplying by 26
	multiple26 := 1
	column := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		letterNum := int(columnTitle[i]-'A') + 1
		column += multiple26 * letterNum // A is 1 not 0
		multiple26 *= 26
	}
	return column
}
