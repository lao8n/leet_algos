package data_structures

import "sort"

// kruskal's algorithm
// 1. sort edges -> could use heap, easier to just sort
// 2. add edge if no cycle -> union find
// 3. stop at n-1 edges

// union find -> quick union with 1. union by rank 2. memoized find
func minCostConnectPoints(points [][]int) int {
	roots := make(map[point]point, len(points))
	rank := make(map[point]int, len(points))

	// create edges
	edges := make([]edge, 0)
	for i, point1 := range points {
		// union find init
		pointI := point{x: point1[0], y: point1[1]}
		roots[pointI] = pointI
		rank[pointI] = 1

		// symmetric edge cost calculation so only need one way
		for j := i + 1; j < len(points); j++ {
			pointJ := point{x: points[j][0], y: points[j][1]}
			e := edge{i: pointI,
				j:    pointJ,
				cost: manhattanDistance(pointI, pointJ)}
			edges = append(edges, e)
		}
	}
	// union find init
	ps := uf{
		root: roots,
		rank: rank,
	}

	// sort edges
	sort.Slice(edges, func(i int, j int) bool { // ascending so smallest at front
		return edges[i].cost < edges[j].cost
	})

	// add edge if no cycle
	numEdges := 0
	totalCost := 0
	for _, edge := range edges {
		if numEdges == len(points)-1 {
			return totalCost
		}
		// if either end of edge are already connected then adding edge -> circular
		if !ps.connected(edge.i, edge.j) {
			ps.union(edge.i, edge.j)
			numEdges += 1
			totalCost += edge.cost
		}
	}
	// return cost
	return totalCost // can get here because of (-10000, -10000), (+10000, +10000) case
}

type edge struct {
	i    point
	j    point
	cost int
}

type point struct {
	x int
	y int
}

type uf struct {
	root map[point]point
	rank map[point]int
}

func (uf *uf) find(i point) point {
	// base case
	if i == uf.root[i] {
		return i
	}
	// recursive case
	uf.root[i] = uf.find(uf.root[i])
	return uf.root[i]
}

func (uf *uf) union(i point, j point) {
	iRoot, jRoot := uf.find(i), uf.find(j)
	if iRoot != jRoot {
		if uf.rank[iRoot] < uf.rank[jRoot] {
			uf.root[iRoot] = jRoot
		} else if uf.rank[jRoot] < uf.rank[iRoot] {
			uf.root[jRoot] = iRoot
		} else {
			uf.root[iRoot] = jRoot
			uf.rank[jRoot] += 1
		}
	}
}

func (uf *uf) connected(i point, j point) bool {
	iRoot, jRoot := uf.find(i), uf.find(j)
	return iRoot == jRoot
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
