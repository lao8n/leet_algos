package data_structures

func maxKilledEnemies(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	maxCount := 0
	rowHits := 0              // [ -> ] going horizontally along a row
	colHits := make([]int, n) // vertical slice
	for i := 0; i < m; i++ {  // row at a time
		for j := 0; j < n; j++ { //
			// reset row hits - go forward to figure out row count
			if j == 0 || grid[i][j-1] == 'W' {
				rowHits = 0
				for k := j; k < n; k++ {
					if grid[i][k] == 'W' {
						break
					} else if grid[i][k] == 'E' {
						rowHits++
					}
				}
			}
			// reset hits on col - go forward to figure out col count
			if i == 0 || grid[i-1][j] == 'W' {
				colHits[j] = 0
				for k := i; k < m; k++ {
					if grid[k][j] == 'W' {
						break
					} else if grid[k][j] == 'E' {
						colHits[j]++
					}
				}
			}
			if grid[i][j] == '0' {
				totalHits := rowHits + colHits[j]
				if totalHits > maxCount {
					maxCount = totalHits
				}
			}
		}
	}
	return maxCount
}
