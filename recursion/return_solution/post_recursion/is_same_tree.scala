/**
 * Definition for a binary tree node.
 * class TreeNode(_value: Int = 0, _left: TreeNode = null, _right: TreeNode = null) {
 *   var value: Int = _value
 *   var left: TreeNode = _left
 *   var right: TreeNode = _right
 * }
 */

// Choices
// recurse down the structure
// 1. breadth first search
// 2. depth first search

// 
object Solution {
    def isSameTree(p: TreeNode, q: TreeNode): Boolean = {
        (p, q) match {
            // base case
            case (null, null) => true
            case (p, null) => false
            case (null, q) => false
            case (p, q) => p.value == q.value && isSameTree(p.left, q.left) && isSameTree(p.right, q.right)
        }
    }
}