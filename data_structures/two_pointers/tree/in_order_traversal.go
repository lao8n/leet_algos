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
// * there are 3 types of traversal: pre-, in-order and post-.. which is determined by the order that current and children are traversed
// * also a choice around how we build up the list -> could return directly or accumulate
// * forward pass is finding what to work on... backward pass is actually doing the work
// * morris method and threaded binary tree
func inorderTraversal(root *TreeNode) []int {
	solution := make([]int, 0)
	cur := root
	var pre *TreeNode
	for cur != nil {
		if cur.Left != nil {
			// find right-most node of left branch and have it follow onto cur
			pre = cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur

			// go into that subtree and process that
			temp := cur
			cur = cur.Left
			temp.Left = nil
		} else {
			// no left so process cur and right
			solution = append(solution, cur.Val)
			cur = cur.Right
		}
	}
	return solution
}
