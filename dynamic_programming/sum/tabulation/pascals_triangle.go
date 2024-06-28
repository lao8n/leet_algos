package data_structures

// approaches
// * brute force - manually calculate everything
// * dp - memoize previous row

// specifics
// * symmetric so can calculate two values at a time
func generate(numRows int) [][]int {
	solution := make([][]int, 0)

	// base case
	if numRows == 0 {
		return solution
	}
	solution = append(solution, []int{1})
	for i := 2; i <= numRows; i++ { // already done row 1
		prev := solution[len(solution)-1]
		cur := make([]int, i) // row length is ith row
		for j := 0; j <= (i-1)/2; j++ {
			if j == 0 {
				cur[j] = 1
				cur[i-1] = 1
			} else {
				cur[j] = prev[j-1] + prev[j]
				cur[i-1-j] = cur[j]
			}
		}
		solution = append(solution, cur)
	}
	return solution
}
