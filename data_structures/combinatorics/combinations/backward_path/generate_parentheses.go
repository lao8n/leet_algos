package data_structures

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	solution := []string{}
	for l := 0; l < n; l++ {
		for _, ls := range generateParenthesis(l) {
			for _, rs := range generateParenthesis(n - 1 - l) {
				solution = append(solution, "(" + ls + ")" + rs)
			}
		}
	}
	return solution
}