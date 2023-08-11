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
// we want to use a queue and bfs but where we go to right to left on each layer?
// need to keep track of each layer too o/w how tell? i guess could use the val?
// make copy or use the same nodes?
// time complexity O(n) - process each node once
// space complexity O(n)
func connect(root *Node) *Node {
	// setup
	if root == nil {
		return nil
	}
	queue := make([]item, 1)
	queue[0] = item{node: root, depth: 0}

	// bfs on queue
	var next item
	for len(queue) > 0 {
		// process node
		head := queue[0]
		queue = queue[1:]
		if next.node != nil && head.depth == next.depth {
			head.node.Next = next.node
		}
		next = head

		// add children
		if head.node.Right != nil {
			// perfect so if right then left also not nil
			queue = append(queue, item{node: head.node.Right, depth: head.depth + 1})
			queue = append(queue, item{node: head.node.Left, depth: head.depth + 1})
		}
	}

	return root
}

// rather than depth could have nested for loop to process each level at a time
type item struct {
	node  *Node
	depth int
}
