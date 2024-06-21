package data_structures

import "math"

// dp problem (technically bellman-ford)
// don't care about cumulative length of path etc. so any given square pick the route that gets there with minimal effort
// therefore can do dp on entirely new square (maybe can optimise with min effort to get to that square
// unclear on how to know if path is legitimate...e.g. route goes backwards...cannot just have a rule which is always down or right
// actually i'm really using dijkstra's here - although no heap
func minimumEffortPath(heights [][]int) int {
	numRows, numCols := len(heights), len(heights[0])
	dp := make([][]int, numRows)
	for i := range dp {
		dp[i] = make([]int, numCols)
		for j, _ := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = 0
	queue := make([]xy, 1)
	queue[0] = xy{x: 0, y: 0}
	k := numRows * numCols // max length of path
	for k > 0 {
		ql := len(queue)
		for i := 0; i < ql; i++ {
			popped := queue[0]
			queue = queue[1:]
			// process neighbours (so know the direction)
			for _, neighbour := range getNeighbours(popped, numRows, numCols) {
				diff := heights[neighbour.x][neighbour.y] - heights[popped.x][popped.y]
				if diff < 0 {
					diff *= -1
				}
				if dp[popped.x][popped.y] > diff {
					diff = dp[popped.x][popped.y]
				}
				if diff < dp[neighbour.x][neighbour.y] {
					dp[neighbour.x][neighbour.y] = diff
					queue = append(queue, neighbour)
				}
			}
		}
		k--
	}
	return dp[numRows-1][numCols-1]
}

type xy struct {
	x int
	y int
}

func getNeighbours(curr xy, numRows int, numCols int) []xy {
	translation := []int{1, 0, -1, 0, 1}
	neighbours := make([]xy, 0)
	for i := 1; i < len(translation); i++ {
		newX, newY := curr.x+translation[i-1], curr.y+translation[i]
		if newX >= 0 && newX < numRows && newY >= 0 && newY < numCols {
			neighbours = append(neighbours, xy{x: newX, y: newY})
		}
	}
	return neighbours
}
