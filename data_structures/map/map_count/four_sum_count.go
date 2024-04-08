package data_structures

// clarifying questions
// * do have negative numbers so cannot rule anything out.

// approaches
// * two pointer approach sorting all array O(4 * n logn) +
// * map approach looking for complements - need to brute force everything but final map
// * brute force 2d array of possible sums - could have a map with count... O(n^2) and do it again for O(n^2) - and then loop over O(3 n^2) looking for complements...
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// first combined map
	map1 := make(map[int]int, 0)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			map1[num1+num2]++
		}
	}
	// second combined map
	map2 := make(map[int]int, 0)
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			map2[num3+num4]++
		}
	}
	// look for complements
	total := 0
	for sum1, count1 := range map1 {
		if count2, ok := map2[-sum1]; ok {
			total += count1 * count2
		}
	}
	return total
}
