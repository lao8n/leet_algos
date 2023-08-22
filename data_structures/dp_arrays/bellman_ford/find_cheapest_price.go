package data_structures

import "math"

// approach = no negative numbers -> dijkstra, could also use bellman-ford
// bellman-ford = dynamic programming
// dynamic programming = 1. top down with memoization or 2. bottom up with tabulation
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// adjacency list
	adj := make(map[int][][]int) // src -> [[destination, cost]]
	for _, flight := range flights {
		fSrc, fDst, fCost := flight[0], flight[1], flight[2]
		adj[fSrc] = append(adj[fSrc], []int{fDst, fCost})
	}

	// dp table = max number of edges x destination
	dpCurr := make([]int, n)
	dpPrev := make([]int, n)
	for i := 0; i < n; i++ {
		dpCurr[i] = math.MaxInt
	}
	dpCurr[src] = 0
	// dp loop
	k += 1 // k is number of stops not number of flights
	for k > 0 {
		copy(dpPrev, dpCurr)
		// possible sources
		for i := 0; i < n; i++ {
			// possible destinations
			for _, flight := range adj[i] {
				flightDst, flightCost := flight[0], flight[1]
				if dpPrev[i] < math.MaxInt && dpPrev[i]+flightCost < dpCurr[flightDst] {
					dpCurr[flightDst] = dpPrev[i] + flightCost
				}
			}
		}
		k--
	}
	if dpCurr[dst] == math.MaxInt {
		return -1
	}
	return dpCurr[dst]
}
