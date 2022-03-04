Disjoint Set Union (DSU) = slowly merge individual points into connected components
- methods: 
  1. find() = two nodes have the same id if in the same component - basically use a tree
  2. union() = connect two components together

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