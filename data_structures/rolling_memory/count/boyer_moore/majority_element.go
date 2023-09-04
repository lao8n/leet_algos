package data_structures

func majorityElement(nums []int) []int {
	// use boyer-moore voting algorithm where there can be at most two
	// majority elements which are more than n/3 times
	// 1st pass find top 2 candidates
	var cand1, cand2, count1, count2 int
	for _, v := range nums {
		if v == cand1 {
			count1++
		} else if v == cand2 {
			count2++
		} else if count1 == 0 {
			cand1 = v
			count1++
		} else if count2 == 0 {
			cand2 = v
			count2++
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	// 2nd pass actually check if candidates > n/3
	for _, v := range nums {
		if v == cand1 {
			count1++
		} else if v == cand2 {
			count2++
		}
	}
	var result []int
	if count1 > len(nums)/3 {
		result = append(result, cand1)
	}
	if count2 > len(nums)/3 {
		result = append(result, cand2)
	}
	return result
}
