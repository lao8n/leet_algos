Time Complexity = O(E+V) = O(2E) + O(V) because every node is visited once and every edge is considered twice
Space Complexity = O(V) - max size of stack

Use when:
* Path finding (but don't care about shortest path) as more space efficient
* For pre-order and post-order traversal
* Topological sorting
* Wider graphs (as opposed to deeper)

Choice between
* naive 
* cache per search
* cache over all searches (memoization)

Algorithm
```
func DFS(cur node, target node, visited map[int]bool) bool {
    if cur.id == target.id {
        return true
    }
    for _, node := range cur.neighbours {
        if !visited[node.id] {
            visited[node.id] = true
            if DFS(node, target, visited) {
                return true
            }
        }
    }
    return false
}

type node struct {
    id int
    neighbours []node
}
```

Under the hood this is just using the call stack