package data_structures

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */
// one approach is to store peekedValue calling next() to get it, or nil
type PeekingIterator struct {
	nextValue int
	iterator  *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	pI := PeekingIterator{iterator: iter}
	pI.next()
	return &pI
}

func (this *PeekingIterator) hasNext() bool {
	return this.nextValue != -1
}

func (this *PeekingIterator) next() int {
	returnValue := this.nextValue
	if this.iterator.hasNext() {
		this.nextValue = this.iterator.next()
	} else {
		this.nextValue = -1
	}
	return returnValue
}

func (this *PeekingIterator) peek() int {
	return this.nextValue
}
