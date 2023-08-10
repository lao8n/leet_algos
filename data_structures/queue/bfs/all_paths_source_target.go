package data_structures

// approach bfs
// how append node to correct paths in solution? -> not appending to all
// how only add if actually reach destination
// how to prevent duplicates? cannot use the node - because can visit multiple times - only the path
// so need a set?
// don't need to memoize because it is a directed graph
// time complexity = O(2^V * V) = at most 2^{V-1} - 1 possible paths to go from the starting vertex to the target and
// each takes O(V) time
// space complexity = queue can contain O(2^V) paths and each path takes O(V) space for O(2^V * V)
// argument is that every new node doubles the number of possible paths
func allPathsSourceTarget(graph [][]int) [][]int {
	// already have adjacency list

	// setup queue, memoized & solution
	queue := make([][]int, 1)
	queue[0] = []int{0} // source
	paths := make([][]int, 0)

	// queue loop
	for len(queue) > 0 {
		// queue update
		popped := queue[0]
		poppedNode := popped[len(popped)-1]
		queue = queue[1:]

		if poppedNode == len(graph)-1 { // found the target
			paths = append(paths, popped)
		}

		// neighbours
		for _, neighbour := range graph[poppedNode] {
			newPath := make([]int, len(popped)+1)
			copy(newPath, popped)
			newPath[len(popped)] = neighbour
			queue = append(queue, newPath)
		}
	}
	return paths
}
