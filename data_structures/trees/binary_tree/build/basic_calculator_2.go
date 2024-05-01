package data_structures

import (
	"regexp"
	"strconv"
	"strings"
)

// clarifying questions
// * need to have ordering of operations with multiplication and division before addition and subtraction
// * also have blank spaces to ignore
// * can there be brackets? no

// approaches
// * original idea is use a stack but that only works if reverse polish
// * can we go through in a single pass? probably not maybe two passes 1. multiplication and division 2. addition and subtraction - but then how update? linked list? - easier just to output new string
// * build a tree recursively where the string gives you pre-order traversal
// specifics
// * use byte slice rather than immutable string
// * what if 3 * 4 * 5 -> want to process it as 12 * 5 -> 60
func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	sSplit := splitWithDelimeters(s, "[+\\-*/ ]")

	v, _ := strconv.Atoi(string(sSplit[0])) // guaranteed to be int
	start := &node{Val: v}
	prev := start
	for i := 1; i < len(sSplit); i++ {
		switch sSplit[i] {
		case "+", "-", "*", "/":
			v, _ = strconv.Atoi(string(sSplit[i+1]))
			prev = createNode(prev, string(sSplit[i]), v)
		default:
			// skip
		}
	}
	// recurse from start
	total := recurse(prev)
	return total
}

// removes space and creates an array with each element separate - needed to group numbers like 42
func splitWithDelimeters(s string, pattern string) []string {
	// first remove spaces
	s = strings.ReplaceAll(s, " ", "")

	re := regexp.MustCompile(pattern)
	indexes := re.FindAllStringIndex(s, -1) // -1 to return all matches and not cap at n

	var result []string
	prevIndex := 0

	for _, idx := range indexes { // idx return start and end index
		result = append(result, s[prevIndex:idx[0]])
		result = append(result, s[idx[0]:idx[1]])
		prevIndex = idx[1]
	}

	result = append(result, s[prevIndex:])

	return result
}

func recurse(n *node) int {
	if n == nil {
		return 0
	}
	switch n.Operator {
	case "*":
		return recurse(n.Left) * recurse(n.Right)
	case "/":
		return recurse(n.Left) / recurse(n.Right)
	case "+":
		return recurse(n.Left) + recurse(n.Right)
	case "-":
		return recurse(n.Left) - recurse(n.Right)
	default:
		// should be a num
		return n.Val
	}
}

// what tree operations would we want for 3 + 4 * 2
//   - we want to be able to create a node with a number e.g. 3
//   - we want to be able to add an operation with the children being 1. calculation up to this point and everything after - but below is wrong - need to be able to swap
//     *
//   - 2
//     3   4
//   - if have + or - operation create new parent
func createNode(prevNode *node, op string, val int) *node {
	newOpNode := &node{Operator: op}
	newValNode := &node{Val: val}
	switch op {
	// add as parent
	//      +
	// prev   new
	case "+", "-":
		newOpNode.Left = prevNode
		newOpNode.Right = newValNode
		return newOpNode
	// add as child:
	//        prev.op
	// prev.left      *
	//        prev.right new
	// or if just prev val
	case "*", "/":
		if prevNode.Operator == "" || prevNode.Operator == "*" || prevNode.Operator == "/" { // just a value
			newOpNode.Left = prevNode
			newOpNode.Right = newValNode
			return newOpNode
		}
		newOpNode.Left = prevNode.Right
		newOpNode.Right = newValNode
		prevNode.Right = newOpNode
		return prevNode
	}
	return nil
}

type node struct {
	Operator string
	Val      int
	Left     *node
	Right    *node
}
