package data_structures

// dfs - 2 approaches - 1. look for counter example 2. check all paths.
// can retry path so memoization not important but need to detect cycles.
// this is especially complicated because we need to memoized both 1. when a node has been visited and 2. the previous calculation
func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	// construct adjacency map
	adjacencyMap := make(map[int][]int)
	memoized := make(map[int]bool, n)
	for _, edge := range edges {
		adjacencyMap[edge[0]] = append(adjacencyMap[edge[0]], edge[1]) // directed graph
	}
	if len(adjacencyMap[destination]) > 0 {
		return false
	}
	return findPath(source, destination, adjacencyMap, memoized)
}

// if find cycle return false, if don't find destination return false
func findPath(node int, destination int, adjacencyMap map[int][]int, memoized map[int]bool) bool {
	// arrived at destination
	if node == destination {
		return true
	}
	// detect cycle
	if _, ok := memoized[node]; ok { // been here before on this specific path - whether true or not
		return false
	}
	memoized[node] = false // no path but we have been here
	// need all paths be true and there to be at least one path
	if len(adjacencyMap[node]) == 0 {
		return false
	}
	for _, neighbour := range adjacencyMap[node] {
		if !memoized[neighbour] && !findPath(neighbour, destination, adjacencyMap, memoized) {
			return false
		}
		memoized[neighbour] = true
	}
	return true
}
