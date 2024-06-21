package data_structures

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var succ *TreeNode

	for root != nil {
		if p.Val >= root.Val {
			root = root.Right
		} else {
			succ = root
			root = root.Left
		}
	}

	return succ
}
