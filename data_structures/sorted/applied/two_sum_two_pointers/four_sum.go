import "sort"

func fourSum(nums []int, target int) [][]int {
	// 1. 2 for loops + 2-sum hashmap O(n^3)
	// 2. 2 for loops + 2-sum 2-pointers O(n^3)
	n := len(nums)
	if nums == nil || n < 4 {
		return nil
	}
	sort.Ints(nums)
	var solution [][]int
	// two for loops
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			// 2-pointer
			remainingTarget := target - nums[i] - nums[j]
			left, right := j+1, n-1
			for left < right {
				sum := nums[left] + nums[right]
				if sum < remainingTarget {
					left++
				} else if sum > remainingTarget {
					right--
				} else {
					quadruplet := []int{nums[i], nums[j], nums[left], nums[right]}
					solution = append(solution, quadruplet)
					// skip through duplicates
					for left < right && nums[left] == quadruplet[2] {
						left++
					}
					for left < right && nums[right] == quadruplet[3] {
						right--
					}
				}
			}
			// skip through duplicates
			for j < n-1 && nums[j] == nums[j+1] {
				j++
			}
		}
		// skip through duplicates
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return solution
}