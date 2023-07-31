package data_structures

import "sort"

// approach 1 = just greedily swap to smaller - problem is might not put smallest at front if indirect connection
// approach 2 = separate into connected components and sort each component - but maintain indices of component to know how to combine... maybe can do it as part of union process?

// time complexity: O((E+V)⋅α(V)+Vlog⁡V)
func smallestStringWithSwaps(s string, pairs [][]int) string {
	// setup
	parent := make([]int, len(s))
	rank := make([]int, len(s))
	// O(n)
	for i := 0; i < len(s); i++ {
		parent[i] = i
		rank[i] = 1
	}
	ds := disjointSet{
		parent: parent,
		rank:   rank,
	}
	// O(n ^ 2) (because union is O(n) because of find O(1) reverse ackermann)
	for _, pair := range pairs {
		swapLeft, swapRight := pair[0], pair[1]
		ds.union(swapLeft, swapRight)
	}
	// brute force is loop through and extract each connected component sort them
	// and then loop through again to place them?
	stringSegments := make(map[int][]byte, 0) // int is root
	// O(n) (find is O(1))
	for i := 0; i < len(s); i++ {
		root := ds.find(i)
		stringSegments[root] = append(stringSegments[root], s[i])
	}
	// O(n) * O(n logn)?
	for k, v := range stringSegments {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		stringSegments[k] = v
	}
	solution := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		root := ds.find(i)
		solution[i] = stringSegments[root][0]
		stringSegments[root] = stringSegments[root][1:]
	}
	return string(solution)
}

type disjointSet struct {
	parent []int
	rank   []int
}

func (ds *disjointSet) find(x int) int {
	if ds.parent[x] == x {
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
	}
}
