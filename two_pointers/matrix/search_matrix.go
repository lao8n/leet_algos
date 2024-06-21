package data_structures

// clarifying questions
// * unique? no but just need one
// approaches
// * n binary searches - every row look for number.. O(n logn)
// * search by quadrant.. look for a num at midpoint.
// * bottom-left or top-right and search across space for O(m + n)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := len(matrix)-1, 0 // bottom left
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else {
			i--
		}
	}
	return false
}
