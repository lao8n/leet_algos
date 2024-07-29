package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// approaches
// * in-order traversal via dfs

// specifics?
// * how update? in-place or new node? return node and might be difficult to replace otherwise
func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	d := data{increasing: dummy}
	d.dfs(root)
	return dummy.Right
}

type data struct {
	increasing *TreeNode
}

func (d *data) dfs(node *TreeNode) {
	// base case
	if node == nil {
		return
	}
	// recursive case
	// 1. go left
	d.dfs(node.Left)
	// 2. process node
	d.increasing.Right = &TreeNode{
		Val: node.Val,
	}
	d.increasing = d.increasing.Right
	// 3. go right
	d.dfs(node.Right)
}
