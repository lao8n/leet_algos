package data_structures

// union find
// two approaches 1. quick union 2. quick find
// better to go with quick union as can optimise with 1. memoized find operation 2. union by rank
// questions
// * info is duplicated so can we avoid unioning twice?
func findCircleNum(isConnected [][]int) int {
	rootSlice := make([]int, len(isConnected))
	rankSlice := make([]int, len(isConnected))
	for i := 0; i < len(rootSlice); i++ {
		rootSlice[i] = i
		rankSlice[i] = 1
	}
	ds := disjointSet{
		root: rootSlice,
		rank: rankSlice,
	}
	for i, city := range isConnected {
		for j, connection := range city { // more efficient is to just check half the table
			if i != j && connection == 1 {
				ds.union(i, j)
			}
		}
	}
	// check for number of roots i.e. provinces
	numProvinces := 0
	for i, parent := range ds.root {
		if i == parent { // means it is a root node
			numProvinces++
		}
	}
	return numProvinces
}

type disjointSet struct {
	root []int // index is node, value is parent
	rank []int // rank is height of the tree as opposed to number of components
}

func (ds *disjointSet) find(x int) int {
	if x == ds.root[x] {
		return x
	}
	ds.root[x] = ds.find(ds.root[x])
	return ds.root[x]
}

func (ds *disjointSet) union(x int, y int) {
	xRoot, yRoot := ds.find(x), ds.find(y)
	if xRoot != yRoot {
		if ds.rank[xRoot] < ds.rank[yRoot] {
			ds.root[xRoot] = yRoot // make smaller rank change root
		} else if ds.rank[yRoot] < ds.rank[xRoot] {
			ds.root[yRoot] = xRoot
		} else {
			ds.root[xRoot] = yRoot
			ds.rank[yRoot] += 1
		}
	}
}
