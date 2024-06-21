package data_structures

// approaches
// * go right to left - but doesn't work for first digit
// * push characters onto a stack -> pop last character if get a digit
// specifics
// * look forward or go backwards
func clearDigits(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if isNum(s[i]) {
			// don't add num and pop from stack
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	result := ""
	for i := 0; i < len(stack); i++ {
		result += string(stack[i])
	}
	return result
}

func isNum(c byte) bool {
	return c-'0' <= 9 && c-'0' >= 0
}
