package data_structures

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no two elements were swapped in the inner loop, then the array is sorted
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)
	bubbleSort(arr)
	fmt.Println("Sorted array:", arr)
}
