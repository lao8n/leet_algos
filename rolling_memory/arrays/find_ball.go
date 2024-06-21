package data_structures

// clarifying questions
// * to right -> 1 and to left is -1

// approaches
// * manually calculate each level down - if left and right together or wall then stuck
// * save calculations only calculate balls remaining rather than all?

// specifics
// * use bfs variation with a queue to track with depth for each level
func findBall(grid [][]int) []int {
	// init queue
	queue := make([]int, len(grid[0])) // length of row
	for i := range queue {             // check index i in row
		queue[i] = i
	}
	// init solution
	solution := make([]int, len(grid[0])) // original index - and current index
	for i := range solution {
		solution[i] = i
	}
	// loop over levels
	row := 0
	for len(queue) > 0 && row < len(grid) {
		// process level
		ql := len(queue)
		for i := 0; i < ql; i++ {
			popped := queue[i]
			// calc new index - could be -1
			solution[popped] = newPosition(grid[row], solution[popped])
			// fmt.Println(queue, i, popped, solution)
			if solution[popped] != -1 {
				queue = append(queue, popped)
			}
		}
		queue = queue[ql:]
		// update row
		row++
	}
	return solution
}

func newPosition(row []int, col int) int {
	// redirects to right i.e. 1
	if col >= 0 && col < len(row) && row[col] == 1 {
		if col+1 < len(row) && row[col+1] == 1 {
			return col + 1
		}
	}
	// redirects to left i.e. -1
	if col >= 0 && col < len(row) && row[col] == -1 {
		if col-1 >= 0 && row[col-1] == -1 {
			return col - 1
		}
	}
	// left wall
	// right wall
	// middle blockage
	return -1
}
