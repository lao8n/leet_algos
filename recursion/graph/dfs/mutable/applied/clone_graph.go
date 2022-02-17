/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// Choice : dfs vs bfs, dfs because of neighbours and no queue
// Choice: how track already been to neighbor -> use set with value,
// but if pass as acc don't they all have different map
// Choice: how get value back if needed, don't have a set but a map
// but a pointer to Node
// Choice: how update seen map if need neighbours first?

func cloneGraph(node *Node) *Node {
	track := make(map[int]*Node, 100)
	return cloneGraphTrack(node, track)
}

func cloneGraphTrack(node *Node, track map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if n, seen := track[node.Val]; seen {
		return n
	}
	clone := &Node{
		Val: node.Val,
		// don't init neighbours yet
	}
	track[node.Val] = clone
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneGraphTrack(neighbor, track))
	}
	return clone
}