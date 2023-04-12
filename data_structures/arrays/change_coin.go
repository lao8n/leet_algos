package data_structures

// interdependent subproblems combined for a solution -> dynamic programming
// dynamic programming has 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// try top down - 3 components
// 1. base cases
// 2. recurrence relation - can add c coin amount to x amounts
// 3. state = amount

// how to handle duplicates - if use smaller coin cannot use larger coins anymore
// optimized for a single array

// time complexity = O(len(coins) * amount)
// space complexity = O(len(coins))
func change(amount int, coins []int) int {
	tabulation := make([]int, amount+1)
	tabulation[0] = 1

	for _, c := range coins {
		for i := c; i < amount+1; i++ {
			tabulation[i] += tabulation[i-c]
		}
	}
	return tabulation[amount]
}
