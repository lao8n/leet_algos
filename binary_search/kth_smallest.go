package data_structures

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	i, j := matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]
	for i < j {
		mid := i + (j-i)/2 // (i + j) / 2 doesn't work
		if smallerCount(matrix, mid) < k {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return i
}

func smallerCount(matrix [][]int, k int) int {
	ret := 0
	for _, v := range matrix {
		ret = ret + sort.Search(len(v), func(i int) bool {
			return v[i] > k
		})
	}
	return ret
}
