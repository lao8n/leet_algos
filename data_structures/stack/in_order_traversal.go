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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 1)
	stack[0] = root
	seen := make(map[*TreeNode]bool, 0)
	traversal := make([]int, 0)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Left != nil && !seen[top.Left] {
			stack = append(stack, top)
			stack = append(stack, top.Left)
		} else if (top.Left == nil || seen[top.Left]) && !seen[top] {
			stack = append(stack, top)
			traversal = append(traversal, top.Val)
			seen[top] = true
		} else if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return traversal
}
