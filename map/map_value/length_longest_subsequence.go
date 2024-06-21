package data_structures
// clarifying questions
// * not ordered? not ordered
// * guaranteed? no
// * always positive numbers? yes
// * can skip numbers
// approaches
// * to get a single subsequence up to target could do a rolling sum with two pointers where increment right pointer until >= and then decrement left until <=
// * rolling map where lookup all possible subsequences
func lengthOfLongestSubsequence(nums []int, target int) int {
    m := make(map[int]int) // number to length
    // for each element we can either add it to the subsequence or ignore it
    for _, n := range nums {
        newM := make(map[int]int)
        for k, v := range m {
            // longest way to get to k + n
            newWay := v + 1 
            if newWay > m[k + n] {
                newM[k+n] = newWay
            } 
        }
        if 1 > m[n] {
            newM[n] = 1
        }
        for k, v := range newM {
            if k <= target {
                m[k] = v
            }
        }
    }
    if m[target] == 0 {
        return -1
    }
    return m[target]
}