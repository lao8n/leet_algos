package data_structures

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		closest = minDist(root.Val, closest, target)
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}

func minDist(x, y int, target float64) int {
	distX := float64(x) - target
	distY := float64(y) - target
	if distX < 0 {
		distX *= -1
	}
	if distY < 0 {
		distY *= -1
	}
	if distX == distY {
		if x < y {
			return x
		}
		return y
	} else if distX < distY {
		return x
	}
	return y
}
