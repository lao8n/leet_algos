package data_structures

// interconnected problems combined for a solution - esp as num of ways -> dynamic programming

// 2 approaches
// 1. top down recursion with memoization
// 2. bottom up iteration with tabulation

// try bottom up - as potential to optimise space into perhaps just the previous column

// 3 components to dp
// 1 base cases
// 2 recurrence
// 3 state -> as we don't care about horizontal symmetry just need to know 1. num of n left 2. whether
// 1 or 2 spaces to fill

// time complexity = O(n)
// space complexity = O(1)
func numTilings(n int) int {
	mod := 1000000007
	oneFilledWays := []int{0, 1, 0}
	twoFilledWays := []int{0, 2, 1}
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	for i := 2; i < n; i++ {
		// choose one space way
		// [] + [][] -> [][][]
		// []   []      [][]
		// or
		// [] +      -> []
		//      [][]    [][]
		oneFilledWays[0] = (oneFilledWays[1] + twoFilledWays[2]) % mod
		// choose two space ways
		// [] + [] -> []
		// []   []    []
		// or
		// [] + [][] -> [][][]
		// []   [][]    [][][]
		// or
		// [] +   [] -> [][] * 2
		//      [][]    [][]
		twoFilledWays[0] = (oneFilledWays[1]*2 + twoFilledWays[1] + twoFilledWays[2]) % mod
		// move along
		oneFilledWays[1], oneFilledWays[2] = oneFilledWays[0], oneFilledWays[1]
		twoFilledWays[1], twoFilledWays[2] = twoFilledWays[0], twoFilledWays[1]
	}
	return twoFilledWays[1]
}
