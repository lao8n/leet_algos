package data_structures

import "sort"

// clarifying questions
// * 0 obstacle, 1 empty, > 1 height of tree

// approaches
// * bfs where choose shortest tree? can't just be greedy choice because could go through empty cell - need to follow lots of separate paths or backtrack
// * shortest path should be bfs
// * dp but feels like a greedy choice?
// * maybe use greediness where sort trees by height -> then try to construct a path between each pair
// specifics
// * what about scenario where you go to cut a tree then backtrack then go to cut another tree? - what should visited look like in this case?
func cutOffTree(forest [][]int) int {
	m, n := len(forest), len(forest[0])
	// count number of trees
	trees := make([]data, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				trees = append(trees, data{x: 0, y: 0, height: 0}) // add this twice potentially
			}
			if forest[i][j] > 1 {
				trees = append(trees, data{x: i, y: j, height: forest[i][j]})
			}
		}
	}

	sort.Slice(trees, func(i, j int) bool {
		return trees[i].height < trees[j].height
	})

	// greedily choose shorter tree until get blocked
	pathLen := 0
	for i := 0; i < len(trees)-1; i++ {
		fromTree, toTree := trees[i], trees[i+1]
		forest[fromTree.x][fromTree.y] = 1 // cut tree
		visited := make([][]bool, m)
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		partPathLen := bfs(forest, visited, fromTree, toTree)
		if partPathLen == -1 {
			return -1
		}
		pathLen += partPathLen
	}
	return pathLen
}

type data struct {
	x      int
	y      int
	height int
}

type queuer struct {
	data
	pathLen int
}

func bfs(forest [][]int, visited [][]bool, cur data, destination data) int {
	// setup queue
	queue := []queuer{{cur, 0}} // 0 pathLen
	pathLen := 0

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		if visited[popped.x][popped.y] {
			continue // already processed this node
		}
		if popped.x == destination.x && popped.y == destination.y {
			return popped.pathLen
		}
		visited[popped.x][popped.y] = true
		pathLen = popped.pathLen + 1
		for _, neighbour := range findNeighbours(forest, visited, popped.x, popped.y) {
			queue = append(queue, queuer{neighbour, pathLen})
		}
	}
	return -1
}

func findNeighbours(forest [][]int, visited [][]bool, x, y int) []data {
	m, n := len(forest), len(forest[0])
	trans := []int{1, 0, -1, 0, 1}
	neighbours := make([]data, 0)
	for i := 1; i < len(trans); i++ {
		dx, dy := trans[i-1], trans[i]
		nx, ny := x+dx, y+dy
		if nx >= 0 && nx < m && ny >= 0 && ny < n && forest[nx][ny] >= 1 && !visited[nx][ny] {
			neighbours = append(neighbours, data{nx, ny, forest[nx][ny]})
		}
	}
	return neighbours
}
