package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// clarifying questions
// * only need one? yes
// * cannot be a subtree

// approaches
// * recurse forward with decrementing targetSum - if leaf and remaining amount == 0 return true o/w false
func hasPathSum(root *TreeNode, targetSum int) bool {
	// base case - is leaf with no children
	if root != nil && root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// recursive case - is not leaft
	if root != nil {
		return hasPathSum(root.Left, targetSum-root.Val) ||
			hasPathSum(root.Right, targetSum-root.Val)
	}
	return false
}
