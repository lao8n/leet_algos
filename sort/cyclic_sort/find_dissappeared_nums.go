package data_structures

import "fmt"

// approaches
// * we know n up front so could create map of n elements and then delete
// * or could create map of disappeared numbers then increment
// * could sort but that is O(n logn)
// * bucket sort or move number to index? swap everything to its index in nums

// specifics
// * how handle duplicate?
func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		for nums[i] != i+1 && nums[i] != nums[nums[i]-1] { // until the same or in correct place
			// fmt.Println(i, nums[i], nums)
			num := nums[i] // e.g. if i == 0 and num = 4
			nums[i], nums[num-1] = nums[num-1], nums[i]
		}
		i++
	}

	fmt.Println(nums)
	result := make([]int, 0)
	for i, num := range nums {
		if i+1 != num {
			result = append(result, i+1)
		}
	}
	return result
}
