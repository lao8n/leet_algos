package data_structures

import "sort"

// approaches
// * dp problem - interrelated but no real trade-off
// * sort choose smallest weight always
func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)
	count, sum := 0, 0
	for _, w := range weight {
		sum += w
		if sum <= 5000 {
			count++
		} else {
			break
		}
	}
	return count
}
