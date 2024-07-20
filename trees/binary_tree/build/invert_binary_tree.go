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
// * can you have any empty nodes? currently think not
// approaches
// * dfs - but hard to keep track of where you are - probabloy would need a map - maybe can just do assymmetrical dfs on left and right - same order just backwards
// * bfs - use a queue - all the nodes going left to right in a slice -> then swap values in those nodes? or can you recurse and swap? i.e. do you need memory to store all the values?

// specifics
// * two queues?
// * two choice swap entire nodes -> or create a new node?
// * problem is we need to attach each node to its own respective parent
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	return recurse(root)
}

// each swap is a local swap in each subtree
func recurse(node *TreeNode) *TreeNode {
	// base case
	if node == nil {
		return nil
	}
	node.Left, node.Right = recurse(node.Right), recurse(node.Left)
	return node
}
