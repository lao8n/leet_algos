// package main
package data_structures

import (
	"fmt"
)

// backpropagation is essentially dynamic programming where we tabulate answers backwards filling in the table
// not sure if this solution works but the implementation is based upon the pseudo-code from bengio's deep learning
// textbook p206-207

type operation struct {
	function   func([]int) int
	derivative func(int) int
}

func (o operation) forward_pass(x []int, n int) []int {
	ni := len(x)
	u := make([]int, n)
	for i := 0; i < ni; i++ {
		u[i] = x[i]
	}
	for i := ni; i < n; i++ {
		parents := u[:i] // really weird parent function
		u[i] = o.function(parents)
	}
	fmt.Printf("u = %v\n", u)
	return u
}

func (o operation) backward_pass(u []int, ni int, n int) []int {
	grad_table := make([]int, n) // du_n/du_i
	for i := 0; i < n-1; i++ {
		grad_table[i] = 1
	}
	grad_table[n-1] = o.derivative(u[n-1]) // this isn't right
	for j := ni - 1; j >= 0; j-- {
		sum_partials := 0
		for i := ni; i < n; i++ { // j's children are all the i that follow
			sum_partials = grad_table[i] * o.derivative(u[j]) // not sure if this is correct? u[i] should be du_i/du_j
			fmt.Println(j, i, sum_partials)
		}
		grad_table[j] = sum_partials
	}
	return grad_table[:ni]
}

func main() {
	operation := operation{
		function: func(x []int) int {
			result := 0
			for i := 0; i < len(x); i++ {
				result += x[i] * x[i]
			}
			return result
		},
		derivative: func(x int) int {
			return 2 * x
		},
	}
	x := []int{3, 2, 4}
	u := operation.forward_pass(x, 4)
	gradients := operation.backward_pass(u, 3, 4)
	fmt.Println(gradients)
}
