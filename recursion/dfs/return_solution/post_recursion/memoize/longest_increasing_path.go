// choice: look for increasing path, or look for decreasing
// alt choice: go in both directions

// TODO
// naive = dfs -> time limit exceeded, time complexity is O(2^{m + n}), space complexity
// O(mn)
// time complexity = all nodes visited once which is O(mn) for mxn nodes, or one node with path
// O(mxn) for all nodes
// visit/search = within each dfs search do not check the same place once
// memoization across searches = add memoization

// notes memoization only works here because it is strictly increasing so cannot have a cycle
func longestIncreasingPath(matrix [][]int) int {
	// loop over every starting point
	longestPath := 0
	// avoid repeat path calculations across dfs searches
	memoizedPaths := make([][]int, len(matrix))
	for i := range memoizedPaths {
		memoizedPaths[i] = make([]int, len(matrix[0]))
	}
	// longestPath
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			pathLength := dfs(i, j, matrix, memoizedPaths)
			if pathLength > longestPath {
				longestPath = pathLength
			}
		}
	}
	return longestPath
}

// we want pairs '0,1', '1,0', '-1,0', '0,-1'
var leftRightUpDown = []int{0, 1, 0, -1, 0}

func dfs(x, y int, matrix [][]int, memoizedPaths [][]int) int {
	if memoizedPaths[x][y] != 0 {
		return memoizedPaths[x][y]
	}
	for px := 0; px < len(leftRightUpDown)-1; px++ {
		newX, newY := x+leftRightUpDown[px], y+leftRightUpDown[px+1]
		if isValid(newX, newY, matrix) && matrix[newX][newY] > matrix[x][y] {
			newLength := dfs(newX, newY, matrix, memoizedPaths)
			if newLength > memoizedPaths[x][y] {
				memoizedPaths[x][y] = newLength
			}
		}
	}
	memoizedPaths[x][y]++
	return memoizedPaths[x][y]
}

func isValid(x, y int, matrix [][]int) bool {
	if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
		return true
	}
	return false
}