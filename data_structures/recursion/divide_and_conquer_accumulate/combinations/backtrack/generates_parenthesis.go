// approaches
// * use ( + recurseLeft + ) + recurseRight
// * acc solution with left or right bracket - return nothing if doesn't work
func generateParenthesis(n int) []string {
	d := data{
		n:        n,
		solution: []string{},
	}
	d.backtrack(0, 0, "")
	return d.solution
}

type data struct {
	n        int
	solution []string
}

func (d *data) backtrack(l int, r int, acc string) {
	// base cases
	if len(acc) == 2*d.n {
		d.solution = append(d.solution, acc)
		return
	}
	// recursive cases
	if l < d.n {
		// acc += "("
		d.backtrack(l+1, r, acc+"(")
		// acc = acc[:len(acc) - 1]
	}
	if r < l {
		// acc += ")"
		d.backtrack(l, r+1, acc+")")
		// acc = acc[:len(acc) - 1]
	}
	return
}
