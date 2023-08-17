package data_structures

import "container/heap"

// prim's algorithm
// 1. create visited and adjacency maps
// 2. heap on first node
// 3. loop where add minimum edge that doesn't include visited nodes, add neighbour edges to heap
// 4. stop when visited all nodes

// time complexity O(n^2 log(n)) = each push/pop takes log(n) for n^2 edges (n nodes)
// space complexity O(n^2) = size of heap
// optimization = store array of minimum cost to reach each node for O(n^2) and O(n)
func minCostConnectPoints(points [][]int) int {
	// create visited map
	visited := make(map[point]bool, len(points))
	adjacencyMap := make(map[point][]edge, len(points))

	// create point adjaceny map
	for i, point1 := range points {
		pointI := point{x: point1[0], y: point1[1]}
		visited[pointI] = false // not yet visited
		for j := i + 1; j < len(points); j++ {
			pointJ := point{x: points[j][0], y: points[j][1]}
			e := edge{i: pointI, j: pointJ, cost: manhattanDistance(pointI, pointJ)}
			adjacencyMap[pointI] = append(adjacencyMap[pointI], e)
			adjacencyMap[pointJ] = append(adjacencyMap[pointJ], e)
		}
	}

	// start with first point
	point1 := point{x: points[0][0], y: points[0][1]}
	point1Edges := adjacencyMap[point1]
	h := make(Heap, len(point1Edges))
	copy(h, point1Edges)
	heap.Init(&h)
	visited[point1] = true
	visitedCount := 1
	rollingCost := 0

	// add lowest edge in heap loop
	for h.Len() > 0 {
		// check if visited all nodes
		if visitedCount == len(visited) {
			return rollingCost
		}
		// get minimum edge and add node
		minEdge := heap.Pop(&h).(edge)
		var newPoint point
		if !visited[minEdge.i] { // exactly one of edge's two nodes should be new
			newPoint = minEdge.i
		} else if !visited[minEdge.j] {
			newPoint = minEdge.j
		} else {
			// already visited so go to the next one
			continue
		}
		visited[newPoint] = true
		visitedCount += 1
		rollingCost += minEdge.cost

		// add all of new nodes edges
		for _, neighbourEdge := range adjacencyMap[newPoint] {
			// do not add edge if already visited
			if !(visited[neighbourEdge.i] && visited[neighbourEdge.j]) {
				heap.Push(&h, neighbourEdge)
			}
		}
	}
	return rollingCost
}

type point struct {
	x int
	y int
}

type edge struct {
	i    point
	j    point
	cost int
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
func (h Heap) Less(i, j int) bool { return h[i].cost < h[j].cost }

func manhattanDistance(i point, j point) int {
	xDist := i.x - j.x
	if xDist < 0 {
		xDist *= -1
	}
	yDist := i.y - j.y
	if yDist < 0 {
		yDist *= -1
	}
	return xDist + yDist
}
