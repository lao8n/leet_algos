package data_structures

// basically just bfs where we use a queue of neighbours
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	queue := make([][]int, 1)
	queue[0] = []int{sr, sc}
	startingColor := image[sr][sc]

	// base case
	if startingColor == color {
		return image
	}
	// queue case
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		x, y := popped[0], popped[1]
		if image[x][y] == startingColor {
			image[x][y] = color
			neighbours := findNeighbours(image, x, y)
			queue = append(queue, neighbours...)
		}
	}
	return image
}

// do not need additional memoization matrix because can just use updated image matrix
// to be a valid neighbour needs to 1. be on grid 2. not be 0 or color already (do above)
func findNeighbours(image [][]int, r int, c int) [][]int {
	numRows := len(image)
	numCols := len(image[0])

	possNeighbours := [][]int{}
	translations := []int{1, 0, -1, 0, 1}
	for i := 1; i < len(translations); i++ {
		vertical, horizontal := translations[i-1], translations[i]
		newR, newC := r+vertical, c+horizontal
		// on grid
		if newR >= 0 && newR < numRows && newC >= 0 && newC < numCols {
			possNeighbours = append(possNeighbours, []int{newR, newC})
		}
	}
	return possNeighbours
}
