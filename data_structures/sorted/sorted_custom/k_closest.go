package data_structures

import "sort"

// approach 1:
// 1. calculate distances O(n) 2. sort distances O(n logn) 3. pick k first O(1)
// approach 2:
// 1. calculate distances O(n) 2. same time insert into heap O(n logn) 3. remove element from heap O(n logn)

func kClosest(points [][]int, k int) [][]int {
	// 1. sort by distances directlyO(n logn)
	sort.SliceStable(points, func(i, j int) bool {
		iDistance := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		jDistance := points[j][0]*points[j][0] + points[j][1]*points[j][1]
		if iDistance < jDistance {
			return true
		}
		return false
	})

	// 2. pick first k
	return points[:k]
}
