package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// approach
// * two steps 1. convert binary to integers 2. sum
// * sum binary code
// * forward pass -> sum binary digits & backward pass sum totals
// specifics
// * problem don't know what power of 2 to start with? but maybe we just keep increasing power2 acc?

func sumRootToLeaf(root *TreeNode) int {
	return recursive(root, 0)
}

func recursive(node *TreeNode, path int) int {
	path = path * 2
	if node != nil && node.Val == 1 {
		path += 1
	}
	// base cases
	if node == nil {
		return 0
	}
	if node != nil && node.Left == nil && node.Right == nil { // leaf
		return path
	}
	// recursive case
	return recursive(node.Left, path) + recursive(node.Right, path)
}
