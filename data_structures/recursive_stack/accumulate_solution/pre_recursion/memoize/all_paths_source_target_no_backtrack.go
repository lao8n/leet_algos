package data_structures

// approach: dfs with backtracking for each path
// recursion 1. accumulate paths 2. return? -> accumulate single path return all paths
// memoization 1. yes to track -> in place or separate track?
func allPathsSourceTarget(graph [][]int) [][]int {
	return recurse(graph, 0, make([]int, 0))
}

func recurse(graph [][]int, cur int, acc []int) [][]int {
	// base case -> no more paths then return solution
	if cur == len(graph)-1 {
		return [][]int{append(acc, cur)}
	}
	// recursive case
	solutions := make([][]int, 0)
	path := make([]int, len(acc))
	copy(path, acc)
	path = append(path, cur)
	seen := make(map[int]bool, len(path))
	// might as well create new set each time as if pass pointer will be shared
	for _, node := range path {
		seen[node] = true
	}
	for _, neighbour := range graph[cur] {
		// only recurse if not seen
		if _, ok := seen[neighbour]; !ok {
			recursed := recurse(graph, neighbour, path)
			// fmt.Println(recursed)
			solutions = append(solutions, recursed...)
		}
	}
	return solutions
}
