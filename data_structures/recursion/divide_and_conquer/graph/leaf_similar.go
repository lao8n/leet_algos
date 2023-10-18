package data_structures

/* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */
// clarifying questions
// * nodes dont' matter?
// approaches
// * traverse in parallel - but a bit tricky as can have two different structures -> try this
// * traverse and build a list and then compare the lists but extra O(n)
// * what type of traversal do you want? in order, pre-orderl post-order etc.
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := traverseTree(root1), traverseTree(root2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := range leaves1 {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}

func traverseTree(node *TreeNode) []int {
	// base cases
	if node == nil {
		return []int{}
	}
	if isLeaf(node) {
		return []int{node.Val}
	}
	// internal node - in order
	return append(traverseTree(node.Left), traverseTree(node.Right)...)
}

func isLeaf(node *TreeNode) bool {
	return node != nil && node.Left == nil && node.Right == nil
}
