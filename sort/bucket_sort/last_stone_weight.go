package data_structures

// use counting sort (or bucket sort with buckets of size 1)
// time complexity O(n + w)
// space complexity O(w)
func lastStoneWeight(stones []int) int {
	// find maximum
	max := 0
	for _, s := range stones {
		if s > max {
			max = s
		}
	}
	// bucket sort
	buckets := make([]int, max+1)
	for _, s := range stones {
		buckets[s] += 1
	}

	// scan through weights
	sIndex, lastWeight := max, 0
	for sIndex > 0 {
		if buckets[sIndex] == 0 {
			sIndex -= 1
		} else if lastWeight == 0 {
			if buckets[sIndex]%2 == 1 {
				lastWeight = sIndex
			}
			sIndex -= 1
		} else {
			buckets[sIndex] -= 1
			if lastWeight-sIndex <= sIndex {
				buckets[lastWeight-sIndex] += 1
				lastWeight = 0
			} else {
				lastWeight -= sIndex
			}
		}
	}
	return lastWeight
}
