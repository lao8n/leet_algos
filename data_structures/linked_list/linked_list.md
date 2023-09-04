**Linked List**

Definition = a linked list is a data structure where each node in the list also has a reference to the next node in the list. Advantages over an array including dynamic sizing and cheap insertion/deletion

Time complexity = `O(1)` for insertion and deletion at the beginning, `O(n)` for traversal and search

Space complexity = `O(n)`

**Types**

Linked list = keep calling `next()` until no longer `hasNext()`
```
// peeking_iterator.go
if this.iterator.hasNext() {
    this.nextValue = this.iterator.next()
}
```
Can also use for `lru_cache.go` where least recently used is removed from the front and new values put at end but need not only next pointer but pointers to head and tail as well.