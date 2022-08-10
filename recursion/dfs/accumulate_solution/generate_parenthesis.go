package recursion

// approach = recursive function
// questions
// 1. how acc solutions? in return function as already have []string
// 2. how acc combinations? in argument
// 3. how avoid repetition?
// 4. problem -> how generate "(())(())" pattern

// general idea - given a combination c there are 3 options
// ()c, (c), ()c
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	result := []string{}
	generateParenthesisRecursive(n, n, "", &result)
	return result
}

func generateParenthesisRecursive(l, r int, accString string, result *[]string) {
	if r < l {
		return
	}
	if l == 0 && r == 0 {
		*result = append(*result, accString)
	}
	if l > 0 {
		generateParenthesisRecursive(l-1, r, accString+"(", result)
	}
	if r > 0 {
		generateParenthesisRecursive(l, r-1, accString+")", result)
	}
}
