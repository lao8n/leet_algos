package data_structures

// approach
// * update in place, or create new num
func sumOfEncryptedInt(nums []int) int {
	sum := 0
	for _, num := range nums {
		e := numToEncryption(num)
		sum += e
	}
	return sum
}

func numToEncryption(num int) int {
	maxDigit := 0
	j := 0
	for i := num; i > 0; i = i / 10 {
		d := i % 10 // extract number
		if d > maxDigit {
			maxDigit = d
		}
		j++
	}
	encryption := 0
	for ; j > 0; j-- {
		encryption = encryption*10 + maxDigit
	}
	return encryption
}
