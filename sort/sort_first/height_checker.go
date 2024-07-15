package data_structures

import "sort"

// clarifying questions
// * non-decreasing (left shorter, taller right

// approaches
// * get expected heights by sorting O(n logn) - O(n with bucket sort)
// * rolling count to detect out of place
func heightChecker(heights []int) int {
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)
	sort.Ints(sortedHeights) // smaller -> larger

	diff := 0
	for i, height := range heights {
		if height != sortedHeights[i] {
			diff++
		}
	}
	return diff
}
