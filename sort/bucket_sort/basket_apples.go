package data_structures

func maxNumberOfApples(weight []int) int {
	bucketSort := make([]int, 10001) // max weight is 10^3
	for _, w := range weight {
		bucketSort[w]++
	}

	sum, count := 0, 0
out:
	for w, c := range bucketSort {
		for c > 0 {
			sum += w
			if sum <= 5000 {
				count++
			} else {
				break out // could also just return
			}
			c--
		}
	}
	return count
}
