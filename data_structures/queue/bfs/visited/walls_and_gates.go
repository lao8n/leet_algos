// shortest distance -> bfs

// choices
// 1. update self or update neighbours?
// 2. from each location or the gate?
// 3. how memoize? based upon value? queue memoized or not?

func wallsAndGates(rooms [][]int) {
	numRows, numCols := len(rooms), len(rooms[0])

	// find gates O(mn)
	queue := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	// bfs from all gates simultaneously
	for len(queue) > 0 {
		// fmt.Printf("queue: %v\n", queue)
		// poll queue
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		// add neighbours
		neighbours := findNeighbours(rooms, x, y)
		for _, neighbour := range neighbours {
			rooms[neighbour[0]][neighbour[1]] = rooms[x][y] + 1
			queue = append(queue, neighbour)
		}
	}
}

func findNeighbours(rooms [][]int, x, y int) [][]int {
	numRows, numCols := len(rooms), len(rooms[0])
	neighbours := make([][]int, 0)
	candidates := []int{0, 1, 0, -1, 0}
	for i := 1; i < len(candidates); i++ {
		newX, newY := x+candidates[i-1], y+candidates[i]
		if newX >= 0 && newX < numRows && newY >= 0 && newY < numCols {
			if rooms[newX][newY] == math.MaxInt32 {
				neighbours = append(neighbours, []int{newX, newY})
			}
		}
	}
	return neighbours
}