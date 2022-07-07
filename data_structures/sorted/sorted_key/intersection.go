package data_structures

import "sort"

// 1. sort both 2x O(n logn) + 2 pointers 2x O(n)
// 2. map both 2x O(n), loop over one map keys O(n) - take minimum of values
func intersection(nums1 []int, nums2 []int) []int {
	// sort - 2x O(n logn)
	sort.Ints(nums1)
	sort.Ints(nums2)

	// 2 pointers - adding to set 2x O(n)
	intersectionSet := make(map[int]struct{})
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			intersectionSet[nums1[i]] = struct{}{}
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	// get keys - O(n)
	intersection := make([]int, len(intersectionSet))
	k := 0
	for n := range intersectionSet {
		intersection[k] = n
		k++
	}
	return intersection
}
