package data_structures

// two pointers l & r
// if sum < target increment left if sum > target decrement right
func twoSum(numbers []int, target int) []int {
    l, r := 0, len(numbers) - 1
    for l <= r {
        sum := numbers[l] + numbers[r]
        if sum == target {
            return []int{l + 1, r + 1}
        } else if sum < target {
            l++
        } else {
            r--
        }
    }
    return []int{} // never get here as guaranteed solution
}