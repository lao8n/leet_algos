package data_structures

// approaches
// * for logn need binary search

// specifics - should we search for target, or numbers </> target? search for largest numbers < target, and search for smallest number > target
func searchRange(nums []int, target int) []int {
	left := binarySearch(nums, target, true)
	right := binarySearch(nums, target, false)
	return []int{left, right}
}

func binarySearch(nums []int, target int, lessThan bool) int {
	start, finish := 0, len(nums)-1
	for start <= finish {
		mid := (start + finish) / 2
		// goal
		if nums[mid] == target {
			// we have found first case of target
			if lessThan {
				if mid == 0 || nums[mid-1] < target {
					return mid
				}
				finish = mid - 1
			}
			if !lessThan {
				if mid == len(nums)-1 || nums[mid+1] > target {
					return mid
				}
				start = mid + 1
			}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			finish = mid - 1
		}
	}
	return -1
}
