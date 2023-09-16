package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// we use a stack rather than recursion
// choice: process nil on value or on children?
func inorderTraversal(root *TreeNode) []int {
	traversal := make([]int, 0)
	stack := make([]*TreeNode, 0)
	top := root

	for top != nil || len(stack) > 0 {
		for top != nil {
			stack = append(stack, top)
			top = top.Left
		}

		// process top of stack
		top = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		traversal = append(traversal, top.Val)
		top = top.Right
	}
	return traversal
}
