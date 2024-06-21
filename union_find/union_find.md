**Union Find**

Definition = Union Find (also Disjoint Set) is a data structure to keep track of a set of elements partitioned into disjoint groups. The underlying idea is to turn an undirected graph into a set of directed graphs with a single root, which can then be used as a test for uniqueness. It is called union-find after the two key operations. 

The approach involves slowly merging individual points into connected components where `find()` is a way to get the root node of a node and `union()` is a way to connect two components by setting one of the roots to have the other as its root.

In the data structure, itself there is a choice between storing the direct parent of each node known as quick union, or storing the root parent known as quick find.

Time complexity = Quick union is best because it can be optimised is `O(a n)<= O(n^2)` versus quick find which is `O(n^2)`. `a n` is the Inverse Ackermann function which is in practice `O(1)` on average

Space complexity = `O(n)`

**Types**

**Quick find**
```
// quick_find.cpp
int find(int x) {
    return root[x];
}
void unionNodes(int x, int y){
    ...
    if (rootX != rootY) {
        for (int i = 0; i < root.size(); i++){
            if (root[i] == rootY){
                root[i] = rootX;
            }
        }
    }
}
```
Quick find has `O(1)` find but `O(n)` union because need to update every node to new root parent.

**Quick union**
```
func find(x: Node) {
    for parent[x] != x {
        x = parent[x]
    }
    return x
}
func union(x, y Node) {
    parent[find(x)] = find(y)
}
```
The naive approach to quick union is to set one of the components to have the other component as its root - this is quick union after all. But to do so need to loop through each generation of parents to find the root node. This can be optimised in two ways where:
1. Path compression = replace with `parent[x] = find(parent[x])` to remember calculation 
2. Union-by-rank = choose whether to `parent[x] = parent[y]` or `parent[y] = parent[x]` based upon smaller joining larger. 
```
// number_provinces.go
type disjointSet struct {
    root []int
    rank []int
}
func (ds *disjointSet) find(x int) int {
    if x == ds.root[x] { return x }
    ds.root[x] = ds.find(ds.root[x])
    return ds.root[x]
}
func (ds *disjointSet) union(x int, y int) {
    xRoot, yRoot := ds.find(x), ds.find(y)
    if xRoot != yRoot {
        if ds.rank[xRoot] < ds.rank[yRoot] {
            ds.root[xRoot] = yRoot
        } else if ds.rank[yRoot] < ds.rank[xRoot] {
            ds.root[yRoot] = xRoot
        } else {
            ds.root[xRoot] = yRoot
            ds.rank[yRoot] += 1
        }
    }
}
```
In `number_provinces.go` the `root` and `rank` slices need to be initialized with `rootSlice[i] = i` and `rankSlice[i] = 1`, and then need to loop over all the connections (i.e. edges) `ds.union(i, j)` and then afterwards check to count the number of nodes that are its own parent - which indicates that it is a separate node, as is also done in `num_islands.go` and `count_components.go`, this can be optimised by starting with `connected: n` and then decrementing everytime there is a union.

Variations include `valid_tree.go` that looks for whether a single component. More complicated variations include `smallest_string_with_swaps.go` which uses allowed swaps to get the lexicographically smallest string, or return a time when all components connected in `earliest_acq.go`. Or even `calc_equation.go` nodes have a `weight float64` as well as a `root` and `rank` which is used to `calculateQuery()`


**Quick union: Kruskal's**
```
// min_cost_connect_points.go
sort.Slice(edges, func(i int, j int) bool {
    return edges[i].cost < edges[j.cost]
})
...
for _, edge := range edges {
    if numEdges == len(points) - 1 {
        return totalCost
    }
    if !ps.connected(edge.i, edge.j) {
        ps.union(edge.i, edge.j)
        numEdges += 1
        totalCost += edge.cost
    }
}
```
Quick union typically assumes that we have an unweighted graph. Kruskal's algorithm is for graphs with weights where the core approach is to 1. add the smallest edge that doesn't create a cycle 2. until all nodes are connected. Could potentially use a min-heap to sort but often easier just to sort the whole slice, the union-find data structure is for cycle detection and the key trick is that you can stop at `n - 1` edges. A complicated variation of this is `min_cost_to_supply_water.go` where you don't just have edges but also wells - the key trick there is to create an additional node to a well with each node having an edge and cost associated with it `pipes = append(pipes, []int{ 0, i + 1, well})`.