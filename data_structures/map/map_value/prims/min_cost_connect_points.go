package data_structures

import "math"

// prim's algorithm
func minCostConnectPoints(points [][]int) int {
	// create visited map and edge maps
	visited := make(map[point]bool, len(points))
	minEdge := make(map[point]int, len(points))
	for _, p := range points {
		pointP := point{x: p[0], y: p[1]}
		visited[pointP] = false
		minEdge[pointP] = math.MaxInt32
	}
	// initialize variable
	rollingCost := 0
	visitedCount := 0
	minEdge[point{x: points[0][0], y: points[0][1]}] = 0

	// prims algorithm - add lowest connected edge
	for visitedCount < len(points) {
		currNodeMinCost := math.MaxInt32
		var currNode point
		// pick min edge of one node
		for node, cost := range minEdge {
			if !visited[node] && cost < currNodeMinCost {
				currNode = node
				currNodeMinCost = cost
			}
		}
		// update values with picked edge
		rollingCost += currNodeMinCost
		visitedCount += 1
		visited[currNode] = true

		// add this node's neighbours as edges if lower costs
		for _, p := range points {
			pPoint := point{x: p[0], y: p[1]}
			cost := manhattanDistance(currNode, pPoint)
			if !visited[pPoint] && cost < minEdge[pPoint] {
				minEdge[pPoint] = cost
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
