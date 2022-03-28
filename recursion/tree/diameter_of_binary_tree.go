/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// choice 1: validate root at each node or before recursing
// choice 2: separate dfs function? only if need to accumulate something
// choice 3: can we just pick longest path of left or right? maybe need two values -> internal path
// and through path
func diameterOfBinaryTree(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}
	_, internalDiameter := dfs(root)
	return internalDiameter
}

func dfs(root *TreeNode) (throughDiameter int, internalDiameter int) {
	// base case
	if root == nil {
		return throughDiameter, internalDiameter
	}
	// recursive case
	lTD, lID := dfs(root.Left)
	rTD, rID := dfs(root.Right)
	throughDiameter = max(lTD+1, rTD+1)
	internalDiameter = max(lID, rID, lTD+rTD)
	return throughDiameter, internalDiameter
}

func max(args ...int) int {
	maxValue := 0
	for _, a := range args {
		if a > maxValue {
			maxValue = a
		}
	}
	return maxValue
}