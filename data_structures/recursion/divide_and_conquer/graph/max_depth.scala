/**
 * Definition for a binary tree node.
 * class TreeNode(_value: Int = 0, _left: TreeNode = null, _right: TreeNode = null) {
 *   var value: Int = _value
 *   var left: TreeNode = _left
 *   var right: TreeNode = _right
 * }
 */
object Solution {
    def maxDepth(root: TreeNode): Int = {
        // greedy or dynamic programming approach?
        // just need to calculate every distance - could either do an accumulator -> leaves or bottom up
        // latter is more natural because can pick between longer path through tree structure
        root match {
            // base case
            case null => 0
            case root => Math.max(maxDepth(root.left), maxDepth(root.right)) + 1
        }
    }
}