package data_structures

import "strconv"

// use a stack
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	var popped1, popped2 int
	for _, t := range tokens {
		endOfStack := len(stack) - 1
		if len(stack) >= 2 {
			popped1 = stack[endOfStack]
			popped2 = stack[endOfStack-1]
		}
		// integer
		push, err := strconv.Atoi(t)
		if err == nil {
			stack = append(stack, push)
			// operator
		} else {
			if t == "+" {
				stack[endOfStack-1] = popped2 + popped1
			} else if t == "-" {
				stack[endOfStack-1] = popped2 - popped1
			} else if t == "*" {
				stack[endOfStack-1] = popped2 * popped1
			} else if t == "/" {
				stack[endOfStack-1] = popped2 / popped1
			}
			stack = stack[:endOfStack]
		}
	}
	return stack[0]
}
