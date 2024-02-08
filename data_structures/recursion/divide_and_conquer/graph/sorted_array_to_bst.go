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

// approaches
// * create bst inserting but then need to rebalance afterwards
// * create bst recursively from midpoint - need to be careful with copying vs pointers
func sortedArrayToBST(nums []int) *TreeNode {
	// base case
	if len(nums) == 0 {
		return nil
	}
	// recursive case
	midpoint := len(nums) / 2                 // 4 / 2 = 2 -> 0, 1, 2, 3 if 5 / 2 = 2 which 0, 1, 2, 3, 4 is midpoint
	left := sortedArrayToBST(nums[:midpoint]) // exclusive
	right := sortedArrayToBST(nums[midpoint+1:])
	return &TreeNode{
		Val:   nums[midpoint],
		Left:  left,
		Right: right,
	}
}
