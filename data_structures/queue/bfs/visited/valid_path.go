// bfs to find a path
func validPath(n int, edges [][]int, source int, destination int) bool {
	edgesMap := make(map[int][]int, n)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		edgesMap[from] = append(edgesMap[from], to)
		edgesMap[to] = append(edgesMap[to], from)
	}

	queue := make([]int, 1)
	memoized := make([]bool, n)
	queue[0] = source
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		// base case
		if popped == destination {
			return true
		}
		if !memoized[popped] { // haven't been to this node before
			queue = append(queue, edgesMap[popped]...)
			memoized[popped] = true
		}

	}
	return false
}