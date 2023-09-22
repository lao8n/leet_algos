package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	solution := [][]int{}
	for len(queue) > 0 {
		// process node
		levelSize := len(queue)
		levelList := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			popped := queue[0]
			queue = queue[1:]
			if popped != nil {
				levelList = append(levelList, popped.Val)
				queue = append(queue, popped.Left)
				queue = append(queue, popped.Right)
			}
		}
		if len(levelList) > 0 {
			solution = append(solution, levelList)
		}
	}
	return solution
}
