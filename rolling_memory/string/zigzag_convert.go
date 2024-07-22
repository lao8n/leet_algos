package data_structures

// approaches
// * want to produce row by row -> so need to calculate distances between letters and what letters go where

func convert(s string, numRows int) string {
	rows := make([][]rune, numRows)
	upDownLen := numRows
	if numRows >= 2 {
		upDownLen += numRows - 2
	}
	upI := numRows - 2
	for i, c := range s {
		upDownI := i % upDownLen
		if upDownI == 0 { // reset
			upI = numRows - 2
		}
		if upDownI <= numRows-1 { // go down
			rows[upDownI] = append(rows[upDownI], c)
		} else { // go up
			for k := 0; k < numRows; k++ {
				if k == upI {
					rows[k] = append(rows[k], c)
				}
			}
			upI--
		}
	}
	solution := ""
	for j := 0; j < numRows; j++ {
		solution += string(rows[j])
	}
	return solution
}
