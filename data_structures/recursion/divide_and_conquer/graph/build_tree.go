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
// is pre-order not sufficient? i guess we don't know where specifically each element is in the tree.

// approaches
// * so if you take inorder there are multiple valid in-order trees that could be possible and we need to use preorder to pick between them
// * start from leaves or start from node? inorder allows you to split the array in half...
// * rough idea use preorder to pick node -> then split into left and right children using inorder
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 { // same length as inorder
		return nil
	}
	// create node
	nodeVal := preorder[0]
	node := &TreeNode{Val: nodeVal}

	// find left and right children inorder and pre-order have same left-right split
	i := 0
	for i < len(inorder) { // build a hash map rather than searching every tme
		if inorder[i] == nodeVal {
			break
		}
		i++
	}
	node.Left = buildTree(preorder[1:i+1], inorder[:i])   // excludes i
	node.Right = buildTree(preorder[i+1:], inorder[i+1:]) // excludes i
	return node
}
