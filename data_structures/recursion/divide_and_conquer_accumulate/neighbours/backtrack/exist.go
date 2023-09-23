package data_structures

// approach dfs where accumulate possible path
// update board for visited - can do in place
// need to backtrack if something doesn't work
func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    var existRecursive func(int, int, int) bool
    existRecursive = func(x, y, i int) bool {
        // base cases
        if x < 0 || x >= m || y < 0 || y >= n {
            return false
        }
        if board[x][y] != word[i] {
            return false
        }
        if i == len(word) - 1 {
            return true
        }
        // recursive case
        temp := board[x][y]
        board[x][y] = '#'
        found := existRecursive(x + 1, y, i + 1) ||
            existRecursive(x - 1, y, i + 1) ||
            existRecursive(x, y + 1, i + 1) ||
            existRecursive(x, y - 1, i + 1)
        // backtrack
        board[x][y] = temp

        return found
    }

    for i, row := range board {
        for j := range row {
            if existRecursive(i, j, 0) {
                return true
            }
        }
    }
    return false
}