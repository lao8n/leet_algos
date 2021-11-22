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
	out := make([][]int, len(matrix[0]))
	// old matrix coordinates r, c, new matrix coordinates i, j
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			r, c := j, len(matrix[0])-1-i // jth col = rth row, and
			out[i] = append(out[i], matrix[r][c])
		}
	}
	return out
}