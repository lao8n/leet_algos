package data_structures

import (
	"sort"
	"strconv"
)

// approaches
// * sort numbers by first digit, then second digit etc.

// specifics
// * complication is if have 321 3210 then we want to go 3213210 i.e. 0 is later, if we have 321 3214 then want to go 3214321
// * need to sort from largest to smallest
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		numStringI, numStringJ := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])

		digitI, remI := getDigit(numStringI)
		digitJ, remJ := getDigit(numStringJ)

		// choose number with higher digit
		for i := 0; i < len(numStringI)+len(numStringJ) && digitI == digitJ; i++ {
			if len(remI) == 0 {
				remI = numStringI
			}
			if len(remJ) == 0 {
				remJ = numStringJ
			}
			digitI, remI = getDigit(remI)
			digitJ, remJ = getDigit(remJ)
		}
		// now digits are different
		if digitI < digitJ {
			return false
		} else {
			return true
		}
	})
	result := ""
	leading0 := true
	for i, num := range nums {
		if num == 0 && leading0 && i != len(nums)-1 {
			continue
		} else {
			leading0 = false
			result += strconv.Itoa(num)
		}
	}
	return result
}

func getDigit(num string) (int, string) {
	d, _ := strconv.Atoi(string(num[0]))
	num = num[1:]
	return d, num
}
