package data_structures

import (
	"fmt"
	"sort"
)

// specifics
// * need to be careful anytime we append
func largestTimeFromDigits(arr []int) string {
	// permutations
	solutions := permute(arr, []int{})
	// sort
	sort.Slice(solutions, func(x, y int) bool {
		for i := 0; i < len(arr); i++ {
			if solutions[x][i] > solutions[y][i] {
				return true
			}
			if solutions[x][i] < solutions[y][i] {
				return false
			}
		}
		return false
	})
	for _, p := range solutions {
		if isValid(p) {
			return fmt.Sprintf("%v%v:%v%v", p[0], p[1], p[2], p[3])
		}
	}
	return ""
}

func permute(arr []int, acc []int) [][]int {
	// base case
	if len(arr) == 0 {
		return [][]int{acc}
	}
	// recursive case
	solutions := [][]int{}
	for i, num := range arr {
		withoutI := copyAppend(arr[:i], arr[i+1:]...)
		accNum := append(acc, num)
		withoutISolutions := permute(withoutI, accNum)
		solutions = append(solutions, withoutISolutions...)
	}
	return solutions
}

func copyAppend(slice []int, elements ...int) []int {
	result := make([]int, len(slice), len(slice)+len(elements))
	copy(result, slice)
	result = append(result, elements...)
	return result
}

func isValid(arr []int) bool {
	switch arr[0] {
	case 2:
		return arr[1] <= 3 && arr[2] <= 5
	case 1:
		return arr[2] <= 5
	case 0:
		return arr[2] <= 5
	default:
		return false
	}
}
