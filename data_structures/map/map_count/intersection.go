// cannot use set instead need a map to track counts - can do it in a single pass
// could also sort
func intersect(nums1 []int, nums2 []int) []int {
    // create map with counts
    nums1Map := make(map[int]int, len(nums1))
    for _, num1 := range nums1 {
        nums1Map[num1]++
    }
    
    // single pass over second array
    intersection := []int{}
    for _, num2 := range nums2 {
        if count, ok := nums1Map[num2]; ok {
            if count >= 1 {
                intersection = append(intersection, num2)
            }
            nums1Map[num2]--
        }
    }
    return intersection
}ß