package data_structures

// tree definition = 1. all points connected 2. no circular - i.e. no edge between connected components
// need a single root node? number of edges always n - 1?
// approach use union - find to keep track of connected components
func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	ds := disjointSet{
		parent: parent,
		rank:   rank,
	}

	for _, edge := range edges {
		start, end := edge[0], edge[1]
		unionFlag := ds.union(start, end)
		if !unionFlag {
			return false
		}
	}
	return true
}

type disjointSet struct {
	parent []int
	rank   []int
}

// 2 approaches 1. quick union 2. quick find - go for former as can optimize it with 1. union by rank and 2. memoized parent
func (ds *disjointSet) find(x int) int {
	if x == ds.parent[x] {
		return x
	}
	ds.parent[x] = ds.find(ds.parent[x])
	return ds.parent[x]
}

func (ds *disjointSet) union(x int, y int) bool {
	xRoot, yRoot := ds.find(x), ds.find(y)
	if xRoot != yRoot { // if these are equal then cannot be a tree
		// rank
		if ds.rank[xRoot] < ds.rank[yRoot] {
			ds.parent[xRoot] = yRoot
		} else if ds.rank[yRoot] < ds.rank[xRoot] {
			ds.parent[yRoot] = xRoot
		} else {
			ds.parent[xRoot] = yRoot
			ds.rank[yRoot] += 1
		}
		return true
	} else {
		return false
	}
}
