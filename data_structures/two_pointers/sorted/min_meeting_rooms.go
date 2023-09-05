package data_structures

import "sort"

func minMeetingRooms(intervals [][]int) int {
	// two sorted slices O(n + 2nlogn)
	starts, ends := make([]int, len(intervals)), make([]int, len(intervals))
	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)

	// two pointers O(n)
	s, e, n := 0, 0, 0
	for s < len(intervals) {
		// if new start is greater than smallest end than we use the free room
		if starts[s] >= ends[e] {
			n -= 1
			e += 1
		}
		// always add new room and increment start
		n += 1
		s += 1
	}
	return n
}
