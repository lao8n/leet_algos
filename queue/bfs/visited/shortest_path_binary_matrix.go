package data_structures

// shortest path -> bfs
// memoize? yes but do it on neighbour not when processing node o/w queue gets too long
// time: O(n), space: O(n)
// problem is we process neighbours arbitrarily when clearly it makes sense to priortise
// neighbours in the correct direction hence we could use a* search (which gives O(n logn))
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	// setup queue
	queue := make([]queueItem, 0)
	queue = append(queue, queueItem{xy: []int{0, 0}, length: 1})
	// memoize
	grid[0][0] = 1
	// bfs
	for len(queue) > 0 {
		// base case
		current := queue[0]
		queue = queue[1:]
		x, y, length := current.xy[0], current.xy[1], current.length
		// reached bottom corner
		if x == n-1 && y == n-1 {
			return length
		}

		// neighbours
		length += 1
		for _, neighbour := range getValidNeighbours(grid, []int{x, y}) {
			grid[neighbour[0]][neighbour[1]] = 1
			queue = append(queue, queueItem{xy: neighbour, length: length})
		}
	}
	return -1
}

type queueItem struct {
	xy     []int
	length int
}

func getValidNeighbours(grid [][]int, xy []int) [][]int {
	n, x, y := len(grid), xy[0], xy[1]
	neighbours := make([][]int, 0)
	// starting from top and going clockwise
	// [-1, -1] [-1, 0] [-1, 1]
	// [0, -1]  []      [0, 1]
	// [1, -1]  [1, 0]  [1, 1]
	transform := []int{-1, 0, 1, 0, -1, -1, 1, 1, -1}
	for i := 1; i < len(transform); i++ {
		nx, ny := x+transform[i-1], y+transform[i]
		// on grid
		if nx >= 0 && nx < n && ny >= 0 && ny < n {
			// only add 0s - we turn to 1 after processed effectively memoizing
			if grid[nx][ny] == 0 {
				neighbours = append(neighbours, []int{nx, ny})
			}
		}
	}

	return neighbours
}
