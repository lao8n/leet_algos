package data_structures

// approaches
// * same function rotate matrix
// * different function in four direction
func spiralOrder(matrix [][]int) []int {
	top, bottom, left, right := 1, len(matrix)-1, 0, len(matrix[0])-1
	solution := make([]int, len(matrix)*len(matrix[0]))
	s, i, j := 0, 0, -1 // start at (0, -1) outside loop

	// outer loop
	for s < len(solution) {
		// inner loop
		// go right
		for j < right {
			j++
			solution[s] = matrix[i][j]
			s++
		}
		right--
		// go down
		for i < bottom {
			i++
			solution[s] = matrix[i][j]
			s++
		}
		bottom--
		// go left
		for s < len(solution) && j > left {
			j--
			solution[s] = matrix[i][j]
			s++
		}
		left++
		// go up
		for s < len(solution) && i > top {
			i--
			solution[s] = matrix[i][j]
			s++
		}
		top++
	}
	return solution
}
