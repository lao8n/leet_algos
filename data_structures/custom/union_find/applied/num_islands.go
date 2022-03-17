func numIslands(grid [][]byte) int {
	numRows := len(grid)
	numCols := len(grid[0])
	uf := new(UnionFind)
	uf.Init(grid)
	for r := 0; r < numRows; r++ {
		for c := 0; c < numCols; c++ {
			if grid[r][c] == '1' {
				grid[r][c] = '0'
				current := r*numRows + c
				if isValid(r-1, c, grid) {
					uf.Union(current, r*(numRows-1)+c)
				}
				if isValid(r+1, c, grid) {
					uf.Union(current, r*(numRows+1)+c)
				}
				if isValid(r, c-1, grid) {
					uf.Union(current, r*numRows+c-1)
				}
				if isValid(r, c+1, grid) {
					uf.Union(current, r*numRows+c+1)
				}
			}
		}
	}
	return uf.Count
}

func isValid(x, y int, grid [][]byte) bool {
	numRows := len(grid)
	numCols := len(grid[0])
	// check boundaries
	if x >= 0 && x < numRows && y >= 0 && y < numCols {
		// check if land
		if grid[x][y] == '1' {
			return true
		}
	}
	return false
}

type UnionFind struct {
	Parent []int
	Size   []int
	Count  int // number of connected components
}

func (u *UnionFind) Init(grid [][]byte) {
	numRows := len(grid)
	numCols := len(grid[0])
	u.Parent = make([]int, numRows*numCols)
	u.Size = make([]int, numRows*numCols)

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			current := i*numRows + j
			if grid[i][j] == '1' {
				u.Parent[current] = current
				u.Size[current] = 1
				u.Count++
			} else {
				u.Parent[current] = -1 // '0' has node's parent -1
			}
		}
	}
}

func (u *UnionFind) Union(x, y int) {
	pX, pY := u.Find(x), u.Find(y)
	if pX == pY {
		return
	}
	// add smaller to larger
	if u.Size[pX] < u.Size[pY] {
		u.Parent[pX] = u.Parent[pY]
		u.Size[pY] += u.Size[pX]
	} else {
		u.Parent[pY] = u.Parent[pX]
		u.Size[pX] += u.Size[pY]
	}
}

func (u *UnionFind) Find(x int) int {
	for u.Parent[x] != x {
		fmt.Println(u.Parent[x], x)
		u.Parent[x] = u.Find(u.Parent[x])
		x = u.Parent[x]
	}
	return u.Parent[x]
}