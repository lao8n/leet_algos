package data_structures

import "container/heap"

// clarifying questions
// * only one missing number?

// approaches
// * brute force loop through looking for smallest number each time
// * heap of positive numbers where get minimum of heap to see if any missing - then pop off until missing value - heap is O(k logn) // where k is missing number
// * get max and min - then use euler's formula to figure out expected sum vs actual sum - doesn't work if duplicates

// specifics
// * how handle duplicates?
func firstMissingPositive(nums []int) int {
	// setup heap - can't just heapify as want to exclude neg numbers
	h := make(Heap, 0)
	heap.Init(&h)
	// loop through pushing to heap
	for _, num := range nums {
		if num > 0 {
			heap.Push(&h, num)
		}
	}
	// pop heap
	missing := 1
	popped, prevPopped := 0, 0
	for h.Len() > 0 {
		popped = heap.Pop(&h).(int)
		if prevPopped == popped { // duplicate
			continue
		}
		if popped > missing {
			return missing
		}
		missing++
		prevPopped = popped
	}
	return popped + 1
}

type Heap []int

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
	previous, popped := *h, 0
	*h, popped = previous[:h.Len()-1], previous[h.Len()-1]
	return popped
}
func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(x, y int) bool { // min heap
	return h[x] < h[y]
}
func (h Heap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
