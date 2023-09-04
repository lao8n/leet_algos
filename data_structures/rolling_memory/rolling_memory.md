**Rolling Memory**

Definition = 

Time complexity = 

Space complexity = `O(1)` for rolling sum, max etc. Rolling arrays need `O(n)`

**Types**

Rolling sum = loop/iterate through values adding to sum
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

Rolling arrays = build array of rolling values
```
// product_except_self.go
result := make([]int, len(nums))
result[0] = 1
for i := 1; i < len(nums); i++ {
    result[i] = result[i - 1] * nums[i - 1]
}
```
For rolling product in `product_except_self.go` the key idea is that you can optimise from a left and right rolling product array to just left and calculating right on the fly for `O(1)`. For matrix multiply in `num_tilings_optim.go` rather than multiplying together `n` times can instead recursively multiply `n/2` matrices together

Rolling pointer = 
```
```