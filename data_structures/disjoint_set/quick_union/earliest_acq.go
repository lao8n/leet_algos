// approach = union find
// specifics = quick union with 1. union by rank and 2. find memoized
// sort logs? seem to be sorted already
func earliestAcq(logs [][]int, n int) int {
	// setup
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	ds := disjointSet{
		parent:    parent,
		rank:      rank,
		connected: n,
	}

	// sort logs
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	// loop through logs
	for _, log := range logs {
		t, a, b := log[0], log[1], log[2]
		ds.union(a, b)
		if ds.connected == 1 {
			return t
		}
	}
	return -1
}

type disjointSet struct {
	parent    []int
	rank      []int
	connected int
}

func (ds *disjointSet) find(x int) int {
	if x == ds.parent[x] {
		return x
	}
	ds.parent[x] = ds.find(ds.parent[x])
	return ds.parent[x]
}

func (ds *disjointSet) union(x int, y int) {
	rootX, rootY := ds.find(x), ds.find(y)
	if rootX != rootY {
		if ds.rank[rootX] < ds.rank[rootY] {
			ds.parent[rootX] = rootY
		} else if ds.rank[rootY] < ds.rank[rootX] {
			ds.parent[rootY] = rootX
		} else {
			ds.parent[rootX] = rootY
			ds.rank[rootY] += 1
		}
		ds.connected -= 1
	}
}