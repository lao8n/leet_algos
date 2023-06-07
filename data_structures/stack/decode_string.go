package stack

func decodeString(str string) string {
	num := 0
	stack := []seq{{1, []rune{}}}

	for _, s := range str {
		switch {
		case s >= '0' && s <= '9':
			num = num*10 + int(s-'0')
		case s == '[':
			stack = append(stack, seq{num, []rune{}})
			num = 0
		case s == ']':
			poppedIndex := len(stack) - 1
			popped := stack[poppedIndex]
			stack = stack[:poppedIndex]
			topIndex := len(stack) - 1
			for j := 0; j < popped.num; j++ {
				stack[topIndex].chars = append(stack[topIndex].chars, popped.chars...)
			}
		default:
			topIndex := len(stack) - 1
			stack[topIndex].chars = append(stack[topIndex].chars, s)
		}
	}
	return string(stack[0].chars)
}

type seq struct {
	num   int
	chars []rune
}
