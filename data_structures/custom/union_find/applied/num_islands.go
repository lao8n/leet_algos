package leet_algos

// Union find - 2 methods find and union
// https://leetcode.com/problems/redundant-connection/solution/
type DisjointSetUnion struct {
	Parent []int
	Size   []int
	Count  int // number of connected components
}

func (d *DisjointSetUnion) Init(grid [][]byte) {
	numRows := len(grid)
	numCols := len(grid[0])
	d.Parent = make([]int, numRows*numCols)
	d.Size = make([]int, numRows*numCols)

}
