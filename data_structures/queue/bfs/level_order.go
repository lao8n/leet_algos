package data_structures

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// level order traversal -> bfs
// key question is how know whether we have transitioned to new level or not?
// one option is to have a queue with an additional depth field
// another option is to have a nested for loop - for each level
// memoize - no need as tree
func levelOrder(root *Node) [][]int {
	// base cases
	solution := make([][]int, 0)
	if root == nil {
		return solution
	}
	// setup queue
	queue := make([]*Node, 1)
	queue[0] = root

	// bfs
	for len(queue) > 0 {
		levelSize := len(queue) // key trick to separate levels
		level := make([]int, 0)
		// process each level
		for i := 0; i < levelSize; i++ {
			// process current node
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Val)
			// process neighbours
			for _, child := range current.Children {
				queue = append(queue, child)
			}
		}
		solution = append(solution, level)
	}
	return solution
}
