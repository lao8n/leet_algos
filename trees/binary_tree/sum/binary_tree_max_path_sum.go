package data_structures

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//
func maxPathSum(root *TreeNode) int {
	s := solution{maxPath: math.MinInt32}
	s.maxPathSumRec(root)
	return s.maxPath
}

// post-order traversal of subtree
type solution struct {
	maxPath int
}

// doesn't go through node
func (s *solution) maxPathSumRec(node *TreeNode) int {
	if node == nil {
		return 0
	}
	maxPathLeft := s.maxPathSumRec(node.Left)
	maxPathRight := s.maxPathSumRec(node.Right)
	// assume that path goes through node - there are 4 cases
	// 1. starts at node and goes through left child
	maxPathLeft = max(maxPathLeft, 0)
	// 2. starts at node and goes through right child
	maxPathRight = max(maxPathRight, 0)
	// 3. goes through both left and right children
	// 4. just node
	s.maxPath = max(s.maxPath, node.Val+maxPathLeft+maxPathRight)
	// fmt.Println(s.maxPath, node.Val + max(maxPathLeft, maxPathRight))
	return node.Val + max(maxPathLeft, maxPathRight)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
