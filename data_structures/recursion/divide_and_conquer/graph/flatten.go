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
// * what should the definition of the linked list be?
// * still use tree node -> should update in place?
// approaches
// * dfs add to an external list
// * dfs return - but could be a bit difficult
// * use a stack
func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(node *TreeNode) *TreeNode {
	// base cases
	if node == nil {
		return nil
	}
	// recursive cases
	// node.Right = node.Left (unless nil)
	// node.Left -> recurse to end.Right = node.Right
	leftPath := dfs(node.Left)
	rightPath := dfs(node.Right)

	// update current node's lists
	node.Left = nil
	node.Right = leftPath
	// find end of left path
	right := node
	for right.Right != nil {
		right = right.Right
	}
	right.Right = rightPath
	return node
}
