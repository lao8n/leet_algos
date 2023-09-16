package data_structures

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
	if root != nil {
		connectNodes(root.Left, root.Right)
	}
	return root
}

func connectNodes(left *Node, right *Node) {
	// base cases
	if left == nil {
		return
	}
	left.Next = right
	// recursive step
	connectNodes(left.Left, left.Right)
	connectNodes(left.Right, right.Left) // across trees
	connectNodes(right.Left, right.Right)
}
