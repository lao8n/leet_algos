package data_structures

// bfs queue question where see if token of rooms is visited = len(rooms)
func canVisitAllRooms(rooms [][]int) bool {
	toDo := make(map[int]bool, len(rooms))
	for i := 0; i < len(rooms); i++ {
		toDo[i] = true
	}

	queue := make([]int, 1)
	queue[0] = 0

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]

		delete(toDo, popped) // visited popped room
		for _, k := range rooms[popped] {
			if toDo[k] {
				queue = append(queue, k)
			}
		}
	}
	return len(toDo) == 0
}
