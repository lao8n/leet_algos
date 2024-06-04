package data_structures

// approaches
// * diff in points - do diagonal first (as faster), then remaining x or y

func minTimeToVisitAllPoints(points [][]int) int {
	time := 0
	for i := 1; i < len(points); i++ {
		time += timeBetweenPoints(points[i-1], points[i])
	}
	return time
}

func timeBetweenPoints(p1 []int, p2 []int) int {
	// diagonal first
	yDelta := abs(p2[1] - p1[1])
	xDelta := abs(p2[0] - p1[0])
	diag := min(xDelta, yDelta)
	// remaining horizontal & vertical
	remX := xDelta - diag
	remY := yDelta - diag
	return diag + remX + remY
}

func abs(d int) int {
	if d < 0 {
		return -d
	}
	return d
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
