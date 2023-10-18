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
// * can stop and start earlier than root
// * at node 3 in the left branch there are paths of 3-5 3-5-10 that can be added

// approaches
// * accumulate or backward pass paths? accumulate
// * as what? map or list? so at 3 have 3, 8, 18, advantage of map is can check if target is hit -> need to be careful about new maps etc.
// * we want different maps for each pass so either create a new map or we backtrack?
func pathSum(root *TreeNode, targetSum int) int {
	initPath := make(map[int]int)
	return pathSumAcc(root, targetSum, initPath)
}

func pathSumAcc(node *TreeNode, targetSum int, paths map[int]int) int {
	// base cases
	if node == nil {
		return 0
	}
	// process node
	extendedPaths := make(map[int]int)
	for path, count := range paths {
		extendedPaths[path+node.Val] += count
	}
	extendedPaths[node.Val] += 1

	numWays := 0
	if ways, ok := extendedPaths[targetSum]; ok {
		numWays += ways
	}
	numWays += pathSumAcc(node.Left, targetSum, extendedPaths)
	numWays += pathSumAcc(node.Right, targetSum, extendedPaths)

	return numWays
}
