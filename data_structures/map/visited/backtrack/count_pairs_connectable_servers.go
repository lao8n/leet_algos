package data_structures

// clarifying questions
// * given slice of edges could be any number of them, we are not actually given n correct?
// * division constraint is weird - because it is the entire distance - not that they are directly connected?
// approaches
// * brute force - from edges build an adjaceny map -> build a tree at each node recurse down
// * dfs recursion down two paths - but what if overlapping paths?
// * dfs along a single path - generate all valid paths to c
// implementation
// * how deal with cycles? visited map - but it is an unrooted weight tree so there are no cycles...
// * can we optimize so don't have to calculate from scratch each time?
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	// build adj map
	adjMap := make(map[int][]edge, 0)
	for _, e := range edges {
		a, b, w := e[0], e[1], e[2]
		adjMap[a] = append(adjMap[a], edge{b, w})
		adjMap[b] = append(adjMap[b], edge{a, w})
	}
	// do dfs from each node - it is the minimum at every node (but includes two)
	solution := make([]int, len(adjMap))
	for i, edges := range adjMap {
		if len(edges) < 1 {
			solution[i] = 0
		} else {
			counts := make([]int, len(edges))
			visited := map[int]bool{i: true}
			for i, e := range edges {
				d := data{0, signalSpeed, adjMap}
				d.dfs(e.to, e.weight, visited)
				counts[i] = d.count
			}
			// for every pairwise combination of i & j add min
			for j := 0; j < len(counts); j++ {
				for k := j + 1; k < len(counts); k++ {
					solution[i] += counts[j] * counts[k]
				}
			}
		}
	}
	return solution
}

type data struct {
	count       int
	signalSpeed int
	adjMap      map[int][]edge
}

func (d *data) dfs(vertex int, pathWeight int, visited map[int]bool) {
	if visited[vertex] {
		return
	}
	if pathWeight%d.signalSpeed == 0 {
		d.count++
	}
	visited[vertex] = true
	// recursive case
	for _, edge := range d.adjMap[vertex] {
		d.dfs(edge.to, pathWeight+edge.weight, visited)
	}
	visited[vertex] = false
	return
}

type edge struct {
	to     int
	weight int
}
