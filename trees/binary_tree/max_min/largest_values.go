package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// approach
// * dfs but need a depth acc
// * bfs pick max for each row
func largestValues(root *TreeNode) []int {
	max := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		// base case
		if node == nil {
			return
		}
		// update max
		if len(max) >= depth+1 {
			if node.Val > max[depth] {
				max[depth] = node.Val
			}
		} else {
			max = append(max, node.Val)
		}

		// recurse
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return max
}
