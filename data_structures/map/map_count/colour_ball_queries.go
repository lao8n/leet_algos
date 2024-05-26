package data_structures

// clarifying questions
// * initially all balls have no colour
// * mark ball x with colour y

// approaches
// * naive approach just have an array of length limit + 1, then loop over counting all balls
// but this costs O(n^2) as O(n) for each query i.e. map of ball -> colour
// * map of balls -> colour to get last colour, map of colours -> count

// specifics
// * why is it not just the number of queries? because could be the same ball? and could be the
// same colour
// * so if same ball we will realise, and if new colour we will realise
func queryResults(limit int, queries [][]int) []int {
	ballColours := make(map[int]int)
	colourCounts := make(map[int]int)
	result := make([]int, len(queries))
	distinct := 0
	for i, query := range queries {
		x, y := query[0], query[1]
		// ball colour
		oldColour := ballColours[x]
		ballColours[x] = y

		// colour counts
		if oldColour != 0 {
			colourCounts[oldColour]-- // reduce count
			if colourCounts[oldColour] == 0 {
				distinct-- // we have lost a colour
			}
		}
		colourCounts[y]++
		if colourCounts[y] == 1 {
			distinct++
		}
		// distinct colours
		result[i] = distinct
	}
	return result
}
