package data_structures

// approaches
// * fairly mechanical calculation - but need additional grid to store new values
// * trick is can have 4 states rather than 2.. i.e. 0, 1, 2 (dead but used to be alive), 3 (alive but used to be dead)

// specifics
// * time complexity O(2 n^2)
func gameOfLife(board [][]int) {
	// process each element on board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = nextState(board, i, j)
		}
	}
	// set each element to 0 or 1
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] = getCurrentState(board, i, j)
		}
	}
	// no return
}

// Any live cell with fewer than two live neighbors dies as if caused by under-population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by over-population.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
func nextState(board [][]int, x int, y int) int {
	aliveCount := countLiveCells(board, x, y)
	// live cell
	if getPrevState(board, x, y) == 1 {
		if aliveCount < 2 {
			return 2
		} else if aliveCount > 3 {
			return 2
		} else {
			return 1 // stay alive
		}
	} else if getPrevState(board, x, y) == 0 {
		// dead cell
		if aliveCount == 3 {
			return 3
		}
	}
	return board[x][y] // unchanged
}

// assumes that x, y is valid
func getPrevState(board [][]int, x int, y int) int {
	switch board[x][y] {
	case 0, 3:
		return 0
	case 1, 2:
		return 1
	}
	return 0
}

func getCurrentState(board [][]int, x int, y int) int {
	return board[x][y] % 2
}

func countLiveCells(board [][]int, x int, y int) int {
	neighbours := possNeighbours(board, x, y)
	aliveCount := 0
	for _, neighbour := range neighbours {
		if validCoords(board, neighbour[0], neighbour[1]) {
			if getPrevState(board, neighbour[0], neighbour[1]) == 1 {
				aliveCount++
			}
		}
	}
	return aliveCount
}

func validCoords(board [][]int, x int, y int) bool {
	return x >= 0 && y >= 0 && x < len(board) && y < len(board[0])
}

func possNeighbours(board [][]int, x int, y int) [][]int {
	neighbours := make([][]int, 0)
	// up, down, left, right
	translations := []int{1, 0, -1, 0, 1}
	for i := 1; i < len(translations); i++ {
		xDelta, yDelta := translations[i-1], translations[i]
		neighbours = append(neighbours, []int{x + xDelta, y + yDelta})
	}
	// diag
	translations = []int{1, 1, -1, -1, 1}
	for i := 1; i < len(translations); i++ {
		xDelta, yDelta := translations[i-1], translations[i]
		neighbours = append(neighbours, []int{x + xDelta, y + yDelta})
	}
	return neighbours
}
