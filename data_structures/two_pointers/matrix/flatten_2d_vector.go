package data_structures

// approaches
// * flatten in constructor
// * two pointers

// implemenation
// * should hasNext move pointers?
type Vector2D struct {
	vec [][]int
	i   int
	j   int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{
		vec: vec,
		i:   0,
		j:   0,
	}
}

func (this *Vector2D) Next() int {
	if this.HasNext() {
		next := this.vec[this.i][this.j]
		this.j++
		return next
	}
	return 0
}

func (this *Vector2D) HasNext() bool {
	for this.i < len(this.vec) {
		// nested is
		if this.checkNested() {
			return true
		}
		// nested has no more
		this.i++
		this.j = 0 // reset
	}
	return false
}

func (this *Vector2D) checkNested() bool {
	if this.j < len(this.vec[this.i]) {
		return true
	}
	return false
}

/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(vec);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
