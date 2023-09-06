/**
 * Definition for a binary tree node.
 * class TreeNode(var _value: Int) {
 *   var value: Int = _value
 *   var left: TreeNode = null
 *   var right: TreeNode = null
 * }
 */

object Solution {
    def lowestCommonAncestor(root: TreeNode, p: TreeNode, q: TreeNode): TreeNode = {
        // brute force
        // dfs/bfs all paths until you find a path with p and q on it
        // then extract those paths and go from root -> p & q in parallel until you find they differ
        // problem with this approach is have to recurse the path twice? - and we have to recurse root 
        // to find p & q
        
        // ideally we'd like to find the parents and just recurse upwards but TreeNode doesn't support this
        // operation

        // maybe we can do a dynamic programming bottom-up approach where we have a flag for whether 
        // p & q are children we then recurse the whole tree looking for the first flag where that is true
        def lowestCommonAncestorDFS(node: TreeNode): TreeNode = {
            node match {
                // base case = found leaf value is null
                case null => null // p & q parent
                // recursive case 
                case node => {
                    val leftAncestor = lowestCommonAncestorDFS(node.left)
                    val rightAncestor = lowestCommonAncestorDFS(node.right)
                    (leftAncestor, rightAncestor) match {
                        case (_, _) if p.value == node.value || q.value == node.value => node
                        case (null, null) => null
                        case (null, _) => rightAncestor
                        case (_, null) => leftAncestor
                        case (_, _) => node
                    }
                }
            }
        }
        lowestCommonAncestorDFS(root)
    }
}