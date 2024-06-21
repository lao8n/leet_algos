package data_structures

// 1. sort both 2x O(n logn) + 2 pointers 2x O(n)
// 2. map both 2x O(n), loop over one map keys O(n) - 1 or 0
func intersection(nums1 []int, nums2 []int) []int {
	// create maps 2x O(n)
	nums1Set, nums2Set := make(map[int]struct{}, len(nums1)), make(map[int]struct{}, len(nums2))
	for _, num1 := range nums1 {
		nums1Set[num1] = struct{}{}
	}
	for _, num2 := range nums2 {
		nums2Set[num2] = struct{}{}
	}

	// compare
	intersection := []int{}
	for k := range nums1Set {
		if _, ok := nums2Set[k]; ok {
			intersection = append(intersection, k)
		}
	}

	return intersection
}
