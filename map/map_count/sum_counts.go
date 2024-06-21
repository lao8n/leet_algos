package data_structures
// clarifying questions
// * number of distinct elements? 

// approaches
// * recursive approach with set for length of values
func sumCounts(nums []int) int {
    sum := 0
    for l := 1; l <= len(nums); l++ { // all subarray lengths
        // setup first map
        rollingMap := make(map[int]int)
        for i := 0; i < l; i++ {
            rollingMap[nums[i]] += 1    
        }
        sum += len(rollingMap) * len(rollingMap)
        // do remaining maps
        for i := l; i < len(nums); i++ {
            // remove previous value
            rollingMap[nums[i-l]] -= 1
            if rollingMap[nums[i-l]] == 0 {
                delete(rollingMap, nums[i-l])
            }
            // add new value
            rollingMap[nums[i]] += 1
            sum += len(rollingMap) * len(rollingMap)
        }
    }
    return sum
}