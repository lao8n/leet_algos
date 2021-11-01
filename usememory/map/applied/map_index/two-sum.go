func twoSum(nums []int, target int) []int {
    nums_map := make(map[int] int)
	for i := range nums {
		complement := target - nums[i]
		if j, ok := nums_map[complement]; ok {
			return []int{i, j}
		} else {
			nums_map[nums[i]] = i
		}
	}
	panic("No two sum solution")
}