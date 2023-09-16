package data_structures

// brute force - pick each vertex and bfs to find height but O(n * logn)
// as we need all maybe we cannot -> but perhaps we can memoize some of the calculations
// essentially we need the longest path from every node
// create adjacency list and create map of longest-paths and single bfs and update?
// for every node we want to know two longest branches that leave it..
// idea is we trim away the leaves of the tree until either one or two nodes left..
// we don't need to track the height directly
func findMinHeightTrees(n int, edges [][]int) []int {
	// base cases
	if n == 1 {
		return []int{0}
	}

	adjList := make(map[int][]int, 0)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		adjList[a] = append(adjList[a], b)
		adjList[b] = append(adjList[b], a) // both ways
	}

	inDegreeMap := make(map[int]int, n)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		inDegreeMap[i] = len(adjList[i])
		if inDegreeMap[i] == 1 { // this is an edge node
			queue = append(queue, i)
		}
	}
	// key trick is do not need to count paths directly just trim leaves
	for n > 2 {
		ql := len(queue)
		for i := 0; i < ql; i++ {
			popped := queue[0]
			queue = queue[1:]
			n--
			// neighbours search for new leaves
			for _, neighbour := range adjList[popped] {
				inDegreeMap[neighbour]--
				if inDegreeMap[neighbour] == 1 {
					queue = append(queue, neighbour)
				}
			}
		}
	}
	return queue
}
