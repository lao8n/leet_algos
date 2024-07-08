package data_structures

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	// start from end
	for i := m + n + 1; i >= 0 && n >= 0; i-- {
		if m < 0 || nums1[m] < nums2[n] {
			nums1[i] = nums2[n]
			n--
		} else {
			nums1[i] = nums1[m]
			m--
		}
	}
}
