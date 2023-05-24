package recursion

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Choice: check current or check children for nil
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
