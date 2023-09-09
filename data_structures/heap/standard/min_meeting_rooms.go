package data_structures

import (
	"container/heap"
	"sort"
)

// tried thinking about this as a dynamic programming problem but unclear how to track all possible meeting room combinations. would need state as index of each meeting each 1, 2, 3, 4 as "14 23" to indicate 14 and together and 23 together
func minMeetingRooms(intervals [][]int) int {
	// sort by start time O(n logn)
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})

	// min heap of end times
	endHeap := make(Heap, 0)
	heap.Init(&endHeap)
	for _, interval := range intervals {
		// earliest ending room is free
		if endHeap.Len() > 0 && endHeap.Peek().(int) <= interval[0] {
			heap.Pop(&endHeap)
		}
		// add room back
		heap.Push(&endHeap, interval[1])
	}
	return len(endHeap)
}

// heap of end times
type Heap []int

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), 0
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h Heap) Peek() interface{} {
	return h[0]
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] } // min heap on end time
