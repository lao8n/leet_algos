package data_structures

// approach
// * every pair of points i.e. m^2 has a line that defines them so create an array of all lines with unique constant and slope then sort points O(m^2 log m^2) to see if constant slope the same.
// * or go over every pair and create a map of constant and slope and increment counts

// specifics
// * although slice is not valid as a key to a map, an array is
// * if there are 3 points on a line then there will be a, b, c, ab, ba, ac, ca, bc, cb 6 combinations

// alternative approach
// * instead for each point x, consider all other points and count that way
func maxPoints(points [][]int) int {
	max := 0
	// loop over all pairwise combinations
	for i, x := range points {
		counts := make(map[float64]int)
		for j, y := range points {
			if i != j {
				xi, xj := x[0], x[1]
				yi, yj := y[0], y[1]
				// slope y = mx + c
				slope := float64(yj-xj) / float64(yi-xi)
				// constant
				counts[slope]++
				if counts[slope] > max {
					max = counts[slope]
				}
			}
		}
	}
	return max + 1
}
