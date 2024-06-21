package data_structures

// clarifying questions:
// 1. i go first always? yes
// 2. only remove 1 or 3? no can remove 1 2 or 3

// approach -> dynamic programming problem where have two options each time
// work up from base cases then we work up to n?
// is there a formula? all you can do is to reduce the pile by an odd amount faster? but what is the advantage of that?

// base cases
// 1 = I win
// 2 = I win
// 3 = I win - can remove 3
// 4 = I lose - whatever remove means that they can remove 1 2 or 3
// 5 = I win remove 1 and force them to see 4 and lose
// 6 = I win remove 2 and force them to see 4
// 7 = I win remove 3 and force them to see 4
// 8 = I lose - cannot avoid having them see 5-7 which means they win
// 9 = I win - force them to see 8
// 10 = I win force them to see 8
// 11 = I win force them to see 8
// 12 = I lose cannot avoid seeing 9 - 12
// 13
// 14
// 15
// seems to just be based on whether even or not
// dynamic programming so work up from base cases with an array - only care about hte last 3 so could do that
// cannot just work up from the beginning - can we have a rule to jump
// is it just a multiple of 4?
func canWinNim(n int) bool {
	// base cases
	// if n <= 3 {
	//     return true
	// }
	if n%4 == 0 {
		return false
	}
	return true
	// removeThree := true
	// removeTwo := true
	// removeOne := true
	// curr := false
	// // dynamic cases
	// for i := 4; i <= n; i++ {
	//     // nothing i can do all cases win
	//     if removeOne && removeTwo && removeThree {
	//         curr = false //
	//     // one of them is false
	//     } else if !removeOne || !removeTwo || !removeThree {
	//         curr = true // we can force them to lose
	//     }
	//     // fmt.Println(removeOne, removeTwo, removeThree, i, curr)
	//     // increment by one
	//     removeThree = removeTwo
	//     removeTwo = removeOne
	//     removeOne = curr
	// }
	// return curr
}
