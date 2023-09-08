package data_structures

// approach: dfs with backtracking for each path
// recursion 1. accumulate paths 2. return? -> accumulate single path return all paths
// memoization 1. yes to track -> in place or separate track?
// time complexity = O(2^V *V) for a dag where O(V) per path for O(2^V *V)
// space complexity = O(V)
// argument is that every new node doubles the number of possible paths
func allPathsSourceTarget(graph [][]int) [][]int {
	return recurse(graph, 0, make([]int, 0), make(map[int]bool))
}

func recurse(graph [][]int, cur int, acc []int, seen map[int]bool) [][]int {
	// base case -> no more paths then return solution
	if cur == len(graph)-1 {
		return [][]int{append(acc, cur)}
	}
	if seen[cur] {
		return [][]int{}
	}
	// recursive case
	solutions := make([][]int, 0)
	path := make([]int, len(acc))
	copy(path, acc)
	path = append(path, cur)

	seen[cur] = true
	for _, neighbour := range graph[cur] {
		recursed := recurse(graph, neighbour, path, seen)
		solutions = append(solutions, recursed...)
	}
	// backtrack
	seen[cur] = false
	return solutions
}
