package data_structures

// Alice or Bob wants to give the other an odd number such that the game ends with the other having n == 1
// If Alice starts with an even number Alice will only subtract 1 to give Bob an odd number.
// Bob cannot give Alice an odd number if Bob starts with an odd number. Since an odd number only has odd factors, when an odd number subtracts an odd number (Any of the number's factor), it becomes even.
// It works the same for Alice too. If Alice is given an odd number, Alice can only give Bob an even number. Bob will subtract 1 to continue to give Alice an odd number.

func divisorGame(n int) bool {
	return n%2 == 0
}
