package data_structures

// question: can multiple oranges start rotting? probably
// approaches:
// 1. do bfs from every rotten orange (i.e. queue could start > 1), until no fresh left (counter to track)
// 2. what about isolated oranges? bfs is inefficient there -> connected components?
// memoize on neighbours or self? happens through rotting
func orangesRotting(grid [][]int) int {
	// queue setup
	queue := make([][]int, 0) // current rotten
	numFresh := 0
	numMinutes := 1
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if grid[i][j] == 1 {
				numFresh += 1
			} else if grid[i][j] == 2 { // rotten
				queue = append(queue, []int{i, j})
			}
		}
	}

	// base cases
	if numFresh == 0 {
		return 0
	}
	if len(queue) == 0 {
		return -1
	}

	// bfs
	for len(queue) > 0 {
		// process queue
		nextMinute := len(queue)
		for i := 0; i < nextMinute; i++ {
			popped := queue[0]
			queue = queue[1:]
			for _, neighbour := range validNeighbours(grid, popped) {
				queue = append(queue, neighbour)
				// update on neighbour not process - as want to avoid duplicates in neighbours
				grid[neighbour[0]][neighbour[1]] = 2
				numFresh -= 1
				if numFresh == 0 {
					return numMinutes
				}
			}
		}
		numMinutes += 1
	}

	return -1
}

func validNeighbours(grid [][]int, xy []int) [][]int {
	x, y, m, n := xy[0], xy[1], len(grid), len(grid[0])
	transformation := []int{0, 1, 0, -1, 0}
	neighbours := make([][]int, 0)
	for i := 1; i < len(transformation); i++ {
		nx, ny := x+transformation[i-1], y+transformation[i]
		if nx >= 0 && nx < m && ny >= 0 && ny < n {
			if grid[nx][ny] == 1 { // fresh
				neighbours = append(neighbours, []int{nx, ny})
			}
		}
	}
	return neighbours
}
