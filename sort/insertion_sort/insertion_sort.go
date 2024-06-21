package data_structures

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		elementToSort := arr[i]
		for j := i - 1; j >= 0 && arr[j] > elementToSort; j-- {
			arr[j+1] = arr[j] // shift sorted elements along
		}
		arr[j+1] = elementToSort
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Unsorted array:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array:", arr)
}
