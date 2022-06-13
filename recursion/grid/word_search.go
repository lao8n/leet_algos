package recursion

// two options
// 1. recurse forwards from each of the board locations, accumulating up to word
// 2. recurse backwards from each of the board locations, returning true if could have word
// q = memoize search, e.g. have a lookup table with 1. bool 2. word substring
// q: recurse forwards or backwards with string?
// q: how make sure don't recurse over locations?
func exist(board [][]byte, word string) bool {
	existsSolution := false
	// all starting locations
	for sx := 0; sx < len(board); sx++ {
		for sy := 0; sy < len(board[0]); sy++ {
			if existRecursive([]int{sx, sy}, board, word) {
				existsSolution = true
			}
		}
	}
	return existsSolution
}

func existRecursive(xy []int, board [][]byte, word string) bool {
	existsSolution := false
	if board[xy[0]][xy[1]] == word[0] {
		// base case
		if word[1:] == "" {
			return true
		}
		for _, nxy := range findNeighbours(xy, board) {
			board[xy[0]][xy[1]] = 0 // set to visited
			if existRecursive(nxy, board, word[1:]) {
				existsSolution = true
			}
			board[xy[0]][xy[1]] = word[0]
		}
	}
	//
	return existsSolution
}

func findNeighbours(xy []int, board [][]byte) [][]int {
	x := xy[0]
	y := xy[1]
	combinations := []int{0, 1, 0, -1, 0} // [0, 1], [1, 0], [0, -1], [-1, 0]
	neighbours := [][]int{}
	for i, j := 0, 1; j < len(combinations); i, j = i+1, j+1 { // no comma operator
		nx := x + combinations[i]
		ny := y + combinations[j]
		if nx >= 0 && nx < len(board) && ny >= 0 && ny < len(board[0]) {
			neighbours = append(neighbours, []int{nx, ny})
		}
	}
	return neighbours
}
