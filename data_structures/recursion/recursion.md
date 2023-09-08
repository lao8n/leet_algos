**Recursion**

Definition = When a function calls itself. Useful for divide and conquer and dynamic programming questions

Time complexity = Typically `O(n)` if recurse all space, and `O(log n)` to go depth

Space complexity = Usually `O(n)` 

**Types**

**Divide and conquer: Graph**
```
// max_depth.go
if root == nil { return 0 }
leftDepth := maxDepth(root.Left)
rightDepth := maxDepth(root.Right)
if leftDepth > rightDepth { return leftDepth + 1}
return rightDepth + 1
```
Graph traversal, usually DFS, involves a base case and then a left and right branch recursion on the forward pass and returning a value on the backward pass. That might be a boolean as in `is_same_tree.go`, a specific `TreeNode` as with `lowest_common_ancestor.go` or nothing at all as with `connect_nodes.go` where instead building a graph. For non-tree graphs we need to track whether we have visited a node as with `clone_graph.go` where we `if n, seen := track[node.Val]; seen { return n }`

**Divide and conquer: Search**
```
// top_k_frequent.go
leftOfPivot := partition(freqSlice, left, right)
if leftOfPivot == k { return freqSlice[:k] }
if leftOfPivot > k { return quickSelect(freqSlice, k, left, leftOfPivot - 1) }
return quickSelect(freqSlice, k, leftOfPivot + 1, right)
```
Quick select involves partitioning around a randomly chosen value, and using quick sort only for the partitioned part that the value might be in. Need to be careful, as in `find_k_largest.go` in partition function to 1. pivot to the end, 2. swap all elements to the left 3. pivot back to final place. Also, as with `find_k_largest_indirect.go` can find it faster to find the `n - k` smallest and then indirectly find the `k` largest. If the data is already sorted then can simple partition the data without needing to swap as in `rotated_search.go`

**Divide and conquer: Sort**
```
// quicksort.go
pivotIndex := partition(s, start, end, rand.Intn(end - start + 1) + start)
sortRecursiveWithSwap(s, start, pivotIndex + 1)
sortRecursiveWithSwap(s, pivotIndex + 1, end)
```
There are main approaches to sorting. Quick sort is conquer and divide because most of the work happens in the forward-pass in the partitioning function looping through values and swapping if less than or greater than a pivot. Merge sort in contrast is divide and conquer because the forward-pass partition is simply dividing the array in half, whereas the backward pass is where the sorting actually happens. `treesort.go` is similar to quick sort in the sense that the sorting occurs on the forward pass into the tree sorting where it should be placed based upon comparison

**Divide and conquer: Sum**
```
// find_target_sum_ways.go
positiveWays := findTargetSumWays(nums[1:], target - nums[0])
negativeWays := findTargetSumWays(nums[1:], target + nums[0])
return positiveWays + negativeWays
```
This is similar to the dynamic programming `num_ways` questions but is simple because all the numbers need to be included once, the only choice is whether to add or subtract and because we don't need any memoization. Similarly, `longest_increasing_path.go` seems like a dynamic programming question but because we do not memoize it isn't dynamic programming.

**Divide and conquer: Combinations**
```
// letter_combinations.go
lastElement := len(digits) - 1
combinations := letterCombinations(digits[:lastElement])
var newCombinations []string
for _, l := range digitLetterMap[digits[lastElement]] {
    for _, c := range combinations {
        newCombinations = append(newCombinations, c + l)
    }
}
```
We use recursion but only along one dimension, combining an element with every combination or path already created. `permute.go` is similar where recurse with `solutionsWithoutN := permute(withoutN)` and then create `solutionsWithN` by adding `n` to each of the `solutionsWithoutN`. Other approaches rather than adding an element to every possible solution already created, instead simply append new solutions which are created separately, such as with `flatten_nested_iterator.go`. There can always be more complicated creation of new combinations such as `merge_intervals.go`

**Divide and conquer: Neighbours**
```
// num_islands.go
for _, n := range getLandNeighbours(l, grid) {
    dfs(n, grid)
}
```
These are questions where a list of possible neighbours, usually constrained to a grid are then recursed on. If all paths need to be created then it would be necessary to backtrack, but when only looking for one as in `valid_path.go`, `topological_sort.go` or `word_search.go` it is fine to have a `visited` matrix which is updated once.

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
```
// change_coin.go
key := fmt.Sprintf("%d %d", len(coins), amount)
if memAmount, ok := s.memoized[key]; ok {
    return memAmount
}
numWays := 0
for i, c := range coins {
    numWays += s.changeRecursive(amount - c, coins[i:]) 
}
s.memoized[key] = numWays
```
Rather than choosing between options as in max/min questions instead we sum where key trick to prevent duplicates is to use `coins[i:]`. `coin_change.go` expands the number of ways question into a minimization problem to choose the `leastNumCoins`. `num_decodings.go` sums across one and digit cases and `num_ways.go` adds `totalNumWays := diffColourNumWays + sameColourNumWays`. In multiple cases we memoize on a string key using `fmt.Sprintf` for example `key := fmt.Sprintf("%d %d", nRemain, nSameColour)`