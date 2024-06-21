package data_structures

// clarifying questions
// * don't need to validate moves?
// approaches
// * have a 2-d array and manually add move then check conditions where O(n^2) space, and need to loop through horizontal and vertical O(2n) and maybe diagonal for O(3n)
// * only check conditions i.e. have n row conditions, n col conditions and 2 diagonal condition for O(2n + 2) space, and then update for each move can do + 1 and -1 - only costs O(4)
type TicTacToe struct {
	n    int
	rows []int
	cols []int
	diag []int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		n:    n,
		rows: make([]int, n),
		cols: make([]int, n),
		diag: make([]int, 2),
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	// update
	delta := 1
	if player == 2 {
		delta = -1
	}
	this.rows[row] += delta
	this.cols[col] += delta
	if row == col { // e.g. 3, 3
		this.diag[0] += delta
	}
	if row == this.n-1-col { // e.g. 0, n-1 or 1, n-2
		this.diag[1] += delta
	}

	// check
	playerWin := delta * this.n
	if this.rows[row] == playerWin || this.cols[col] == playerWin || this.diag[0] == playerWin || this.diag[1] == playerWin {
		return player
	}
	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
