package data_structures

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */
// clarifying questions
// * properties -> monotonically increasing i.e. increase x or y -> increase z

// approaches
// * x and y can be between 1 and 1000 so try every pair of combinations i.e. 1000^2
// * binary search on x to find where custom function intercepts with z i.e. for every single y do binary search on x which is O(n * logn) -> should only be a single value -> if smallest value doesn't intercept then can give up early?
// * can we uniquely identify the formula? - if it was a straight line yes - but we don't know the form?

func findSolution(customFunction func(int, int) int, z int) [][]int {
	solutions := make([][]int, 0)
	minY, maxY := 1, 1000

	// fix y, search x space
	for y := minY; y <= maxY; y++ {
		xValues := findXValues(y, customFunction, z)
		if xValues != nil {
			solutions = append(solutions, xValues)
		}
	}

	return solutions
}

func findXValues(y int, customFunction func(int, int) int, z int) []int {
	minX, maxX := 1, 1000
	for minX <= maxX {
		midX := int((minX + maxX) / 2)
		genZ := customFunction(midX, y)
		if genZ == z {
			return []int{midX, y}
		} else if genZ > z { // too big
			maxX = midX - 1
		} else if genZ < z {
			minX = midX + 1
		}
	}
	return nil
}
