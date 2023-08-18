package data_structures

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	// create adjacency list
	adjList := make(map[int][]edge, 0)
	for _, time := range times {
		u, v, w := time[0], time[1], time[2]
		adjList[u] = append(adjList[u], edge{source: u, target: v, time: w})
	}

	// shortest path
	minimumTime := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		if i == k {
			minimumTime[i] = 0
		} else {
			minimumTime[i] = math.MaxInt
		}
	}

	// create heap with node 1 neighbours
	h := make(Heap, len(adjList[k]))
	heap.Init(&h)
	copy(h, adjList[k])
	heap.Init(&h)

	// bfs
	for h.Len() > 0 {
		// get new edge
		popped := heap.Pop(&h).(edge)

		// new time < old time
		prevTime := minimumTime[popped.source]
		if popped.time+prevTime < minimumTime[popped.target] {
			minimumTime[popped.target] = popped.time + prevTime
			// neighbours - only add if time is improved (don't need visited)
			for _, neighbour := range adjList[popped.target] {
				heap.Push(&h, neighbour)
			}
		}
	}

	max := math.MinInt
	for _, min := range minimumTime {
		if min > max {
			max = min
		}
	}

	if max == math.MaxInt { // unreachable node
		return -1
	}
	return max
}

type edge struct {
	source int
	target int
	time   int
}

type Heap []edge

func (h *Heap) Push(x interface{}) { *h = append(*h, x.(edge)) }
func (h *Heap) Pop() interface{} {
	previous, n, popped := *h, h.Len(), edge{}
	*h, popped = previous[:n-1], previous[n-1]
	return popped
}
func (h Heap) Len() int           { return len(h) }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Heap) Less(i, j int) bool { return h[i].time < h[j].time }
