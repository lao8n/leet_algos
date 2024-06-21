**Stack**

Definition = a LIFO data structures that supports push and pop operations

Time complexity = `O(n)` for each element popped once

Space complexity = `O(n)` for max size of data structure

**Types**

**Standard Stack**
```
// daily_temperatures.go
stack := make([]info, 0)
for i, t := range temperatures {
    ...
    popped := stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]
    ...
    stack = append(stack, info{temp: t, day: i})
}
```
3 main operations are peek, pop and push which are achieved above in `daily_temperatures.go`, usually a slice is fine. `trap_water.go` uses this to sum horizontal strips (as opposed to vertical strips as with two pointers) storing the values down the valley in a stack and popping them as you go up. Work for any nested structure such as `flatten_nested_iterator.go` and can also be used to implement other data structures like queues in `queue_using_stack.go` or `min_stack.go`. 

**Stack for decoding**
```
// is_valid.go
case '(', '[', '{':
    stack = append(stack, c)
case ')', ']', '}':
    ...
    stack = stack[:len(stack - 1)]
```
Stacks are useful for decoding tasks like in `decode_string.go`, `eval_rpn.go` or `is_valid.go` where you can build order of brackets, operators. 

**Stack for recursion**
```
// in_order_traversal.go
for len(stack) > 0 {
    top := stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]
    if top.Left != nil && !seen[top.Left] {
        stack = append(stack, top)
        stack = append(stack, top.Left)
    }
    ...
}
```
Can extend to have no seen map by dfs adding left to stack, and then itself then set `top = top.Right` as in `in_order_traversal_no_seen.go`