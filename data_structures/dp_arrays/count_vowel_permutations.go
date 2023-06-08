package data_structures

// interconnected problems combined for a solution -> dynamic programming
// 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 components
// 1. base cases
// 2. recurrence
// 3. state variables -> 5 possible states (each vowel) -> state machine between them

// time complexity O(n)
// space complexity O(1)

const (
	a = iota
	e
	i
	o
	u
)

func countVowelPermutation(n int) int {
	currentWays := []int{1, 1, 1, 1, 1}
	for ; n > 1; n-- {
		nextWays := make([]int, 5)
		// a -> e
		nextWays[e] += currentWays[a]
		// e -> a or i
		nextWays[a] += currentWays[e]
		nextWays[i] += currentWays[e]
		// i -> a, e, o, u (not i)
		nextWays[a] += currentWays[i]
		nextWays[e] += currentWays[i]
		nextWays[o] += currentWays[i]
		nextWays[u] += currentWays[i]
		// o -> i, u
		nextWays[i] += currentWays[o]
		nextWays[u] += currentWays[o]
		// u -> a
		nextWays[a] += currentWays[u]

		for j, w := range nextWays {
			currentWays[j] = w % 1000000007
		}
	}
	sum := 0
	for _, w := range currentWays {
		sum += w
	}
	return sum % 1000000007
}
