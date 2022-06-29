package data_structures

import "sort"

// approach 1 = 1. sort O(n logn) 2. merge O(n)
// approach 2 = 1. merge sort combined
func merge(intervals [][]int) [][]int {
	// 1. custom sort - comparing just start index
	sort.Slice(intervals,
		func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
	mergedIntervals := [][]int{}
	// 2. merge intervals
	for _, interval := range intervals {
		if len(mergedIntervals) == 0 {
			mergedIntervals = append(mergedIntervals, interval)
		} else {
			// overlap
			if mergedIntervals[len(mergedIntervals)-1][1] >= interval[0] {
				// but some not overlapping
				if interval[1] >= mergedIntervals[len(mergedIntervals)-1][1] {
					mergedIntervals[len(mergedIntervals)-1][1] = interval[1]
				}
				// no overlap
			} else {
				mergedIntervals = append(mergedIntervals, interval)
			}
		}
	}
	return mergedIntervals
}
