package data_structures

// naive approach manually handle indices
// better approach make it the same problem over and over again - by rotating
// can we go row by row or something cleverer?
// maybe we can go backwards from 9 -> 1
// given an index generate a direction - but then need to track whether already been there?
func generateMatrix(n int) [][]int {
	// initialize
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n) // init to 0
	}
	// naive approach
	num := 1
	dir := 0 // 0 right, 1, down, 2 left, 3 up
	cur := coord{}
	for num < n*n {
		m[cur.x][cur.y] = num
		num++
		cur, dir = nextIndex(m, n, cur, dir)
	}
	m[cur.x][cur.y] = num
	return m
}

func nextIndex(m [][]int, n int, cur coord, dir int) (coord, int) {
	// try next index
	next := cur
	switch dir {
	case 0: // go right
		if cur.y+1 < n && m[cur.x][cur.y+1] == 0 { // valid to go right
			next.y = cur.y + 1
		} else {
			return nextIndex(m, n, cur, 1)
		}
	case 1: // go down
		if cur.x+1 < n && m[cur.x+1][cur.y] == 0 { // valid to go right
			next.x = cur.x + 1
		} else {
			return nextIndex(m, n, cur, 2)
		}
	case 2: // go left
		if cur.y-1 >= 0 && m[cur.x][cur.y-1] == 0 { // valid to go right
			next.y = cur.y - 1
		} else {
			return nextIndex(m, n, cur, 3)
		}
	case 3: // go up
		if cur.x-1 >= 0 && m[cur.x-1][cur.y] == 0 { // valid to go right
			next.x = cur.x - 1
		} else {
			return nextIndex(m, n, cur, 0)
		}
	}
	return next, dir
}

type coord struct {
	x int
	y int
}
