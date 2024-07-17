package data_structures

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
// clarifying questions
// * leaf = all 1s or all 0s
// * o/w = set val to any values and define grid into 4 sub grids -> recurse

// approach
// * manually build quad tree

// specifics
// * should node be current node or previous node?
// * how optimise bailing out of leaf calculation?
func construct(grid [][]int) *Node {
	return recurse(coord{0, 0}, coord{len(grid) - 1, len(grid[0]) - 1}, grid)
}

type coord struct {
	x int
	y int
}

func recurse(topLeft, bottomRight coord, grid [][]int) *Node {
	if isLeaf(topLeft, bottomRight, grid) {
		return &Node{
			Val:    grid[topLeft.x][topLeft.y] == 1,
			IsLeaf: true,
		}
	} else {
		// 4 recursions
		midX, midY := int((topLeft.x+bottomRight.x)/2), int((topLeft.y+bottomRight.y)/2)
		return &Node{
			Val:         grid[topLeft.x][topLeft.y] == 1,
			IsLeaf:      false,
			TopLeft:     recurse(topLeft, coord{midX, midY}, grid),
			TopRight:    recurse(coord{topLeft.x, midY + 1}, coord{midX, bottomRight.y}, grid),
			BottomLeft:  recurse(coord{midX + 1, topLeft.y}, coord{bottomRight.x, midY}, grid),
			BottomRight: recurse(coord{midX + 1, midY + 1}, bottomRight, grid),
		}
	}
}

func isLeaf(topLeft, bottomRight coord, grid [][]int) bool {
	same := grid[topLeft.x][topLeft.y]
	for i := topLeft.x; i <= bottomRight.x; i++ {
		for j := topLeft.y; j <= bottomRight.y; j++ {
			if grid[i][j] != same {
				return false
			}
		}
	}
	return true
}
