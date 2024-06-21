package data_structures

import "math"

// approaches
// * dp - with start and finish? or just finish?
func maxProduct(nums []int) int {
	// setup dp
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	// calculate products
	max := math.MinInt
	// can optimise further without need for array just constant space
	for i := 0; i < len(nums); i++ { // best seq that ends at i
		if i == 0 {
			dpMax[i] = nums[i]
			max = dpMax[i]
			dpMin[i] = nums[i]
		} else {
			// max
			max1 := nums[i]              // start new seq at i
			max2 := dpMax[i-1] * nums[i] // continue seq to i
			max3 := dpMin[i-1] * nums[i] // take min seq and multiply by i
			if max1 > max2 && max1 > max3 {
				dpMax[i] = max1
			} else if max2 > max3 {
				dpMax[i] = max2
			} else {
				dpMax[i] = max3
			}
			// update max
			if dpMax[i] > max {
				max = dpMax[i]
			}
			// min
			min1 := nums[i]              // start new seq at i
			min2 := dpMin[i-1] * nums[i] // continue seq to i
			min3 := dpMax[i-1] * nums[i] // take max seq and multiply by i
			if min1 < min2 && min1 < min3 {
				dpMin[i] = min1
			} else if min2 < min3 {
				dpMin[i] = min2
			} else {
				dpMin[i] = min3
			}
		}
	}
	return max
}
