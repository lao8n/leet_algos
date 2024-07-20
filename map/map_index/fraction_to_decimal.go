// clarifying questions
// * to repeat how many numbers need to repeat? just one? two? depends upon the carry?

// approach
// * do division but stop after length of answer string is > 10^4

// specifics
// * use token buffer rather than immutable strings
// * how to do cycle detection?
// * strings to ascii - but do it digit by digit
// * ways to detect cycle -> rotational argument? have O(n^2) because need to check start index and end index
// * maybe use map of different cycle lengths -> if any two numbers repeat is that sufficient?
// * cycle is more important d and rem ->
func fractionToDecimal(numerator int, denominator int) string {
	// greater than zero
	neg := false
	if numerator < 0 {
		neg = !neg
		numerator *= -1
	}
	if denominator < 0 {
		neg = !neg
		denominator *= -1
	}

	d := numerator / denominator   // 1 / 2 = 0
	rem := numerator % denominator // 1 % 2 = 1
	s := strconv.Itoa(d)
	answer := []byte(s)
	if neg && !(d == 0 && rem == 0) {
		answer = append([]byte{'-'}, answer...)
	}
	if rem > 0 {
		answer = append(answer, '.')
	}

	// cycle detection - cant just check first number, need to check any number afterwards
	// e.g. could be 0.3534012012012 etc. - need to check every seq from index i to end ->
	cycle := make(map[[2]int]int) // d and rem -> index where cycle started
	i := len(answer)
	for rem > 0 {
		numerator = rem * 10                 // 10
		d = numerator / denominator          // 10 / 2 = 5
		rem = numerator % denominator        // 10 % 2 = 0
		answer = append(answer, byte('0'+d)) // not sure if correct
		repeat := [2]int{d, rem}
		if j, ok := cycle[repeat]; ok {
			// copy
			repeating := make([]byte, i-j)
			copy(repeating, answer[j:i])
			answer = answer[:j] // remove whole repeating section
			// fmt.Println(string(answer), string(repeating))
			answer = append(answer, '(')
			answer = append(answer, repeating...)
			answer = append(answer, ')')
			break
		}
		cycle[repeat] = i
		i++
		// assume two ways cycles can occur
	}
	return string(answer)
}