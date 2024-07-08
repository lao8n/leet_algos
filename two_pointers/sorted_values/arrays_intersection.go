package data_structures

// approach
// * 3 pointers

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	i1, i2, i3 := 0, 0, 0
	solution := []int{}
	for i1 < len(arr1) && i2 < len(arr2) && i3 < len(arr3) {
		if arr1[i1] == arr2[i2] && arr2[i2] == arr3[i3] {
			solution = append(solution, arr1[i1])
			i1++
			i2++
			i3++
		} else {
			// increment minimum value
			if arr1[i1] < arr2[i2] {
				if arr1[i1] < arr3[i3] {
					i1++
				} else {
					i3++
				}
			} else { // i2 or i3 smallest
				if arr2[i2] < arr3[i3] {
					i2++
				} else {
					i3++
				}
			}
		}
	}
	return solution
}
