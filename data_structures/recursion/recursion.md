**Recursion**

Definition = When a function calls itself. Useful for divide and conquer and dynamic programming questions

Time complexity = Typically `O(n logn)` if recurse all space

Space complexity = Usually `O(n)` 

**Types**

**Divide and conquer**

**Divide and conquer accumulate**

**Dynamic programming: Summation**
```
// fibonacci/main.ts
if(cache[i] != undefined){
    return cache[i]!;
} 
cache[i] = fibRecMemoize(n - 1n, cache) + fibRecMemoize(n - 2n, cache)
``` 
Key thing is try to access memoized value, and only if that fails calculate recursively. This can be extended as in `tribonacci.go` where sum result of `tribonacciRecursive` with `n-1`, `n-2`, `n-3`

**Dynamic programming: Boolean**
```
// is_interleave.go
key := s1 + " " + s2
...
chooseLeft, chooseRight := false, false
if s1[0] == s3[0] {
    chooseLeft = isInterleaveRecursive(memoized, s1[1:], s2, s3[1:])
}
if s2[0] == s3[0] {
    chooseRight = isInterleaveRecursive(memoized, s1, s2[1:], s3[1:])
}
memoized[key] = chooseLeft || chooseRight
```
Key part is recursive section where try both paths and see if one of them is true.
**Dynamic programming: Max, Min and Longest**
```
// max_profit.go
key := fmt.Sprintf("%d %d %d", d, k, o)
if profit, ok := data.memoized[key]; ok {
    return profit
}
...
notSellNotBuy = data.maxProfitRecursive(d + 1, k, o)
notSellBuy, sellNotBuy := 0, 0
if o == false {
    notSellBuy = data.maxProfitRecursive(d + 1, k, true) - data.prices[d]
}
if o == true && k > 0 {
    sellNotBuy = data.prices[d] + data.maxProfitRecursive(d + 1, k - 1, false)
}
```
Key thing is 1. define state by a string with days, number of transactions and bought price (this is better than storing profit because it leads to more options) 2. calculate each possible option such as not sell and not buy and then pick the most profitable one. Variations like `max_profit_cooldown.go` involve adding cooldown as a state variable or `max_profit_fee.go` adding a fee to any sell decision to make it less lucrative. Other max variations include `maximum_score.go` where the state is defined as a `[][]int` or `delete_and_earn.go` where it is `map[int]int`
```
// paint_house.go
if unavailableColour >= 0 && d.memoization[houseNum][unavailableColour] != 0 {
    return d.memoization[houseNum][unavailableNum]
}

redCost, blueCost, greenCost := math.MaxInt, math.MaxInt, math.MaxInt
if red != unavailableColour {
    redCost = d.costs[houseNum][red] + d.minCostsRecursion(houseNum+1, red)
}
if blue != unavailableColour {
    blueCost = d.costs[houseNum][blue] + d.minCostsRecursion(houseNum+1, blue)
}
if green != unavailableColour {
    greenCost = d.costs[houseNum][green] + d.minCostsRecursion(houseNum+1, green)
}
```
Minimization questions are no different where again calculate value for each option and then pick the lowest one. For `min_difficulty.go` the recursive state is also a `[][]int` but the options require a for loop to iterate through, whilst `min_cost_tickets.go` is similar to `max_profit.go` comparing between hold and buy decisions. `min_cost_climbing_stairs.go` is different in that it returns the two last step (effectively memoizing only those).
```
// longest_common_subsequence_optimised.go
if text1[0] == text2[0] {
    longest = longestCommonSubsequenceRecursive(text1[1:, text2[1:], memoized]) + 1
} else {
    longest = int(math.Max(
        float64(longestCommonSubsequenceRecursive(text1[1:], text[2:], memoized)),
        float64(longestCommonSubsequenceRecursive(text1, text2[1:], memoized)),
    ))
}
```
These longest questions involving building longest, but do not accumulate the sequence instead embedding it in the memoized `map[string]map[string]int` assigning `memoized[text1][text2] = longest`. Similarly for `longest_increasing_path.go` we have `memoizedPaths [][]int` whereas `longest_increasing_subsequence.go` we just use `memoized []int`

**Dynamic programming: Number of Ways**
