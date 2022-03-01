package leet_algos

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	numRows, numCols := len(grid), len(grid[0])
	numIslands := 0
	dirs := []int{0, 1, 0, -1, 0}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if grid[i][j] == '1' {
				numIslands++
				bfs()
			}
		}
	}
}
