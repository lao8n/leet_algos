package data_structures

import "fmt"

// clarifying questions
// * unique? no but just need one
// approaches
// * n binary searches - every row look for number.. O(n logn)
// * search by quadrant.. look for a num at midpoint.
// * can be cleverer by not just looking at midpoint element but top-left and bottom-right of every matrix whether it is valid.
func searchMatrix(matrix [][]int, target int) bool {
	// edge cases
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// setup
	d := data{
		matrix: matrix,
		target: target,
	}

	// search
	return d.searchQuadrant(0, 0, len(matrix)-1, len(matrix[0])-1)
}

type data struct {
	matrix [][]int
	target int
}

func (d *data) searchQuadrant(x1, y1, x2, y2 int) bool { // inclusive
	if x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 || x1 >= len(d.matrix) || x2 >= len(d.matrix) || y1 >= len(d.matrix[0]) || y2 >= len(d.matrix[0]) || x1 > x2 || y1 > y2 {
		return false
	}
	mx, my := (x1+x2)/2, (y1+y2)/2
	fmt.Println(x1, y1, x2, y2, mx, my, d.matrix[mx][my])
	if d.matrix[mx][my] == d.target {
		return true
	} else if d.matrix[mx][my] > d.target {
		// search in quadrant that is smaller i.e.
		return d.searchQuadrant(x1, y1, mx-1, y2) || // above
			d.searchQuadrant(mx, y1, x2, my-1) // bottom left
	} else {
		return d.searchQuadrant(mx+1, y1, x2, my) || // bottom left quadrant
			d.searchQuadrant(x1, my+1, x2, y2) // right quadrant
	}
	return false
}
