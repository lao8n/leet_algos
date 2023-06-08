Disjoint Set Union (DSU) 
Key Idea = turn undirected graph into a directed graph with a single root, then compare roots to see if connected
Approach = slowly merge individual points into connected components
- methods: 
  1. find() = two nodes have the same id if in the same component i.e. the same root
  2. union() = connect two components together

2 approaches
1. Quick Find: O(1) find but O(n) union -> store root so union is more expensive but easy to see what root is
2. Quick Union: O(1) union but O(n) find -> store parent so have to recurse to root, but faster union

Naive find implementation
```
func find(x: Node) {
    for parent[x] != x {
        x = parent[x]
    }
    return x
}
```
Naive union implementation
```
func union(x, y) {
    parent[find(x)] = find(y)
}
```
Improvements
1. Path compression = replace with `parent[x] = find(parent[x])` to remember calculation 
2. Union-by-rank = choose whether to `parent[x] = parent[y]` or `parent[y] = parent[x]` based upon smaller joining larger

Complexity
* Time: `O(mxn)` required by the time complexity where at most `mxn` islands need to be unioned.
* Space: `O(mxn)` required by the `UnionFind` `Parent` and `Size` are `mxn` 