package data_structures
// clarifying questions
// * beautiful? substring has even length & only 0s and 1s
// * minimum number of changes? to make beautiful -> can cahnge s[1] to 1 and s[0

// approaches
// * dynamic programming - feels like word break question -> is there any partition where all even and only 1s and 0s 
// * so then have true or false for each index whether to partition into 1s and 0s but if have a bunch of falses what should you do then?
// * probably want dp to be num of changes required? correct...
func minChanges(s string) int {
    dp := make([]int, len(s))
    // base case
    if s[0] != s[1] { // default is set at 0 if equal
        dp[0] = 1
        dp[1] = 1
    }
    // try bottom up data tabulation
    for i := 2; i < len(s) - 1; i = i + 2 {
        // consider next pair
        // if the same -> then no problem
        if s[i] == s[i + 1] {
            dp[i] = dp[i - 1] // just have a new partition no new changes necessary
            dp[i+1] = dp[i - 1] 
        } else if s[i] == s[i - 1] || s[i] != s[i-1] { // overlap are the same i.e. don't change s[i] - change s[i+1]
            dp[i] = dp[i-1]
            dp[i+1] = dp[i-1] + 1 // need to change dp[i-1]
        } 
    } 
    return dp[len(s) - 1]
}
