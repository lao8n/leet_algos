package data_structures

import "math"

// can't just greedily go from squareroot downwards
// could maybe do some dynamic programming approach but no constraints on using same integer twice etc.

// approach use bfs until find a solution
// state is remaining amount, length of sequence, last number
// we prevent duplicate approaches by constraining new number to be the same or smallest than last

// complexity
// time = O(sqrt(n) * n)
// space = O(n)
func numSquares(n int) int {
	// setup
	squares := allPerfectSquares(n)
	queue := make([]state, 1)
	queue[0] = state{
		nRemaining:      n,
		length:          0,
		lastNumberIndex: 0,
	}

	// queue
	for len(queue) > 0 {
		polled := queue[0]
		queue = queue[1:]

		// base case
		if polled.nRemaining == 0 {
			return polled.length
		}

		// add to queue
		for i := polled.lastNumberIndex; i < len(squares); i++ {
			if squares[i] <= polled.nRemaining {
				s := state{
					nRemaining:      polled.nRemaining - squares[i],
					length:          polled.length + 1,
					lastNumberIndex: i,
				}
				queue = append(queue, s)
			}
		}
	}

	return n // can always do all 1s
}

type state struct {
	nRemaining      int
	length          int
	lastNumberIndex int
}

func allPerfectSquares(n int) []int {
	squares := make([]int, 0)
	sqrt := int(math.Sqrt(float64(n)))
	for i := sqrt; i > 0; i-- {
		squares = append(squares, i*i)
	}
	return squares
}
