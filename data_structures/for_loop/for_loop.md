**For Loop**

Definition = way to iterate over the elements of a data structure

Time complexity = `O(n)`. For bubble sort and insertion sort `O(n^2)`, but slightly better for insertion as it is an online algorithm (doesn't need full list), and makes less swaps on average

Space complexity = O(0) in place

**Types**

Nested for loops = for `n` iterations move sort one number
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
For insertion sort you have `sorted | unsorted` where you pick the first element in unsorted and then find where it fits in the sorted section, and you have to shift all the larger elements along