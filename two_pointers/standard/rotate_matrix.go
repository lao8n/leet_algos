package data_structures

// rotation can sometimes do something clever with reverse
// probably just need to keep multiple indices so when swap into correct place
// also need to keep track of where it used to be?
// swap across axes
// [1, 2, 3]    [9, 6, 3]
// [4, 5, 6] -> [8, 5, 2]
// [7, 8, 9]    [7, 4, 1]
// another idea is to rotate groups of four cells
func rotate(matrix [][]int) {
	// rotation is a flip on diagonal and a flip on horizontal
	// diagonal flip is 0, 1 index -> 1, 2
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; i+j < n; j++ {
			matrix[i][j], matrix[m-1-j][n-1-i] = matrix[m-1-j][n-1-i], matrix[i][j]
		}
	}
	// flip on horizontal
	for i := 0; i < (m / 2); i++ {
		matrix[i], matrix[m-1-i] = matrix[m-1-i], matrix[i]
	}
}
