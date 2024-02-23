package data_structures

// clarifying questions

// approaches
// brute force - loop over matrix O(mn) looking for zeroes - if find a zero then go over row and column for O(m + n) so total time is O(mn + z(m + n))
// find zeroes first - O(mn) with look up for rows and columns so as loop over elements know.. then loop over again where see if row or column is in the map and if so set to zero..
// go diagonally from top left to bottom right - although still might need to go back?
// can you do a rolling set zero - where just set the neighbour to zero and then that sets its neighbour to zero? i guess y0ou need to know whether it is a starting zero where go up-down and left-right versus just continue
func setZeroes(matrix [][]int) {
	// setup
	m, n := len(matrix), len(matrix[0])
	zeroRows, zeroCols := make(map[int]bool), make(map[int]bool) // O(1) space optimisation is to use first cell as flag

	// look for zeroes - if find update backwards only
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// look for zero
			if matrix[i][j] == 0 {
				if !zeroRows[i] {
					// if find update up and left - only do this if not already true
					for y := j - 1; y >= 0; y-- {
						matrix[i][y] = 0
					}
					zeroRows[i] = true
				}
				if !zeroCols[j] {
					for x := i - 1; x >= 0; x-- {
						matrix[x][j] = 0
					}
					zeroCols[j] = true
				}
				continue
			}
			// check if zero row or col and update
			if zeroRows[i] || zeroCols[j] {
				matrix[i][j] = 0
			}
		}
	}

}
