package recursion

import "math"

// interconnected problems combined for a solution -> dynamic programming

// 2 approaches to dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// top down
// 1. base cases
// 2. recurrence case
// 3. state variables = house number + available colours -> cost

// time complexity O(n)
// space complexity O(n)
func minCost(costs [][]int) int {
	memoization := make([][]int, len(costs))
	for i := 0; i < len(costs); i++ {
		memoization[i] = make([]int, len(costs[0]))
	}
	d := data{
		costs:       costs,
		memoization: memoization,
	}
	return d.minCostsRecursion(0, -1)
}

type data struct {
	costs       [][]int
	memoization [][]int
}

type colour int

const (
	red colour = iota
	blue
	green
)

func (d *data) minCostsRecursion(houseNum int, unavailableColour colour) int {
	// base cases
	if houseNum >= len(d.costs) {
		return 0
	}
	// memoization
	if unavailableColour >= 0 && d.memoization[houseNum][unavailableColour] != 0 {
		return d.memoization[houseNum][unavailableColour]
	}

	// recurrence
	redCost, blueCost, greenCost := math.MaxInt, math.MaxInt, math.MaxInt
	if red != unavailableColour {
		redCost = d.costs[houseNum][red] + d.minCostsRecursion(houseNum+1, red)
	}
	if blue != unavailableColour {
		blueCost = d.costs[houseNum][blue] + d.minCostsRecursion(houseNum+1, blue)
	}
	if green != unavailableColour {
		greenCost = d.costs[houseNum][green] + d.minCostsRecursion(houseNum+1, green)
	}
	minCost := redCost
	if blueCost < minCost {
		minCost = blueCost
	}
	if greenCost < minCost {
		minCost = greenCost
	}
	if unavailableColour >= 0 {
		d.memoization[houseNum][unavailableColour] = minCost
	}

	return minCost
}
