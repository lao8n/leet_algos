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
// * dfs but how take account of level? want min per level and max per level 0 could have a map?
// * bfs with queue for each level?
type data struct {
	node *TreeNode
	pos  int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// bfs
	queue := []data{{node: root, pos: 0}}
	largestWidth := 0
	for len(queue) > 0 {
		levelSize := len(queue)
		minPos, maxPos := math.MaxInt, math.MinInt
		for i := 0; i < levelSize; i++ {
			popped := queue[i]
			if popped.pos < minPos {
				minPos = popped.pos
			}
			if popped.pos > maxPos {
				maxPos = popped.pos
			}
			if popped.node != nil {
				if popped.node.Left != nil {
					queue = append(queue, data{popped.node.Left, popped.pos * 2})
				}
				if popped.node.Right != nil {
					queue = append(queue, data{popped.node.Right, popped.pos*2 + 1})
				}
			}
		}
		queue = queue[levelSize:]
		if maxPos-minPos+1 > largestWidth {
			largestWidth = maxPos - minPos + 1
		}
	}
	return largestWidth
}
