package data_structures

// approaches
// * dfs with backtrack
// * dfs in both directions?
// * dfs where have visited path map at each point
// * dp i.e. dfs with memoization where cahce represents longest path starting at a point
// * topological sort - process in order

// specifics
// * build one-way based upon if smaller -> bigger.
func longestIncreasingPath(matrix [][]int) int {
	// in-degree map and adjacency list (no need can calculate on the fly)
	m, n := len(matrix), len(matrix[0])
	inDegree := make([][]int, m)
	queue := make([][]int, 0)

	for i := 0; i < m; i++ {
		inDegree[i] = make([]int, n)
		for j := 0; j < n; j++ {
			neighbours := validNeighbours(matrix, i, j)
			for _, neighbour := range neighbours {
				if matrix[neighbour[0]][neighbour[1]] < matrix[i][j] {
					inDegree[i][j]++
				}
			}
			if inDegree[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	// bfs (peeling onion) with queue
	pathLen := 0
	for len(queue) > 0 {
		batch := len(queue)
		for i := 0; i < batch; i++ {
			popped := queue[i]
			px, py := popped[0], popped[1]
			for _, neighbour := range validNeighbours(matrix, px, py) {
				nx, ny := neighbour[0], neighbour[1]
				if matrix[nx][ny] > matrix[px][py] {
					inDegree[nx][ny]--
					if inDegree[nx][ny] == 0 {
						queue = append(queue, []int{nx, ny})
					}
				}
			}
		}
		pathLen++
		queue = queue[batch:]
	}

	return pathLen
}

func validNeighbours(matrix [][]int, i int, j int) [][]int {
	dirs := []int{1, 0, -1, 0, 1}
	neighbours := [][]int{}
	for d := 1; d < len(dirs); d++ {
		dx, dy := dirs[d-1], dirs[d]
		nx, ny := i+dx, j+dy
		// in-degree is number of dirs that are smaller so check if neighbour is smaller
		if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) {
			neighbours = append(neighbours, []int{nx, ny})
		}
	}
	return neighbours
}
