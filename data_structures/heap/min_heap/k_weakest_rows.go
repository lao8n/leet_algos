package data_structures

import "container/heap"

// one option is to sort (quick sort) but by two criteria 1. number of soldiers and 2. row number
// could also create a heap and remove all elements

// we have a few options
// 1. min heap of size m where remove k elements
// 2. max heap of size k with test where we pop and push if smaller than max - then we pop in reverse order
func kWeakestRows(mat [][]int, k int) []int {
	// heapify O(mn)
	rows := make(Heap, len(mat))
	for i, _ := range mat {
		j := 0
		for j < len(mat[0]) && mat[i][j] == 1 {
			j++
		}
		rows[i] = Row{soldiers: j - 1, row: i}
	}
	// O(m)
	heap.Init(&rows)

	// remove elements O(k logn)
	solution := make([]int, k)
	i := 0
	for i < k {
		solution[i] = heap.Pop(&rows).(Row).row
		i++
	}
	return solution
}

type Row struct {
	soldiers int
	row      int
}
type Heap []Row

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Row)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), Row{}
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int { return len(h) }
func (h Heap) Less(i, j int) bool {
	if h[i].soldiers == h[j].soldiers {
		return h[i].row < h[j].row
	}
	return h[i].soldiers < h[j].soldiers
}
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
