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
// * can be any nodes in each tree? yes
// * unique combination? just return true if at least one?

// approaches
// * can we do a parallel dfs? probably not as for each node would need to check all the other nodes
// * intuition is this is 2 sum where build a map of values then look for target - value in the other tree
// * optimise which tree to navigate based upon minimizing size of map? don't bother
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	// create map from tree 1
	map1 := make(map[int]bool)
	makeMap(root1, map1)

	// search for complement in tree 2
	return searchComplement(root2, map1, target)
}

func makeMap(node *TreeNode, map1 map[int]bool) {
	// base case
	if node == nil {
		return
	}
	// recursive case
	map1[node.Val] = true
	makeMap(node.Left, map1)
	makeMap(node.Right, map1)
}

func searchComplement(node *TreeNode, map1 map[int]bool, target int) bool {
	// base case
	if node == nil {
		return false
	}
	// recursive case
	if map1[target-node.Val] { // complement - existence and checking true are same here
		return true
	}
	return searchComplement(node.Left, map1, target) || searchComplement(node.Right, map1, target)
}
