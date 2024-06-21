package data_structures

// approaches
// * actually play the game
// * it is a form insertion sort where lower number goes to the back but back numbers aren't sorted
// skills[2,4,1,0,3] -> [6,9,2,4,3]
// * local maximum? maybe there is a clever way but start but just playing the game
func findWinningPlayer(skills []int, k int) int {
	gamesWon := 0
	winner := 0
	queue := make([]int, len(skills)-1)
	for i := 1; i < len(skills); i++ {
		queue[i-1] = i
	}
	count := 0
	for len(queue) > 0 {
		if count >= len(skills) { // already processed all skills
			return winner
		}
		if gamesWon == k {
			return winner
		}
		challenger := queue[0]
		if skills[winner] > skills[challenger] {
			gamesWon++
			queue = append(queue[1:], challenger)
		} else {
			queue = append(queue[1:], winner)
			winner = challenger
			gamesWon = 1
		}
		count++
	}
	return winner
}
