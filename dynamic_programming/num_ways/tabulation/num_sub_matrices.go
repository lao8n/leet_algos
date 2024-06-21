package data_structures

// clarifying questions
// * rectangles not necessarily squares
// * different layers but can have for example a 1x2 compared with a 1x1 to give a 1x3

// approaches
// * num of ways -> dp but what is the memoization structure?
// * any rectangle is the combination of two rectangles
// * loop through every type of rectangle - but there is surely some repeated computation? could start bfs type search at index 1
// * union find but what all possible ways? not just one connected structure

// assertion - every rectangle can be uniquely described as a combination of a left and a right rectangle - maybe can have at each point all the vertical and horizontal rectangles
// dfs neighbour search from every point
// individual points, with 1 x n and nx1 just looping through and incrementing..
// how handle larger rectangles?
// maybe just check 3 points - bottom, right and bottom right and take the minimum? but we also need to get all subsets?
func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	// count horizontal strips
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if mat[i][j] == 1 {
				mat[i][j] += mat[i][j-1]
			}
		}
	}

	numWays := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			numWays += mat[i][j]

			localMin := mat[i][j]
			// add vertical strips
			for k := i + 1; k < m; k++ {
				if mat[k][j] == 0 {
					break
				}
				localMin = min(localMin, mat[k][j])
				numWays += localMin
			}
		}
	}
	return numWays
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
