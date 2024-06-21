package data_structures

import "unicode"

// approaches
// * use stack where push if addition/subtraction (just negate if latter) and pop if * or /
// * actually just need two numbers - last number and rolling sum
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	cur, last, result := 0, 0, 0
	sign := '+'
	for i, r := range s {
		if unicode.IsDigit(r) {
			cur = (cur * 10) + int(r-'0')
		}
		if !unicode.IsDigit(r) && r != ' ' || i == len(s)-1 {
			if sign == '+' {
				result += last
				last = cur
			} else if sign == '-' {
				result += last
				last = -cur
			} else if sign == '*' {
				last *= cur
			} else if sign == '/' {
				last /= cur
			}
			sign = r
			cur = 0
		}
	}
	result += last
	return result
}
