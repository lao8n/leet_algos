// dynamic programming
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// 3 components
// 1. base cases
// 2. recurrence
// 3. state = string, 2 digit key flag -> number of ways - recursing forward but return solution - state is more complicated as need to check previous number as 1 can have up to 9 but 2 only up to 6

// time complexity O(n)
// space complexity O(1)
import "strconv"

func numDecodings(s string) int {
	i1, i2 := 1, 0 // numDecodings from index + 1 and index + 2
	// recursion goes from end -> front for num ways - iteration might be more natural to flip
	for i := len(s) - 1; i >= 0; i-- {
		numWays := 0
		d, _ := strconv.Atoi(string(s[i]))
		if d > 0 {
			numWays += i1
		}
		if i < len(s)-1 {
			dNext, _ := strconv.Atoi(string(s[i+1]))
			if d == 1 || (d == 2 && dNext <= 6) {
				numWays += i2
			}
		}
		i1, i2 = numWays, i1
	}
	return i1
}