/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// two large prime numbers
const MOD_1 = 1000000007
const MOD_2 = 2147483647

// hashing approach
// if nil -> hash it to 3 (any prime is fine)
// else -> shift hash value of left node by fixed value, and right value by 1 then sum and add to node.val to get new hash

// to avoid minimum spurious hits we hash twice to two different values
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	h := Hashes{visited: make(map[HashPair]bool)}
	h.hashNode(root, true) // track visited
	subRootHash := h.hashNode(subRoot, false)
	_, exists := h.visited[subRootHash]
	return exists
}

type Hashes struct {
	visited map[HashPair]bool
}

func (h *Hashes) hashNode(node *TreeNode, addToMemo bool) HashPair {
	// base cases
	if node == nil {
		return HashPair{5, 7}
	}
	// calculate two hashes
	left := h.hashNode(node.Left, addToMemo)
	right := h.hashNode(node.Right, addToMemo)

	leftHash := ((left.h1<<5)%MOD_1 + (right.h1<<1)%MOD_1 + node.Val) % MOD_1
	rightHash := ((right.h2<<7)%MOD_2 + (right.h2<<1)%MOD_2 + node.Val) % MOD_2
	hashPair := HashPair{leftHash, rightHash}

	// memoize
	if addToMemo {
		h.visited[hashPair] = true
	}

	return hashPair
}

type HashPair struct {
	h1 int
	h2 int
}