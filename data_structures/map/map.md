**Map**
Definition = map is a data structure with keys and values for O(1) insertion and lookup. Most common example uses hashing function to map from unique key to an index in an array.
Time complexity = O(n) creation, although O(1) lookup
Space complexity = O(n)

**Types**
Map of indices = need to return indices, but save one for loop with map lookup instead
```
// two_sum.go
nums_map := make(map[int]int)
for i := range nums {
    ...
    if j, ok := nums_map[complement]; ok { ... }
    else { nums_map[nums[i]] = i }  
}
```
For extensions like three-sum and four-sum sorting and two pointers works better to prevent duplicates. Also good for O(1) read e.g. randomized insert with O(1) deletion. 

Map of values = useful for string to int mapping such as with roman_to_int questions
```
// roman_to_int.go
int_to_roman_map := map[rune]int{
		'M': 1000,
		'D': 500,
        ...
}
```
Also useful if same value is going to be accessed multiple times as with sparse_matrix_multiply.go