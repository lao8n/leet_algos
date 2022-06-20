package data_structures

// approach 1: merge O(log(m+n)) & find median O((m+n)/2)
// approach 2: get lens O(m) + O(n) & then merge half-way O(log((m+n)/2))
// complication 1: even vs odd numbers
// complication 2: switching to other array
// choice 1 : recursion or for loop
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// get median index
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)
	// odd numbers (1 + 2)/2 = 1.5 -> 1 okay because index start at 0
	// even numbers (2 + 2)/2 = 2 -> 2 okay because it is the highest index with start at 0 but also need 1
	medianIndex := int((lenNums1 + lenNums2) / 2)
	var medianValue float64

	for f, i, j := 0, 0, 0; f <= medianIndex; f++ {
		// increment through values
		var v int
		if i < lenNums1 && j < lenNums2 && nums1[i] < nums2[j] {
			v = nums1[i]
			i++
		} else if j < lenNums2 {
			v = nums2[j]
			j++
		} else {
			v = nums1[i]
			i++
		}
		// calc median value
		if (lenNums1+lenNums2)%2 == 0 {
			if f == medianIndex-1 {
				medianValue = float64(v)
			} else if f == medianIndex {
				medianValue = (medianValue + float64(v)) / 2
			}
		} else {
			if f == medianIndex {
				medianValue = float64(v)
			}
		}
	}
	return medianValue
}
