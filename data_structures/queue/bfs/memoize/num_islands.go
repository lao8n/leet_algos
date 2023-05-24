package recursion

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	numRows, numCols := len(grid), len(grid[0])
	numIslands := 0

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if grid[i][j] == '1' {
				numIslands++
				bfs([]int{i, j}, grid, numRows, numCols)
			}
		}
	}
	return numIslands
}

func bfs(xy []int, grid [][]byte, numRows, numCols int) {
	dirs := []int{0, 1, 0, -1, 0}
	queue := [][]int{xy}
	for len(queue) != 0 {
		pop := queue[0]
		queue = queue[1:]
		for k := 0; k < 4; k++ {
			x := pop[0] + dirs[k]
			y := pop[1] + dirs[k+1]
			if x >= 0 && x < numRows && y >= 0 && y < numCols && grid[x][y] == '1' {
				grid[x][y] = '0'
				queue = append(queue, []int{x, y})
			}
		}
	}
}
