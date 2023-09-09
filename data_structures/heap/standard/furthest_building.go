package data_structures

import "container/heap"

// trick = use ladders on largest jumps (to preserve the bricks)
// approach maintain a heap of size ladders, it should be a min heap, where we pop out smallest number
func furthestBuilding(heights []int, bricks int, ladders int) int {
	// heap setup
	h := make(Heap, ladders)
	heap.Init(&h)

	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			heap.Push(&h, diff)
		}
		if h.Len() > ladders {
			popped := heap.Pop(&h)
			bricks -= popped.(int)
		}
		if bricks < 0 {
			return i - 1
		}
	}
	return len(heights) - 1
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
func (h Heap) Less(i, j int) bool { return h[i] < h[j] } // we want min heap
