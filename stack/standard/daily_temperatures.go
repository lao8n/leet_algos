package data_structures

// Given an array of integers temperatures represents the daily temperatures, return an array answer
// such that answer[i] is the number of days you have to wait after the ith day to get a warmer
// temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.

// have a stack either of increasing or decreasing temps
// time complexity O(n) (each element can only be popped once)
// space complexity O(n)
func dailyTemperatures(temperatures []int) []int {
	solution := make([]int, len(temperatures))

	stack := make([]info, 0)

	// recurrent case
	for i, t := range temperatures {
		// stack is highest -> lowest temps
		// if new temp is higher than stack - then remove from stack until no more higher
		for len(stack) > 0 && t > stack[len(stack)-1].temp {
			popped := stack[len(stack)-1]
			solution[popped.day] = i - popped.day
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, info{temp: t, day: i})
	}
	return solution
}

type info struct {
	temp int
	day  int
}
