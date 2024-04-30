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
// * does it matter if left or right? no
// * cannot have children? yes

// approaches
// * double recursion 1. func check if matches 2. func to try all possible nodes
// * single recursion? simultaneously do both? but need to keep track of where in sub-tree and where in testing root?
// * forward or backward pass? probably backward pass

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// base case
	if root == nil {
		return false
	}
	// recursive cases
	return subTreeMatch(root, subRoot) || // try with root as current val
		isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func subTreeMatch(node *TreeNode, subNode *TreeNode) bool {
	// base cases
	if node == nil && subNode == nil {
		return true
	}
	// recursive cases
	if node != nil && subNode != nil && node.Val == subNode.Val {
		return subTreeMatch(node.Left, subNode.Left) &&
			subTreeMatch(node.Right, subNode.Right)
	}

	// remaining base cases where one nil and other not nil
	return false
}
