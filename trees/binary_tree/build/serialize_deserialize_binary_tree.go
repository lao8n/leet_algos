package data_structures

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
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
	// Split the data into a list of tokens
	lst := strings.Split(data, ",")

	// Start deserialization from the beginning of the list
	idx := 0

	var helper func() *TreeNode
	helper = func() *TreeNode {
		// If the index is out of bounds or the token is "n", return nil
		if idx >= len(lst) || lst[idx] == "n" {
			idx++ // Move to the next token
			return nil
		}

		// Convert the token to an integer and move to the next token
		val, _ := strconv.Atoi(lst[idx])
		idx++

		// Recursively deserialize the left and right subtrees
		left := helper()
		right := helper()

		// Create a new TreeNode and return it
		return &TreeNode{Val: val, Left: left, Right: right}
	}

	// Start the deserialization process
	return helper()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
