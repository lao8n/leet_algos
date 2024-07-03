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

// approaches
// * dfs over all possible paths
// * single path check - but don't know that 5 isn't closest until explore

// specifics
// * do everything on backward pass - calculate distance and see if better than childs best solution
func closestValue(root *TreeNode, target float64) int {
	best, bestDist := 0, math.MaxFloat64
	// always <= so left overrides cur, which overrides right
	if root.Right != nil {
		right := closestValue(root.Right, target)
		if rightDist := dist(right, target); rightDist <= bestDist {
			best = right
			bestDist = rightDist
		}
	}
	if curDist := dist(root.Val, target); curDist <= bestDist {
		best = root.Val
		bestDist = curDist
	}
	if root.Left != nil {
		left := closestValue(root.Left, target)
		if leftDist := dist(left, target); leftDist <= bestDist {
			best = left
			bestDist = leftDist
		}
	}
	return best
}

func dist(val int, target float64) float64 {
	valFloat := float64(val)
	if target > valFloat {
		return target - valFloat
	}
	return valFloat - target
}
