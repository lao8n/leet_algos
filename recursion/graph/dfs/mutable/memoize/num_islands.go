package leet_algos

// time complexity = O(MxN) -> every possible starting point
// space complexity = O(MxN) worst case where all land and dfs searches MxN deep

// plan: for each found 1 -> do dfs look for all 1s that are connected
// seen: tick off seen 1s and 0s in a seen graph

// note here we track whether seen based upon updating grid[i][j] rather than a separate seen matrix
func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == byte('1') {
				count++
				dfs(location{x: i, y: j}, grid)
			}
		}
	}
	return count
}

func dfs(l location, grid [][]byte) {
	grid[l.x][l.y] = byte('0')
	for _, n := range getLandNeighbours(l, grid) {
		dfs(n, grid)
	}
	return
}

func getLandNeighbours(l location, grid [][]byte) []location {
	neighbours := make([]location, 0, 4)
	possibleNeighbours := []location{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	for _, pn := range possibleNeighbours {
		if validateGridLocation(location{x: l.x + pn.x, y: l.y + pn.y}, len(grid), len(grid[0])) {
			if grid[l.x+pn.x][l.y+pn.y] == byte('1') {
				neighbours = append(neighbours, location{x: l.x + pn.x, y: l.y + pn.y})
			}
		}
	}
	return neighbours
}

func validateGridLocation(l location, gridWidth int, gridLength int) bool {
	if l.x >= 0 && l.x < gridWidth && l.y >= 0 && l.y < gridLength {
		return true
	}
	return false
}

type location struct {
	x int
	y int
}
