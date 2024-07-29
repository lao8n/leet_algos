package data_structures

import (
	"container/heap"
	"sort"
)

// clarifying question
// * are they sorted by start-time? no they are not

// approaches
// * dp question where if you attend at one time can't attend another etc.
// * 2 approaches to dp either 1. top down recursion with memoization or 2. bottom up tabulation
// * rolling window where choose which event to attend
// * greedy - just choose to attend event with least end time
// * sort by start time then by end time

// edge cases
// * shouldn't attend an event unless no other events at that time because
func maxEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	// setup heap
	h := make(Heap, 0)
	heap.Init(&h)

	numEvents := 0
	i, d := 0, 0
	for i < len(events) || h.Len() > 0 {
		// push all events
		for i < len(events) && events[i][0] <= d {
			heap.Push(&h, event{events[i][0], events[i][1]})
			i++
		}
		for h.Len() > 0 {
			popped := heap.Pop(&h).(event)
			// fmt.Println(i, d, popped, h, numEvents)
			if popped.end >= d {
				numEvents++
				break
			}
		}
		d++
	}
	return numEvents
}

type event struct {
	start int
	end   int
}
type Heap []event

func (h *Heap) Push(x interface{}) {
	(*h) = append(*h, x.(event))
}
func (h *Heap) Pop() interface{} {
	previous := *h
	popped := previous[h.Len()-1]
	*h = previous[:h.Len()-1]
	return popped
}
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].end < h[j].end } // want to be min heap
