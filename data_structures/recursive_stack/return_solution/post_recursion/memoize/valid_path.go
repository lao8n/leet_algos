package data_structures

// dfs from source until one of them reaches destination
// use a stack - indirectly using recursion
// recursion - 1. accumulate or 2. return -> accumulate path but return bool
// track where been 1. memoization - global rather than local
func validPath(n int, edges [][]int, source int, destination int) bool {
	edgesMap := make(map[int][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		edgesMap[from] = append(edgesMap[from], to)
		edgesMap[to] = append(edgesMap[to], from)
	}
	return recurse(edgesMap, source, destination, make(map[int]bool))
}

func recurse(edgesMap map[int][]int, curr int, destination int, path map[int]bool) bool {
	// base case
	if curr == destination {
		return true
	}
	// visited so ignore
	if path[curr] == true {
		return false
	}
	path[curr] = true // visited
	// recurse
	for _, neighbour := range edgesMap[curr] {
		if recurse(edgesMap, neighbour, destination, path) {
			return true
		}
	}
	// no backtrack as looking for a single path
	return false
}
