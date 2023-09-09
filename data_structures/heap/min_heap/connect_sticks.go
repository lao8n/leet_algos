package data_structures

import "container/heap"

// trick is to also combine the smallest sticks
// we will maintain a stack of sticks
func connectSticks(sticks []int) int {
	h := make(Heap, len(sticks))
	copy(h, sticks)
	heap.Init(&h)

	total := 0
	for h.Len() > 1 {
		smallest, secondSmallest := heap.Pop(&h).(int), heap.Pop(&h).(int)
		sum := smallest + secondSmallest
		heap.Push(&h, sum)
		total = total + sum
	}
	return total
}

type Heap []int

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), 0
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
