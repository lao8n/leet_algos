package data_structures

// approaches
// * majority element -> boyer-moore - can you just add + 1 if same number and -1 if different

// specifics
// * suppose had majority element and then all other elements. worst case is have another element which is n/2 - 1then it takes n to minus and then n to plus
func majorityElement(nums []int) int {
	maj, count := nums[0], 0
	for _, num := range nums {
		if num == maj {
			count++
		} else {
			count--
			if count <= 0 {
				maj = num
				count = 1
			}
		}
	}
	return maj
}
