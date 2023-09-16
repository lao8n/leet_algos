**Matrices**

Definition = 1D, 2D, or 3D data structures useful for storing discrete data

Time complexity = `O(n^d)` where `d` is the number of dimensions. For dynamic programming questions the time complexity is `O(n r)` where `r` is the recurrence step.

Space complexity = `O(n^d)` where `d` is the number of dimensions

**Types**

**Dynamic programming: Boolean**
```
// word_break.go
memoized := make([]bool, len(s) + 1)
memoized[0] = true
for lenI := 1; lenI <= len(s); lenI++ {
    for partI := 0; partI < lenI; partI++ {
        if memoized[partI] && wordMap[s[partI:lenI]] {
            memoized[lenI] = true
            break // to next loop
        }
    }
}
```
Memoized is just a slice, as we try every possible length and every part.

**Dynamic programming: Max, min & longest**
```
// max_profit.go
dp := make([][]int, k + 1)
for i := 0; i < k + 1; i++ {
    dp[i] = make([]int, len(prices))
}
for i := 1; i < k + 1; i++ {
    maxDiff := -prices[0]
    for j := 1; j < len(prices); j++ {
        maxDiff = max(maxDiff, dp[i-1][j-1] - prices[j-1])
        dp[i][j] = max(dp[i][j-1], maxDiff + prices[j])
    }
}
```
Just like the recursive case we need to choose between buying and selling, although arguably this is less transparent then the recursion code. If only what happened in the previous round matters, however, sometimes the code can be significantly simplified as with `max_profit_cooldown.go` where `sold, held, reset = held + p, max(held, reset - p), max(reset, sold)`. 
```
// paint_houses.go
prevHouse := make([]int, 3) // 3 colours
for i := 0; i < len(costs); i++ {
    currHouse := make([]int, 3)
    // paint red
    currHouse[0] = costs[i][0] + min(prevHouse[1], prevHouse[2])
    ...
    prevHouse = currHouse
}
```
With cost minimization, the same applies where the recursions can get complicated unless you just focus on the previous round. Variations such as `paint_house_k_colours.go` extend the same idea to `currHouse := make([]int, k)`. Typically the hardest step is coming up with the recurrence relation and what will uniquely define the memoization matrix such as with `min_difficulty.go` where the recurrence relation is itself a separate function. Again this questions can be optimised by just considering previous rather than keeping all rounds. The Bellman-Ford algorithm used in `find_cheapest_price.go` is essentially just cost minimization for dynamic programming, the slight variation coming that it is applied to graphs.
```
// longest_common_subsequence.go
if text1[i] == text2[j] {
    tabulation[i][j] = tabulation[i + 1][j + 1] + 1
} else {
    tabulation[i][j] = int(math.Max(float64, tabulation[i + 1][j]), float64(tabulation[i][j + 1]))
}
```
Longest questions are similar where the key step is again defining the recurrence relation.

**Dynamic programming: Number of ways**
```
// change_coin.go
tabulation[0] = 1
for _, c := range coins {
    for i := c; i < amount + 1; i++ {
        tabulation[i] += tabulation[i - c]
    }
}
```
In the recursive case for `change_coin.go` we looped over `coins` adding to `numWays`, here we directly increment `tabulation[i]`, although the result is the same. Unlike recursion cases we do not need to worry about memoization and unnecessary repeated calculations as we do the minimum amount by virtue of going through the tabulation. Other variations include text based like `count_vowel_permutations.go`, or `num_tilings.go` where we optimise to just `oneFilledWays` and `twoFilledWays` slices, but the general principal is the same.

**Optimisation: Backpropagation**
```
type graph struct {
    graph map[int]*node
    oper_table map[int]float64
    grad_table map[int]float64
}
```
Backpropagation involves maintaining matrices (although here they are maps) of forward pass outputs and backward pass gradients that are calculated once and then re-used later.
