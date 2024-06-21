package data_structures

// approaches
// * tarjan's algorithm - actual algorithm for this problem
// * edge is a critical connection iff it is not a cycle
// * cycle detection usually just with dfs + visited nodes - however need a way to know all edges that are part of cycle
// * solution is dfs + visited count where if there is a more direct route then it must be a cycle.

// complexity
// * time = O(V + E)
// * space = O(E)
func criticalConnections(n int, connections [][]int) [][]int {
	// setup
	rank := make(map[int]int)
	adj := make(map[int][]int)
	conn := make(map[[2]int]struct{}) // we order by smaller node number
	for _, connection := range connections {
		x, y := connection[0], connection[1]
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
		conn[connKey(x, y)] = struct{}{}
	}
	d := data{rank, adj, conn}
	d.dfs(0, 1)
	result := [][]int{}
	for c, _ := range d.conn {
		result = append(result, []int{c[0], c[1]})
	}
	return result
}

type data struct {
	rank map[int]int
	adj  map[int][]int
	conn map[[2]int]struct{}
}

func (d *data) dfs(node int, discoveryRank int) int {
	if d.rank[node] > 0 {
		return d.rank[node]
	}
	d.rank[node] = discoveryRank
	minRank := discoveryRank + 1
	for _, neighbour := range d.adj[node] {
		// skip parent
		if neighbourRank, ok := d.rank[neighbour]; ok && neighbourRank == discoveryRank-1 {
			continue
		}
		recursiveRank := d.dfs(neighbour, discoveryRank+1)
		if recursiveRank <= discoveryRank { // there is a faster or equivalent route here
			delete(d.conn, connKey(node, neighbour))
		}
		if recursiveRank < minRank {
			minRank = recursiveRank
		}
	}
	return minRank
}

func connKey(x, y int) [2]int {
	if x < y {
		return [2]int{x, y}
	}
	return [2]int{y, x}
}
