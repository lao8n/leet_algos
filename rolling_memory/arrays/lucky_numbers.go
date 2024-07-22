package data_structures

// approaches
// * test each element -> i.e. O(mn) elements with O(m + n) scan
// * slice max and mins
func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	rowMins := make([]int, m)
	colMaxs := make([]int, n)
	rowMinsSet := make(map[int]struct{}, n)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if rowMins[r] == 0 || matrix[r][c] < rowMins[r] {
				rowMins[r] = matrix[r][c]
			}
			if colMaxs[c] == 0 || matrix[r][c] > colMaxs[c] {
				colMaxs[c] = matrix[r][c]
			}
		}
		// finished a column
		rowMinsSet[rowMins[r]] = struct{}{}
	}
	intersection := make([]int, 0)
	for c := 0; c < n; c++ {
		if _, ok := rowMinsSet[colMaxs[c]]; ok {
			intersection = append(intersection, colMaxs[c])
		}
	}
	return intersection
}
