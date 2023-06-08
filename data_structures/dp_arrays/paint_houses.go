package data_structures

// interconnected problems combined for a solution -> dynamic programming

// 2 approaches to dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// bottom up iteration with optimised tabulation
// 1. base cases
// 2. recurrence case
// 3. state variables = house number + available colours -> cost

// time complexity O(n)
// space complexity O(1)

func minCost(costs [][]int) int {
	prevHouse := make([]int, 3) // 3 colours
	for i := 0; i < len(costs); i++ {
		currHouse := make([]int, 3)
		// paint red
		currHouse[0] = costs[i][0] + min(prevHouse[1], prevHouse[2])
		// paint blue
		currHouse[1] = costs[i][1] + min(prevHouse[0], prevHouse[2])
		// paint green
		currHouse[2] = costs[i][2] + min(prevHouse[0], prevHouse[1])
		prevHouse = currHouse
	}
	return min(prevHouse[0], min(prevHouse[1], prevHouse[2]))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
