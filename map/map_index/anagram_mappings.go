package data_structures

// which way around? mapping of num2 - i.e. where num2 data is
// 1. create a list over nums 1 and then search for each nums2 element - but this costs O(n^2)
// 2. loop over nums2 building a map for O(n) then loop over nums1 looking up indices
func anagramMappings(nums1 []int, nums2 []int) []int {
	// build map of indices O(n)
	nums2Map := make(map[int]int)
	for i, num2 := range nums2 {
		nums2Map[num2] = i
	}
	// build solution
	solution := make([]int, len(nums1))
	for i, num1 := range nums1 {
		solution[i] = nums2Map[num1]
	}
	return solution
}
