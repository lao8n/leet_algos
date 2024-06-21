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
// * trim is inclusive? yes
// * can right tree child ever be greater than node? no because would have gone left.

// approaches
// * if removing a node then will remove all left-nodes but maybe not all right - can i just take right child and put to left child?
// * in-place dfs? recurse to node if less than replace right child with node and vice-versa?
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// base cases
	if root == nil {
		return nil
	}
	// fmt.Println(node.Val)
	// recursive cases
	if root.Val < low { // need to remove
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
