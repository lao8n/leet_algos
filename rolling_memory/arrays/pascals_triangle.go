package data_structures

// appraoches
// * brute force - loop through each layer adding things up.. this will cost 2^n?
// * symmetry - the calculation is symmetric so we can get rid of half
// * could use the formula as well

// details
// * need to be careful about first row
func generate(numRows int) [][]int {
	solution := make([][]int, numRows)
	solution[0] = []int{1}
	for r := 1; r < numRows; r++ {
		row := make([]int, r+1) // r + 1 is length of row
		// make half
		row[0], row[r] = 1, 1
		for i := 1; i < r; i++ { // 5 / 2 = 2 i.e. includes middle by zero-indexing
			row[i] = solution[r-1][i-1] + solution[r-1][i] // add one behind and one in-line
		}
		solution[r] = row
	}
	return solution
}
