**Sorted**

Definition = For some problems the key step is to sort the data in order to make the problem easier

Time complexity = `O(n logn)`

Space complexity = Can be done in-place

**Types**

**Sort first** = some problems become much easier if we sort the input first
```
// merge_intervals.go
sort.Slice(intervals,
    func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
```
`merge_intervals.go` becomes much easier if the intervals are sorted by start index because then have a simple choice between: 1. if the interval overlaps update current interval with new end time 2. if they don't add new interval entirely. This is also useful in anagram related questions like `is_anagram.go` and `group_anagram.go` where a string or slice is sorted for comparison. 

**Custom sorts** = implementation of custom sorts
```
// song_sorting.go
type customSort struct { ... }
func (x customSort) Len() int { ... }
func (x customSort) Less(i, j int) bool { ... }
func (x customSort) Swap(i, j int) { ... }

sort.Sort(customSort{tracks, func(x, y *Track) bool {
    if x.Title != y.Title { 
        return x.Title < y.Title 
    }
    ...
}})
```
To implement a truly custom sort as with `song_sorting.go` need to make sure struct implement the sort interface. Most examples can use `sort.Slice` or `sort.SliceStable` however as in `k_closest.go`