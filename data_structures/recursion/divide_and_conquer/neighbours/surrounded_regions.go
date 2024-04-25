package data_structures

// approach
// * find all O on the border, then find all O adjacent to O on the border O(2m + 2n)- flip everything else i.e. do dfs from the borders? O(V + E)
func solve(board [][]byte) {
	// find O on borders
	m, n := len(board), len(board[0])
	borderO := make([][]int, 0) // coords
	for i := 0; i < m; i++ {
		// first col
		if board[i][0] == 'O' {
			borderO = append(borderO, []int{i, 0})
		}
		// last col
		if board[i][n-1] == 'O' {
			borderO = append(borderO, []int{i, n - 1})
		}
	}
	for j := 1; j < n-1; j++ {
		// first row except first and last element
		if board[0][j] == 'O' {
			borderO = append(borderO, []int{0, j})
		}
		// last row except first and last element
		if board[m-1][j] == 'O' {
			borderO = append(borderO, []int{m - 1, j})
		}
	}
	// do dfs turning O to X from each border point - update in place
	keepAsZero := make(map[int]map[int]bool) // xy map to keep as zero
	for _, b := range borderO {
		dfs(board, b, keepAsZero)
	}
	// turn all others to X
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !keepAsZero[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, xy []int, keepAsZero map[int]map[int]bool) {
	if keepAsZero[xy[0]] == nil {
		keepAsZero[xy[0]] = make(map[int]bool)
	}
	keepAsZero[xy[0]][xy[1]] = true
	for _, neighbour := range getONeighbours(board, xy, keepAsZero) {
		dfs(board, neighbour, keepAsZero)
	}
}

func getONeighbours(board [][]byte, xy []int, keepAsZero map[int]map[int]bool) [][]int {
	neighbours := make([][]int, 0)
	translations := []int{1, 0, -1, 0, 1}
	for i := 1; i < len(translations); i++ {
		dx, dy := translations[i-1], translations[i]
		nx, ny := xy[0]+dx, xy[1]+dy
		if onBoard(len(board), len(board[0]), []int{nx, ny}) && board[nx][ny] == 'O' && !keepAsZero[nx][ny] {
			neighbours = append(neighbours, []int{nx, ny})
		}
	}
	return neighbours
}

func onBoard(m int, n int, xy []int) bool {
	return xy[0] >= 0 && xy[0] < m && xy[1] >= 0 && xy[1] < n
}
