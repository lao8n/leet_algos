/**
 * Definition for a binary tree node.
 * class TreeNode(_value: Int = 0, _left: TreeNode = null, _right: TreeNode = null) {
 *   var value: Int = _value
 *   var left: TreeNode = _left
 *   var right: TreeNode = _right
 * }
 */
// options
// brute force
// for loop over every starting node and every ending node and calculating the length

// dynamic programming - recurse up from the leaves
// wasted computation because a lot of shared effort - solution is memoization
// natural topological ordering from the leaves to the root
// can calculate max length of path from leaf to a node
// problem is might also require a second-pass over the tree to calculate the max sum of effectively two branches - solution might be to return two values (length from tree and max sum)
// basically a dfs (or bfs - don't think it matters)

// greedy approach?
// don't know the future length of tree - so hard to see what this could be 
object Solution {
    def diameterOfBinaryTree(root: TreeNode): Int = {
        // return (longestPath, diameter)
        def dfsDiameterOfBinaryTree(root: TreeNode): (Int, Int) = {
            root match {
                // base case - reached a leaf
                case null => (-1, -1) // because want to have simpler recursive step - also works with (0, 0) and no +2 on line 34
                // recursive case - could have null children and then length is still 0
                case root => 
                    (dfsDiameterOfBinaryTree(root.left), dfsDiameterOfBinaryTree(root.right)) match {
                        // path longest of left or right + 1
                        // diameter longest of sub tree diameters, or lP + rP + 2
                        case ((lP, lD), (rP, rD)) => (Math.max(lP, rP) + 1, Math.max(Math.max(lD, rD), lP + rP + 2))
                    }
            }
        }
        dfsDiameterOfBinaryTree(root)._2 // extract diameter
    }
}