**Map**
Definition = map is a data structure with keys and values for O(1) insertion and lookup. Most common example uses hashing function to map from unique key to an index in an array.
Time complexity = O(n) creation, although O(1) lookup
Space complexity = O(n)

**Types**
Map of indices = replace two for loops with one and a map lookup. 
```
two_sum.go
nums_map := make(map[int]int)

```

