package data_structures

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(value int) error {
	if n == nil {
		return fmt.Errorf("node is nil")
	}
	if value == n.value {
		return fmt.Errorf("value %d already exists", value)
	} else if value < n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			return n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			return n.right.Insert(value)
		}
	}
	return nil
}

func (n *Node) FindMin() int {
	if n.left == nil {
		return n.value
	} else {
		return n.left.FindMin()
	}
}

func (n *Node) FindMax() int {
	if n.right == nil {
		return n.value
	} else {
		return n.right.FindMax()
	}
}

func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}
	// 3 things: recurse left, recurse right & do something
	// in order traversal: 1. left 2. do something 3. right
	// pre-order traversal: 1. do something 2. left 3. right
	// post-order traversal: 1. left 2. right 3. do something
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Print(n.value)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}
