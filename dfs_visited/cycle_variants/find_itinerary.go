package data_structures

import "sort"

// approaches
// * hierholzer's algorithm - 1. dfs until find cyclic path 2. if unvisited edges do detour
// * adapted algorithm - 1 . dfs until get stuck at a vertex 2. backtrack to nearest neighbour vertex until all edges used.
// complexity = O(E)

// definitions
// * eulerian path = path in a finite graph that visits every edge exactly once
// * eulerian cycle = is a eulerian path that starts and ends on the same vertex
func findItinerary(tickets [][]string) []string {
	adjMap := make(map[string][]string)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		adjMap[from] = append(adjMap[from], to)
	}
	for _, destinations := range adjMap {
		sort.Strings(destinations)
	}

	result := make([]string, 0)
	d := data{adjMap, result}
	d.dfs("JFK")
	reverse(d.result)
	return d.result
}

type data struct {
	adjMap map[string][]string
	result []string
}

func (d *data) dfs(from string) {
	for len(d.adjMap[from]) > 0 {
		to := d.adjMap[from][0]
		d.adjMap[from] = d.adjMap[from][1:]
		d.dfs(to)
	}
	d.result = append(d.result, from)
}

func reverse(output []string) {
	l, r := 0, len(output)-1
	for l < r {
		output[l], output[r] = output[r], output[l]
		l++
		r--
	}
}
