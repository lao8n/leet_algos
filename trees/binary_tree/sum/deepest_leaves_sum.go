package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 2 approaches
// 1. dfs where track depth as descend - always return deepest depth if choice
// 2. bfs where loop through until final round and then sum those leaves
func deepestLeavesSum(root *TreeNode) int {
	returnSum, _ := deepest(root, 0)
	return returnSum
}

func deepest(root *TreeNode, depth int) (int, int) { // return sum & depth
	// base case
	if root == nil {
		return 0, depth // don't add to depth because nil
	}

	// forward pass
	lSum, lDepth := deepest(root.Left, depth+1)
	rSum, rDepth := deepest(root.Right, depth+1)
	// backwards pass - choosest deepest depth
	returnSum, returnDepth := 0, 0
	if lDepth > rDepth {
		returnDepth = lDepth
		returnSum = lSum
	} else if rDepth > lDepth {
		returnDepth = rDepth
		returnSum = rSum
	} else if lSum == 0 && rSum == 0 { // we are at leaf
		returnDepth = depth
		returnSum = root.Val
	} else { // we are higher up tree but are merging two equal depth leaves
		returnDepth = lDepth
		returnSum = lSum + rSum
	}
	return returnSum, returnDepth
}
