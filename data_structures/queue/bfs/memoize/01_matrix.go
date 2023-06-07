package data_structures

// want shortest path distance -> bfs i.e. use a queue
// choice bfs from points or bfs from zeros? probably latter is easier
// question how deal with two different zeros leading to different distances?
// one answer is to loop over all zeroes at once..
// another question - how do we know when we have updated all points? counter?
func updateMatrix(mat [][]int) [][]int {
	queue := make([]point, 0)
	toDo := make(map[point]bool, 0)
	// O(mn) - find zeros and count non-zeros
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				queue = append(queue, point{i, j, 0})
			}
			toDo[point{i, j, 0}] = true // always 0
		}
	}
	// bfs with queue
	for len(queue) > 0 && len(toDo) > 0 {
		popped := queue[0]
		queue = queue[1:]

		// new point
		if toDo[point{popped.x, popped.y, 0}] {
			mat[popped.x][popped.y] = popped.distance
			delete(toDo, point{popped.x, popped.y, 0})
			queue = append(queue, findNeighbours(mat, popped)...)
		}
	}
	return mat
}

// only finds valid neighbours, doesn't check for memoization
func findNeighbours(mat [][]int, popped point) []point {
	numRows := len(mat)
	numCols := len(mat[0])
	neighbours := make([]point, 0)
	translations := []int{1, 0, -1, 0, 1}
	for i := 1; i < len(translations); i++ {
		xTranslate, yTranslate := translations[i-1], translations[i]
		newX, newY := popped.x+xTranslate, popped.y+yTranslate
		if newX >= 0 && newY >= 0 && newX < numRows && newY < numCols {
			neighbours = append(neighbours, point{newX, newY, popped.distance + 1})
		}
	}
	return neighbours
}

type point struct {
	x        int
	y        int
	distance int
}
