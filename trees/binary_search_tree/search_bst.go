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
// * just do dfs - forward pass look for val backward pass return pointer
func searchBST(root *TreeNode, val int) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	} else if root.Val < val { // we are too small - should go right
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}
