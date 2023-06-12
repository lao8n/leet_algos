Disjoint Set Union (DSU) 
Key Idea = turn undirected graph into a directed graph with a single root, then compare roots to see if connected
Approach = slowly merge individual points into connected components
- methods: 
  1. find() = two nodes have the same id if in the same component i.e. the same root
  2. union() = connect two components together

2 approaches
1. Quick Find: O(1) find but O(n) union -> store root so union is more expensive but easy to see what root is
2. Quick Union: O(1) union but O(n) find -> store parent so have to recurse to root, but faster union

Best approach
* Quick Union = Overall time complexity of Quick Find is O(n^2) (O(n) per element), versus Quick Union <= O(n^2) (O(n) is the worst case)

Analogy
* have a room full of people, how to see if people are related? (assume only one parent e.g. just father's side)
* one approach is to have every person list their 10th paternal ancestor - i.e. quick find
* other approach is just to have every person list their direct father.

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
2. Union-by-rank = choose whether to `parent[x] = parent[y]` or `parent[y] = parent[x]` based upon smaller joining larger. Only relevant for Quick Union.

Complexity
* Time: `O(mxn)` required by the time complexity where at most `mxn` islands need to be unioned.
* Space: `O(mxn)` required by the `UnionFind` `Parent` and `Size` are `mxn` 