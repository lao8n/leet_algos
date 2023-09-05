**Rolling Memory**

Definition = Whilst looping through input data update some memory with rolling sums, products, counts and maxes

Time complexity = Typically `O(n)` as still need to iterate through loop

Space complexity = `O(1)` for rolling sum, max etc. Rolling arrays need `O(n)`

**Types**

**Rolling sum** = loop/iterate through values adding to sum
```
// add_two_numbers.go
for l1 != nil || l2 != nil {
    if l1 != nil {
        sum += l1.Val
        l1 = l1.Next
    }
    ...
}
```
Particularly useful for questions where need to multiply by 10 and handle remainder like `reverse_integer.go`.

**Rolling arrays** = build array of rolling values
```
// product_except_self.go
result := make([]int, len(nums))
result[0] = 1
for i := 1; i < len(nums); i++ {
    result[i] = result[i - 1] * nums[i - 1]
}
```
For rolling product in `product_except_self.go` the key idea is that you can optimise from a left and right rolling product array to just left and calculating right on the fly for `O(1)`. For matrix multiply in `num_tilings_optim.go` rather than multiplying together `n` times can instead recursively multiply `n/2` matrices together

**Rolling max** = store rolling max
```
// max_sub_array.go
for i := 1; i < len(nums); i++ {
    if currentRollingSum < 0 {
        currentRollingSum = 0
    }
    currentRollingSum += nums[i]
    if currentRollingSum > maxRollingSum {
        maxRollingSum = currentRollingSum
    }
}
```
Kadane's algorithm in `max_sub_array.go` involves storing `currentRollingSum` and `maxRollingSum` and resetting `currentRollingSum` if it falls below 0. Variations include comparing to `marginal_profit` in `maxProfit.go` and for `maxSubarraySumCircular.go` having to maintain a `minSum`, `maxSum` and `rollingSum` to compare whether middle subarray or outer subarray is bigger.

**Rolling counts** = store rolling counts
```
// majority_element.go
var cand1, cand2, count1, count2 int
for _, v := range nums {
    if v == cand1 { count1++} 
    else if v == cand2 { count2++ }
    else if count1 == 0 { cand1 = v; count1++ }
    else if count2 == 0 { cand2 = v; count2++ }
    else { count1--; count2-- }
}
```
In the Boyer-Moore voting algorithm `majority_element.go` the key trick is there can only be two majority elements with more than `n/3` votes so 1st pass is to find two top 2 candidates and second pass to check if greater than `n/3`.s