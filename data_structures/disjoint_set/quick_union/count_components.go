package data_structures

// approach: use union find
// quick union + 1. union by rank and 2. memoized find
func countComponents(n int, edges [][]int) int {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	ds := disjointSet{
		parent:     parent,
		rank:       rank,
		components: n,
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		ds.union(from, to)
	}
	return ds.components
}

type disjointSet struct {
	parent     []int
	rank       []int
	components int
}

func (ds *disjointSet) find(x int) int {
	if x == ds.parent[x] {
		return x
	}
	ds.parent[x] = ds.find(ds.parent[x])
	return ds.parent[x]
}

func (ds *disjointSet) union(x int, y int) {
	xRoot, yRoot := ds.find(x), ds.find(y)
	if xRoot != yRoot {
		if ds.rank[xRoot] < ds.rank[yRoot] {
			ds.parent[xRoot] = yRoot
		} else if ds.rank[yRoot] < ds.rank[xRoot] {
			ds.parent[yRoot] = xRoot
		} else {
			ds.parent[xRoot] = yRoot
			ds.rank[yRoot] += 1
		}
		ds.components -= 1
	}
}
