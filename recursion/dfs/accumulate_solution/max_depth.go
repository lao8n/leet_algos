package recursion

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// accumulator is much faster because it is tail recursive
// https://medium.com/@meeusdylan/tail-recursion-in-go-fb5cf69a0f26
// even though no tail-call optimisation
func maxDepth(root *TreeNode) int {
	return maxDepthRecursion(root, 0)
}

func maxDepthRecursion(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}

	newDepth := depth + 1
	lDepth := maxDepthRecursion(root.Left, newDepth)
	rDepth := maxDepthRecursion(root.Right, newDepth)

	if lDepth > rDepth {
		return lDepth
	}
	return rDepth
}
