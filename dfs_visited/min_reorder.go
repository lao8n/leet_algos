package data_structures

// clarifying questions
// * n cities
// * one way to travel - i.e. it is a tree

// approaches
// * min reorder is only reorder? i.e. just flip every connection which is going the wrong way
// * recurse from root to leaves treating all roads as bidirectional - for every road that is wrong way round count...
func minReorder(n int, connections [][]int) int {
	// two-way adjList & road correct way
	adjList := make(map[int][]direction)
	for _, connection := range connections {
		from, to := connection[0], connection[1]
		adjList[from] = append(adjList[from], direction{node: to, toNode: true})
		adjList[to] = append(adjList[to], direction{node: from, toNode: false})
	}

	visited := make([]bool, n)
	// recurse from root to leaves counting along the way -> bfs or dfs?

	flip := recurse(0, adjList, visited)
	return flip
}

func recurse(node int, adjList map[int][]direction, visited []bool) int {
	// base case
	if visited[node] {
		return 0
	}
	// count
	visited[node] = true
	// recurse
	flip := 0
	for _, neighbour := range adjList[node] {
		if !visited[neighbour.node] {
			// fmt.Println(node, neighbour.node, neighbour.toNode)
			if neighbour.toNode {
				flip++
			}
			flip += recurse(neighbour.node, adjList, visited)
		}
	}
	return flip
}

type direction struct {
	node   int
	toNode bool
}
