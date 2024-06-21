package data_structures

// clarifying questions
// * k days sum? T < lower -> - point if T > upper + 1 point
// * diet is to gain weight not lose weight?

// approaches
// * rolling sum - maintain with O(n) time
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	points := 0
	rollingCalories := 0
	count := 0
	for i, cal := range calories {
		count++
		rollingCalories += cal
		// k > 0 days remaining
		if count < k {
			continue
		}

		// rolling k -> remove first element and add curr
		if count > k {
			rollingCalories -= calories[i-k]
		}
		// update points
		if rollingCalories < lower {
			points--
		} else if rollingCalories > upper {
			points++
		}
	}
	return points
}
