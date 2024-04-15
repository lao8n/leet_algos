package data_structures

import "container/heap"

// approaches
// * feels like a standard traversal problem - dfs where store minimum path to get to any location.
// * probably want to use bfs to represent each step through path

// specifics
// * choice between calculate all distances and times - and then see if time is valid? but actually may have been another path which is slower but not doesn't dissappear do need to update each time period.
// * create find neighbour function where add condition that node has to be alive?
// * what condition do we want to break out of loop?
func minimumTime(n int, edges [][]int, disappear []int) []int {
	// adjList
	adjList := make(map[int][]edge)
	for _, e := range edges {
		from, to, weight := e[0], e[1], e[2]
		// add edge both ways
		adjList[from] = append(adjList[from], edge{to: to, weight: weight})
		adjList[to] = append(adjList[to], edge{to: from, weight: weight})
	}

	// dp array of min times
	dp := make([]int, n)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	// setup heap
	h := make(Heap, 1)
	h[0] = 0
	heap.Init(&h)

	// bfs dijkstra
	for h.Len() > 0 {
		popped := heap.Pop(&h).(int)
		ns := neighbours(adjList, dp, popped, disappear)
		if len(ns) > 0 {
			for _, n := range ns {
				heap.Push(&h, n)
			}
		}
	}
	return dp
}

type edge struct {
	to     int
	weight int
}

func neighbours(adjList map[int][]edge, dp []int, curNode int, disappear []int) []int {
	neighbours := make([]int, 0)
	curTime := dp[curNode]
	if curTime == -1 || curTime >= disappear[curNode] {
		return neighbours
	}
	for _, e := range adjList[curNode] {
		// either first time getting there or new faster time
		newTime := curTime + e.weight
		if newTime < disappear[e.to] { // can visit
			if dp[e.to] == -1 || newTime < dp[e.to] {
				dp[e.to] = newTime
				neighbours = append(neighbours, e.to) // got here with new time
			}
		}
	}
	return neighbours
}

type Heap []int

func (h *Heap) Push(x interface{}) {
	(*h) = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
	previous, popped := *h, 0
	*h, popped = previous[:h.Len()-1], previous[h.Len()-1]
	return popped
}
func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(x, y int) bool {
	return h[x] < h[y]
}
func (h Heap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
