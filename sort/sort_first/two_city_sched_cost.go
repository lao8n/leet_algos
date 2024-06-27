package data_structures

import "sort"

// approaches
// * if choose a then affect is b - could be dp
// * if everyone is in a then need to choose n to go to b - sort by delta?

func twoCitySchedCost(costs [][]int) int {
	minCost := 0
	deltaCosts := make([]int, len(costs))
	for i, cost := range costs {
		aCost, bCost := cost[0], cost[1]
		minCost += aCost
		deltaCosts[i] = bCost - aCost // choose smallest
	}

	sort.Ints(deltaCosts)
	n := len(costs) / 2
	for i := 0; i < n; i++ {
		minCost += deltaCosts[i]
	}

	return minCost
}
