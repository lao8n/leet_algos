package data_structures

import "sort"

// clarifying questions
// * constraint is the size of the rectangles
// * seems like vertical height is irrelevant. just horizontal matters

// approaches
// * dp because looking for minimum - and decision for one rectangle affects others
// * can't you just greedily cover points from left and right? what scenario is that not the right approach? maybe could have put one in the middle? rather two left and right?

// specifics
// * what is dp array - list of x coordinates and number of rectangles covered up to which value?
// * worth sorting first? but costs O(n logn)
func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] // only care about x
	})
	// greedily choose max rectangle
	minNum := 0
	currentCoveredUpTo := -1
	for _, point := range points {
		if point[0] > currentCoveredUpTo {
			minNum++
			currentCoveredUpTo = point[0] + w
		} else {
			// already covered
		}
	}
	return minNum
}
