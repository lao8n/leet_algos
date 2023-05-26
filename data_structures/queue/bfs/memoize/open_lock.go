package data_structures

// bfs -> initial
// greedily move towards target 4 times at once - if approach results in dead-end do not add to queue
// how keep track of number of moves? maybe need to add it?
func openLock(deadends []string, target string) int {
	// create map O(d) - for deadends but also memoized
	codeMap := make(map[string]bool, len(deadends))
	for _, m := range deadends {
		codeMap[m] = true
	}
	// create queue O(n)
	queue := make([]moves, 0)
	start := "0000"
	if seen := codeMap[start]; !seen {
		queue = append(queue, moves{code: start, moves: 0})
	}

	// loop over queue
	for len(queue) > 0 {
		popped := queue[0]
		if popped.code == target {
			return popped.moves
		}
		queue = queue[1:]
		possMoves := genPossMoves(popped.code, codeMap)
		numMove := popped.moves + 1
		for _, move := range possMoves {
			queue = append(queue, moves{code: move, moves: numMove})
		}
	}

	return -1
}

func genPossMoves(code string, codeMap map[string]bool) []string {
	possMoves := make([]string, 0)
	for i, s := range code {
		c := int(s - '0')
		code1, code2 := []byte(code), []byte(code)
		code1[i], code2[i] = byte((c+1)%10+'0'), byte((c-1+10)%10+'0')
		code1S, code2S := string(code1), string(code2)
		if _, seen := codeMap[code1S]; !seen {
			possMoves = append(possMoves, code1S)
			codeMap[code1S] = true
		}
		if _, seen := codeMap[code2S]; !seen {
			possMoves = append(possMoves, code2S)
			codeMap[code2S] = true
		}
	}
	return possMoves
}

type moves struct {
	code  string
	moves int
}
