package data_structures

// approaches
// * only a parent knows whether a node is left or right - so need argument to tell it
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	return recursive(root, false)
}

func recursive(node *TreeNode, isLeft bool) int {
	// base case
	if node == nil {
		return 0
	}
	// recursive case
	if isLeft && node.Left == nil && node.Right == nil {
		return node.Val
	}
	left := recursive(node.Left, true)
	right := recursive(node.Right, false)
	return left + right
}
