import "sort"

func threeSum(nums []int) [][]int {
	results := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	if n == 0 || n < 3 {
		return results
	}
	// find target
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// find two numbers that 2-sum to -target
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				results = append(results, []int{nums[left], nums[right], nums[i]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}

	return results
}