package data_structures

import (
	"strconv"
	"strings"
)

type Codec struct {
	i     int
	parts []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	// base case
	if root == nil {
		return "n,"
	}
	// recursive case (pre-order traversal)
	return strconv.Itoa(root.Val) + "," +
		this.serialize(root.Left) +
		this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	this.i = 0 // reset
	this.parts = parts[:len(parts)-1]
	return this.deserializeRecursive()
}

// need to keep update i for every recursion even if just find a nil
func (this *Codec) deserializeRecursive() *TreeNode {
	// base case
	if this.i >= len(this.parts) || this.parts[this.i] == "n" {
		this.i++
		return nil
	}

	num, _ := strconv.Atoi(this.parts[this.i])
	this.i++
	// recursive case
	left := this.deserializeRecursive()
	right := this.deserializeRecursive()
	return &TreeNode{
		Val:   num,
		Left:  left,
		Right: right,
	}
}
