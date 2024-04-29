package data_structures

// approaches
// * dfs with backtrack
// * dfs in both directions?
// * dfs where have visited path map at each point
// * dp i.e. dfs with memoization where cahce represents longest path starting at a point
// specifics
// * can we avoid exploring the same paths again?
// * time complexity O(V + E) which is O(V) = O(mn) and O(E) = (4V) = O(mn), space = O(mn)
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}
	best := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pathLen := dfs(matrix, i, j, cache)
			if pathLen > best {
				best = pathLen
			}
		}
	}
	return best
}

func dfs(matrix [][]int, i int, j int, cache [][]int) int {
	m, n := len(matrix), len(matrix[0])
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	dirs := []int{1, 0, -1, 0, 1}
	for d := 1; d < len(dirs); d++ {
		dx, dy := dirs[d-1], dirs[d]
		nx, ny := i+dx, j+dy
		if nx >= 0 && nx < m && ny >= 0 && ny < n && matrix[nx][ny] > matrix[i][j] {
			pathLen := dfs(matrix, nx, ny, cache)
			if pathLen > cache[i][j] {
				cache[i][j] = pathLen
			}
		}
	}
	cache[i][j]++
	return cache[i][j]
}
