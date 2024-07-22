package data_structures

// counting sort runs in O(A + B) A is the number of items to be sorted and B is the difference between the smallest and largest element

func diagonalSort(mat [][]int) [][]int {
	// sort all first column
	for i := 0; i < len(mat); i++ {
		sortDiagonal(mat, i, 0)
	}
	// sort first row (except 1st element)
	for i := 1; i < len(mat[0]); i++ {
		sortDiagonal(mat, 0, i)
	}
	return mat
}

// inplace
func sortDiagonal(mat [][]int, x int, y int) {
	// extract nums
	nums := make([]int, 0)
	i, j := x, y
	for i < len(mat) && j < len(mat[0]) {
		nums = append(nums, mat[i][j])
		i++
		j++
	}
	// sort nums
	nums = countingSort(nums)

	// assign nums
	i, j, k := x, y, 0
	for i < len(mat) && j < len(mat[0]) {
		mat[i][j] = nums[k]
		i++
		j++
		k++
	}
}

func countingSort(nums []int) []int {
	// count
	counts := make([]int, 101)
	for _, num := range nums {
		counts[num]++
	}

	// create list
	sortedList := make([]int, len(nums))
	j := 1
	for i := 0; i < len(nums); i++ {
		for counts[j] == 0 {
			j++
		}
		sortedList[i] = j
		counts[j]--
	}
	return sortedList
}
