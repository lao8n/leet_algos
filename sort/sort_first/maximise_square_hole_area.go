package data_structures

import "sort"

// clarifying questions
// * n + 2 horizontal and m + 2 vertical.. this is an n x m matrix
// * constraint to not just remove all the bars is that it has to be square?
// * can we assume that they are ordered? no they are not sorted

// thoughts
// * if you remove a vertical bar you also need to remove a horizontal bar and vice-versa - basically need to remove cross-point and if multiple they need to be along the diagonal

// approaches
// * dynamic programming approach feels like overkill the decisions are related but in a simple way...
// * just need the longest continuous unbroken string of numbers incrementing by 1...
func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	longestHBars := 0
	currentHBars := 0
	sort.Ints(hBars)
	for i := 0; i < len(hBars); i++ {
		// streak continuing?
		if i == 0 {
			currentHBars = 1
		} else if hBars[i] == hBars[i-1]+1 {
			currentHBars++
		} else {
			currentHBars = 1 // reset
		}
		// longest?
		if currentHBars > longestHBars {
			longestHBars = currentHBars
		}
	}
	longestVBars := 0
	currentVBars := 0
	sort.Ints(vBars)
	for i := 0; i < len(vBars); i++ {
		if i == 0 {
			currentVBars = 1
		} else if vBars[i] == vBars[i-1]+1 {
			currentVBars++
		} else {
			currentVBars = 1
		}
		if currentVBars > longestVBars {
			longestVBars = currentVBars
		}
	}
	// return minimum
	if longestHBars < longestVBars {
		return (longestHBars + 1) * (longestHBars + 1)
	}
	return (longestVBars + 1) * (longestVBars + 1)
}
