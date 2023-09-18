package data_structures

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// comparison in parent or comparison in each child?
// need to accumulate min and max for both trees
func isValidBST(root *TreeNode) bool {
	return isValidBSTAcc(root, math.MinInt, math.MaxInt)
}

func isValidBSTAcc(node *TreeNode, minVal int, maxVal int) bool {
	// base case
	if node == nil {
		return true
	}
	if !(node.Val > minVal && node.Val < maxVal) {
		return false
	}
	// recursive case
	return isValidBSTAcc(node.Left, minVal, node.Val) && isValidBSTAcc(node.Right, node.Val, maxVal)
}
