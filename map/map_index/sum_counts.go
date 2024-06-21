package data_structures

// clarifying questions
// * how is it different from the other question? it is not just need better time complexity
// approaches
// * for the basic approach we had two for loops 1. all subarray lengths and 2. for building the rolling map
// * instead we can have a single loop where for each number 
// consider case where we have [1, 2]
// this means we have 3 possible subarrays [1], [2] and [1, 2] where the sum of distinct values is 1 + 1 + 2 = 4 where adding the 2 made the additional 3
// don't udnerstand these formulas
func sumCounts(nums []int) int {
    mod := 1000000007
    // init 
    s1 := make([]int, len(nums)) // sum
    s1[0] = 1
    s2 := make([]int, len(nums)) // sum of squares
    s2[0] = 1
    lastIndex := make(map[int]int) // num to index
    lastIndex[nums[0]] = 0
    
    for i, num := range nums {
        if i == 0 {
            continue
        }
        if idx, ok := lastIndex[num]; !ok {
            // haven't seen the number before therefore it is a distinct value which increases sum of squares
            s1[i] += s1[i-1] + i + 1
            s2[i] += s2[i-1] + 2 * s1[i-1] + i + 1
        } else {
            set := make(map[int]bool)
            l := 0
            for j := i-1; j > idx; j-- {
                set[nums[j]] = true
                l += len(set)
            }
            s1[i] += s1[i-1] + i - idx
            s2[i] += s2[i-1] + 2 * l + i - idx
            // fmt.Println(l, idx, set)
        }
        s1[i] %= mod
        s2[i] %= mod
        lastIndex[num] = i
        // fmt.Println(s1, s2, lastIndex)
		// focusing just on s1 when you have a new number you add 1 to all combinations
		// so [1] : 1, [2] : 1, [1, 2] : 1 + 1 + 2 = 4
		// becomes [1]: 1, [2, 1]: 2 and [1, 2, 1]: 2, 
		// so s1 becomes [1, 3, 5]..\ 
    }
    sum := 0
    for _, squares := range s2 {
        sum += squares
    }
    return sum % mod
}