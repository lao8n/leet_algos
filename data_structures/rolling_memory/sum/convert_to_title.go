package data_structures

// approaches
// * use rolling sum which we decrement with letter
// * choice go from smallest to largest -> don't know how large it could be..

// specifics
// * use ascii - rather than map
// * could use string builder
func convertToTitle(columnNumber int) string {
	// init
	title := ""
	// for loop decrementing
	for columnNumber > 0 {
		// remainder
		rem := (columnNumber - 1) % 26 // 25 for ZY
		// append to string
		title = string(int('A')+rem) + title

		// update columnNumber
		columnNumber = (columnNumber - rem) / 26
	}
	return title
}
