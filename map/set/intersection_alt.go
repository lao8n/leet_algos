package data_structures

// 1. sort both 2x O(n logn) + 2 pointers 2x O(n)
// 2. map both 2x O(n), loop over one map keys O(n) - 1 or 0
func intersection(nums1 []int, nums2 []int) []int {
	// create maps 2x O(n)
	nums1Set := make(map[int]struct{}, len(nums1))
	for _, num1 := range nums1 {
		nums1Set[num1] = struct{}{}
	}

	// compare
	intersection := []int{}
	for _, num2 := range nums2 {
		if _, ok := nums1Set[num2]; ok {
			intersection = append(intersection, num2)
			delete(nums1Set, num2)
		}
	}

	return intersection
}
