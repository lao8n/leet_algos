package data_structures

// clarifying questions
// * x and y can be < 0

// approaches
// * add two rectangles and minus overlap
// * could try to add areas separately but this seems harders
// * for overlap top right hand corner can you just take minimums (what about neg numbers?) and for bottom left take maxs?
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// a area
	aArea := calcArea(ax1, ay1, ax2, ay2)
	// b area
	bArea := calcArea(bx1, by1, bx2, by2)
	// overlap area
	maxX1, maxY1 := max(ax1, bx1), max(ay1, by1)
	minX2, minY2 := min(ax2, bx2), min(ay2, by2)
	overlapArea := calcArea(maxX1, maxY1, minX2, minY2)
	return aArea + bArea - overlapArea
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}

func min(x1, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

func calcArea(x1, y1, x2, y2 int) int {
	// area = width * height
	if x2 >= x1 && y2 >= y1 {
		return (x2 - x1) * (y2 - y1)
	}
	return 0
}
