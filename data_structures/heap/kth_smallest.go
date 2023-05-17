package data_structures

// one option is to insert all elements into a min heap O(mn) and then extract k for smallest
// second option is to maintain a max heap of size k but only insert into it if smaller than max
// third option - can we exploit fact they are sorted? k pointers perhaps? it is going to be some sort of connected triangular shape -> maybe heap again
import (
	"container/heap"
	"fmt"
)

func kthSmallest(matrix [][]int, k int) int {
	// make min heap O(1)
	h := make(Heap, 0)
	heap.Init(&h)

	// track duplicates with indices
	duplicates := make(map[string]bool, len(matrix)*len(matrix[0]))

	// rather than adding all elements we add just the new ones
	// add at most O(k elements) which for tree is O(k log k)
	heap.Push(&h, Data{row: 0, col: 0, val: matrix[0][0]})
	duplicates[fmt.Sprintf("%d %d", 0, 0)] = true
	for k > 0 {
		// pop minimal data
		popped := heap.Pop(&h).(Data)
		k--
		if k == 0 {
			return popped.val
		}
		// add below
		belowKey := fmt.Sprintf("%d %d", popped.row+1, popped.col)
		if popped.row+1 < len(matrix) && !duplicates[belowKey] {
			heap.Push(&h, Data{
				row: popped.row + 1,
				col: popped.col,
				val: matrix[popped.row+1][popped.col],
			})
			duplicates[belowKey] = true
		}
		// and right
		rightKey := fmt.Sprintf("%d %d", popped.row, popped.col+1)
		if popped.col+1 < len(matrix[0]) && !duplicates[rightKey] {
			heap.Push(&h, Data{
				row: popped.row,
				col: popped.col + 1,
				val: matrix[popped.row][popped.col+1],
			})
			duplicates[rightKey] = true
		}
	}
	return 0
}

type Data struct {
	row int
	col int
	val int
}
type Heap []Data

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(Data)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), Data{}
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i].val < h[j].val }
