package recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	// first time both branches non-nil this is common ancestor
	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	} else if r != nil {
		return r
	} else {
		return nil
	}
}
