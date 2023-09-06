**Two pointers**

Definition = way to track two points in a data structure

Time complexity = `O(n)` for managing the two pointers

Space complexity = O(0) in place

**Types**

**Standard two pointers** = compare two different points
```
// trap_water.go
for left <= right {
    if height[left] <= height[right] {
        if height[left] >= maxLeft { ... }
        left++
    } else {
        if height[right] >= maxRight { ... }
        right--
    }
}
```

**Two pointers for sorting** = for `n` iterations move sort one number
```
// bubble_sort.go
for i := 0; i < n; i++ {
    for j :=0; j < n - i - 1; j++ { // only need to pick from unsorted section
        if arr[j] > arr[j + 1] {
            arr[j], arr[j + 1] = arr[j + 1], arr[j]
        }
    }
}
```
For bubble sort you have `unsorted | biggest` where you find the largest number in the unsorted section and put at front of biggest. https://stackoverflow.com/questions/17270628/insertion-sort-vs-bubble-sort-algorithms
```
// insertion_sort.go
for i := 1; i < len(arr); i++ {
    elementToSort := arr[i]
    for j := i - 1; j >= 0 && arr[j] > elementToSort; j-- {
        arr[j + 1] = arr[j] // shift sorted elements along
    }
    arr[j+1] = elementToSort
}
```
For insertion sort you have `sorted | unsorted` where you pick the first element in unsorted and then find where it fits in the sorted section, and you have to shift all the larger elements along. The time complexities for bubble sort and insertion sort are `O(n^2)`, but slightly better for insertion as it is an online algorithm (doesn't need full list), and makes less swaps on average

There are also a number of reordering questions like `sort_colors.go` where we perform swaps on specific data. Or `rotate.go` where sort elements by divisor set.

**Two pointers on sorted values** = sort the elements so can find a target sum
```
// three_sum.go
sort.Ints(nums)
...
for left < right {
    sum := nums[left] + nums[right]
    if sum == target { ...; left++; right++ }
    if sum > target { right-- }
    if sum < target { left++ }
}
```
You might also need to skip duplicates or test if closest as in `three_sum_closest.go`. Can also be used where a pointer in two different arrays as with `min_meeting_rooms.go`

**Two pointers with midpoint** = do left and right indices off of new midpoint 
```
// longest_palindrome.go
for int(midpoint) < len(s) {
    left, right = int(math.Floor(midpoint)), int(math.Ceil(midpoint))
    ...
    midpoint += 0.5
}
```
Manacher's algorithm takes this further memoizing some of the calculations with mirrored values as in `longest_palindrome_manacher.go`. It can also be used to partition and search on sorted elements as with `rotated_search.go`

**Two pointers with map** = Use a map to check values seen before, jump to index and store counts
```
// longest_non_repeating_substring.go
for substringStart, substringFinish ... {
    if j, ok := m[s[substringFinish]]; ok && j >= substringStart {
        substringStart = j + 1
    }
    m[s[substringFinish]] = substringFinish
    ...
}
```
In `longest_non_repeating_substring.go` removed `O(n)` complexity checking for existence in substring using map and another `O(n)` by jumping straight to the next index. In `min_window.go` build a map of counts of letters in one string, then increment right pointer until have one of each, then see if i can increment left pointer.