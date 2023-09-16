**Map**

Definition = a map is a data structure with keys and values for O(1) insertion and lookup. Most common example uses hashing function to map from unique key to an index in an array.

Time complexity = O(n) creation, although O(1) lookup

Space complexity = O(n)

**Types**

**Map of indices** = need to return indices, but save one for loop with map lookup instead
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

**Map of values** = useful for string to int mapping such as with roman_to_int questions
```
// roman_to_int.go
int_to_roman_map := map[rune]int{
		'M': 1000,
		'D': 500,
        ...
}
```
Also useful if same value is going to be accessed multiple times as with `sparse_matrix_multiply.go`

**Map of counts** = count number of each element, often a useful way to track anagrams
```
// is_anagram.go
anagramMap := make(map[rune]int, len(s))
for _, sc := range s {
    anagramMap[sc]++
}
```
Note cannot use slice as key in map, but can use array as in `groupAnagrams.go` where `m := make(map[[26]int][]string)`. When combined with cumulative sum can always be used to find the number of sums equalling a value as in `sub_array_sums.go`

**Map as set** = golang doesn't have a set so just use map instead
```
// num_jewels_in_stones.go
jewelsSet := make(map[byte]bool)
for i := 0; i < len(jewels); i++ {
    jewelsSet[jewels[i]] = true
}
```
Can also use sets for comparison as in `intersection.go` or as a way to track whether seen a value before as in `is_happy.go`