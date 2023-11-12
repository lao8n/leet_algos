package data_structures

// clarifying questions
// * () = 1 - base case by itself
// * (A) = 2A wrapper
// * AB = A + B
// will always be balanced? ((()())

// approaches
// * just loop through string problem is if have () and need to double how handle that?
// * push to stack with the score

// specifics
// * once on stack do we care about what brackets are? use as that is how track whether closing
// * see ) then if previous was left bracket we pop and add to score 1
func scoreOfParentheses(s string) int {
	// stack for tracking brackets
	stack := []int{0}
	// loop to process through brackets
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, 0)
		} else if s[i] == ')' {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] += max(2*popped, 1)
		}
	}
	return stack[len(stack)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
