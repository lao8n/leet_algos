**Stack**

Definition = a LIFO data structures that supports push and pop operations

Time complexity = 

Space complexity 

**Types**

**Stack**
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
3 main operations are peek, pop and push which are achieved above in `daily_temperatures.go`, usually a slice is fine.