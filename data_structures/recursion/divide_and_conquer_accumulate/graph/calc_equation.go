package data_structures

// approaches
// * union find calculate to see if connected if not then -1 o/w try to find a path
// * path find algorithm i.e. dfs - build a graph out of equations
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// adjaceny matrix for neighbours
	adjMap := make(map[string][]edge)

	for i, eq := range equations {
		from, to := eq[0], eq[1]
		e1 := edge{from, to, values[i], 1 / values[i]}
		adjMap[from] = append(adjMap[from], e1)
		e2 := edge{to, from, 1 / values[i], values[i]}
		adjMap[to] = append(adjMap[to], e2)
	}
	// dfs from that node to see if path to other node in query
	solution := []float64{}
	for _, query := range queries {
		from, to := query[0], query[1]
		// problems is could be circular
		visited := make(map[string]bool)
		solution = append(solution, dfs(from, to, 1, visited, adjMap))
	}
	return solution
}

// dfs want to accumulate values along the path
func dfs(cur string, to string, acc float64, visited map[string]bool, adjMap map[string][]edge) float64 {
	// base cases
	if visited[cur] {
		// handle already visited this node
		return -1.0
	}
	// found destination - should handle a/a case if initialize search with 1 for acc
	edges := adjMap[cur]
	if len(edges) == 0 {
		return -1
	}
	if cur == to {
		return acc
	}
	visited[cur] = true

	// recursive cases - loop over adj map
	for _, e := range edges {
		path := dfs(e.to, to, acc*e.fromDivTo, visited, adjMap)
		if path != -1.0 {
			return path
		}
	}
	// need to backtrack on visited path? no because just want a single path - and for each loop we start with a new visited map
	return -1.0
}

// how represent equation value? at the node or the edge?
// probably edge
// whenever add edge also add inverse
type edge struct {
	from      string
	to        string
	fromDivTo float64
	toDivFrom float64
}
