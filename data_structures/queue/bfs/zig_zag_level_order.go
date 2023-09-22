package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	solution := [][]int{}
	rightToLeft := false
	for len(queue) > 0 {
		// process node
		levelSize := len(queue)
		levelList := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			popped := queue[0]
			queue = queue[1:]
			if popped != nil {
				queue = append(queue, []*TreeNode{popped.Left, popped.Right}...)
				if rightToLeft {
					levelList = append([]int{popped.Val}, levelList...)
				} else {
					levelList = append(levelList, popped.Val)
				}
			}
		}
		if len(levelList) > 0 {
			solution = append(solution, levelList)
		}
		rightToLeft = !rightToLeft
	}
	return solution
}
