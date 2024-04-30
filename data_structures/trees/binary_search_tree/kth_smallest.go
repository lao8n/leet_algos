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
// * in bst the in-order traversal represents the order of all the elements so 1. generate entire in-order traversal 2. look up kth index for value
// * just do in-order traversal with a counter..
// * could do it with a stack or with a accumulated counter but need a pointer - easier with recursion

// implementation
// * k is 1-indexed
func kthSmallest(root *TreeNode, k int) int {
	c := counter{count: 0, k: k}
	return c.kthSmallest(root)
}

type counter struct {
	count int
	k     int
}

func (c *counter) kthSmallest(node *TreeNode) int {
	// base cases
	if node == nil {
		return 0
	}

	// recursive
	left := c.kthSmallest(node.Left) // go left
	if left != 0 {
		return left
	}
	c.count++ // only increment on backward pass
	if c.count == c.k {
		return node.Val
	}
	right := c.kthSmallest(node.Right)
	if right != 0 {
		return right
	}
	return 0
}
