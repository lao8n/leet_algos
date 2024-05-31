package data_structures

import (
	"container/heap"
	"sort"
)

// clarifying questions
// * sorted - yes in terms of left point

// approaches
// * could have increasing stack but then what do we do with green building?, maybe could pop until green is relevant?
// * heap sorted by 1. heights and 2. ending right coordinate.
func getSkyline(buildings [][]int) [][]int {
	// setup edges
	edges := make([][]int, 0)
	for i, b := range buildings {
		l, r := b[0], b[1]
		edges = append(edges, []int{l, i})
		edges = append(edges, []int{r, i})
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] < edges[j][0] // sort by position
	})

	// setup heap
	h := make(Heap, 0)
	heap.Init(&h)

	// loop
	result := make([][]int, 0)
	for i := range edges {
		currX := edges[i][0]
		// process all edges at x - there may be multiple
		for i < len(edges) && edges[i][0] == currX {
			bi := edges[i][1]
			// if left edge add r & h to heap
			if buildings[bi][0] == currX {
				heap.Push(&h, building{right: buildings[bi][1], height: buildings[bi][2]})
			}
			i++
		}
		// if have passed tallest building
		for h.Len() > 0 && h.Peek().right <= currX {
			heap.Pop(&h)
		}
		// change in height
		currHeight := 0
		if h.Len() > 0 {
			currHeight = h.Peek().height
		}
		if len(result) == 0 || result[len(result)-1][1] != currHeight {
			result = append(result, []int{currX, currHeight})
		}
	}
	return result
}

type building struct {
	right  int
	height int
}

type Heap []building

func (h *Heap) Push(x interface{}) {
	(*h) = append(*h, x.(building))
}
func (h *Heap) Pop() interface{} {
	previous := *h
	popped := previous[h.Len()-1]
	(*h) = previous[:h.Len()-1]
	return popped
}
func (h Heap) Peek() building {
	return h[0] // not h.Len() - 1
}
func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(x, y int) bool {
	if h[x].height == h[y].height {
		return h[x].right < h[y].right // smallest right index
	}
	return h[x].height > h[y].height // tallest building
}
func (h Heap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
