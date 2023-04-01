package data_structures

// backpropagation is essentially dynamic programming where we tabulate answers backwards filling in the table
// not sure if this solution works but the implementation is based upon the pseudo-code from bengio's deep learning
// textbook p206-207

type operation struct {
	function   func([]int) int
	derivative func(int) int
}

func (o operation) forward_pass(x []int, n int) int {
	ni := len(x)
	u := make([]int, n)
	for i := 0; i < ni; i++ {
		u[i] = x[i]
	}
	for i := ni; i < n; i++ {
		parents := u[:i] // really weird parent function
		u[i] = o.function(parents)
	}
	return u[n-1]
}

func (o operation) backward_pass(x []int, n int) []int {
	ni := len(x)
	grad_table := make([]int, n) // du_n/du_i
	for i := 0; i < n; i++ {
		grad_table[i] = 1
	}
	for j := n - 1; j >= 0; j-- {
		sum_partials := 0
		for i := j; i < n; i++ {
			// all the i where j is the parent of i
			sum_partials = grad_table[i] * o.derivative(x[i]) // not sure if this is correct? x[i] should be du_i/du_j
		}
		grad_table[j] = sum_partials
	}
	return grad_table[:ni]
}
