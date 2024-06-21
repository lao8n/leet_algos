package data_structures

// approach
// * use a tree data structure where each letter is a node

// specifics
// * we use * to indicate the end of a word
type Trie struct {
	root *node
}

type node struct {
	c        byte
	children map[byte]*node
}

func Constructor() Trie {
	return Trie{
		root: &node{
			c:        ' ',
			children: make(map[byte]*node),
		},
	}
}

func (this *Trie) Insert(word string) {
	insert(word+"*", this.root)
}

func insert(word string, parent *node) {
	// base case
	if len(word) == 0 {
		return
	}
	// recursive case
	c := word[0]
	if child, ok := parent.children[c]; ok {
		insert(word[1:], child)
	} else {
		parent.children[c] = &node{
			c:        c,
			children: make(map[byte]*node),
		}
		insert(word[1:], parent.children[c])
	}

}

func (this *Trie) Search(word string) bool {
	return search(word, this.root)
}

func search(word string, parent *node) bool {
	if len(word) == 0 {
		_, ok := parent.children['*']
		return ok
	}

	c := word[0]
	if child, ok := parent.children[c]; ok {
		return search(word[1:], child)
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return startsWith(prefix, this.root)
}

func startsWith(word string, parent *node) bool {
	if len(word) == 0 {
		return true
	}

	c := word[0]
	if child, ok := parent.children[c]; ok {
		return startsWith(word[1:], child)
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
