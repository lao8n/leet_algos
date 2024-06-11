package data_structures

// treat as graph with clone graph dfs - if already visited then just return pointer
// cannot use values for uniqueness so use pointers map from old to new
func copyRandomList(head *Node) *Node {
	visited := make(map[*Node]*Node) // map from old nodes to new
	return copyRandomListRecursive(head, visited)
}

func copyRandomListRecursive(head *Node, visited map[*Node]*Node) *Node {
	// base case
	if head == nil {
		return nil
	}

	// visited
	if n, ok := visited[head]; ok {
		return n // return new node
	}
	// recursive case
	copied := Node{Val: head.Val}
	visited[head] = &copied
	copied.Next = copyRandomListRecursive(head.Next, visited)
	copied.Random = copyRandomListRecursive(head.Random, visited)

	return &copied
}
