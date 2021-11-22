func spiralOrder(matrix [][]int) []int {
	// we use dynamic programming and recursion
	if len(matrix) == 0 {
		return []int{}
	} else if len(matrix) == 1 {
		return matrix[0]
	} else {
		rotatedMatrix := rotate90Left(matrix[1:]) // everything but the first row
		return append(matrix[0], spiralOrder(rotatedMatrix)...)
	}
}

func rotate90Left(matrix [][]int) [][]int {
	// if matrix is mxn then rotated matrix is nxm
	m, n := len(matrix), len(matrix[0])
	out := make([][]int, n)
	// old matrix coordinates r, c, new matrix coordinates i, j
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			r, c := j, n-1-i // jth col = rth row, and
			out[i] = append(out[i], matrix[r][c])
		}
	}
	return out
}