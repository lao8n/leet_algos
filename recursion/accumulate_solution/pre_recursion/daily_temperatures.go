package recursion

// have a stack either of increasing or decreasing temps
func dailyTemperatures(temperatures []int) []int {
	solution := make([]int, len(temperatures))
	stack := make([]info, 0)

	// recurrent case
	tempRecursion(temperatures, 0, stack, solution)
	return solution
}

func tempRecursion(temperatures []int, i int, stack []info, solution []int) {
	// base case
	if i == len(temperatures) {
		return
	}
	// recursive case
	stack = recursiveStack(i, temperatures[i], stack, solution)
	tempRecursion(temperatures, i+1, stack, solution)
}

// replace for loop with recursion
func recursiveStack(i, t int, stack []info, solution []int) []info {
	// recursive
	if len(stack) > 0 && t > stack[len(stack)-1].temp {
		popped := stack[len(stack)-1]
		solution[popped.day] = i - popped.day
		stack = recursiveStack(i, t, stack[:len(stack)-1], solution)
	}
	return append(stack, info{temp: t, day: i})

}

type info struct {
	temp int
	day  int
}
