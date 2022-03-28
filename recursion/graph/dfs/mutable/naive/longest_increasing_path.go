// choice: look for increasing path, or look for decreasing
// alt choice: go in both directions

// TODO
// do not check same path twice

// complexity
// time = each dfs search has complexity O(d), longest path is O(m + n), and there are 2 directions
// to go in so it is 2^{m+n}
// space = O(mn) where need O(d) for depth of recursion (system stack), which is in the worst case
// have depth = mn
func longestIncreasingPath(matrix [][]int) int {
	// loop over every starting point
	longestPath := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			pathLength := dfs(i, j, matrix, true) + dfs(i, j, matrix, false) + 1
			if pathLength > longestPath {
				longestPath = pathLength
			}
		}
	}
	return longestPath
}

// we want pairs '0,1', '1,0', '-1,0', '0,-1'
var leftRightUpDown = []int{0, 1, 0, -1, 0}

func dfs(x, y int, matrix [][]int, increasingFlag bool) int {
	bestLength := 0
	for px := 0; px < len(leftRightUpDown)-1; px++ {
		newX, newY := x+leftRightUpDown[px], y+leftRightUpDown[px+1]
		if isValid(newX, newY, matrix) && matrix[newX][newY] > matrix[x][y] {
			// fmt.Printf("%d %d\n", matrix[x][y], matrix[newX][newY])
			newLength := 0
			if increasingFlag && matrix[newX][newY] > matrix[x][y] {
				newLength = dfs(newX, newY, matrix, increasingFlag) + 1
			}
			if !increasingFlag && matrix[newX][newY] < matrix[x][y] {
				newLength = dfs(newX, newY, matrix, increasingFlag) + 1
			}
			if newLength > bestLength {
				bestLength = newLength
			}
		}
	}
	// fmt.Println(bestLength)
	return bestLength
}

func isValid(x, y int, matrix [][]int) bool {
	if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) {
		return true
	}
	return false
}