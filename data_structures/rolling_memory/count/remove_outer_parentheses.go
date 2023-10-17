package data_structures

import "strings"

// clarifying questions
// * remove outer parentheses
// * guaranteed valid? yes

// approaches
// * stack -> build stack of brackets - stack would be with a count for outer layer
// * string manipulation -> can always take the outer brackets? and then maybe count left and right i.e. have counts of num left and num right and ignore if num right or num left = 0 or something
func removeOuterParentheses(s string) string {
	lCount := 0
	var solution strings.Builder
	// loop over string
	for _, c := range s {
		if c == '(' {
			if lCount > 0 {
				solution.WriteRune(c)
			}
			lCount++
		} else if c == ')' {
			lCount--
			if lCount > 0 {
				solution.WriteRune(c)
			}
		}
		//fmt.Printf("%c %d %d\n", c, lCount, rCount)
	}
	return solution.String()
}
