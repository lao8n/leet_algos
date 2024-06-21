// in order to determine an answer you need either
// 1. combination already calculated
// 2. inverse of combination
// 3. combination is connected e.g. a/b b/c so then can either multiply or inverse & multiply
// 4. solve for individual values? probably not

// build connected components graph where for each equation need both numerator and denominator in
// the same graph - actually calculating might be quite hard - how represent whether numerator/denominator?
// use optimised quick union with 1. memoized find 2. rank by union

// maybe this isn't correct as find is doing something different - it's finding the update node not the root
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	uf := UnionFind{
		nodes: make(map[string]*node),
	}

	for i, eq := range equations {
		for _, operand := range eq {
			if _, ok := uf.nodes[operand]; !ok {
				uf.nodes[operand] = &node{
					root:   operand,
					rank:   1,
					weight: 1.0,
				}
			}
		}
		uf.union(eq[0], eq[1], values[i])
	}

	results := make([]float64, len(queries))
	for i, q := range queries {
		results[i] = uf.calculateQuery(q)
	}
	return results
}

type node struct {
	root   string
	rank   int
	weight float64
}

type UnionFind struct {
	nodes map[string]*node
}

func (uf *UnionFind) find(x string) *node {
	if node, ok := uf.nodes[x]; !ok {
		return nil
	} else {
		if node.root != x {
			parent := uf.find(node.root)
			node.root = parent.root
			node.weight *= parent.weight
		}
		return node
	}
}

func (uf *UnionFind) union(x string, y string, value float64) {
	// where we are doing x / y
	xNode, yNode := uf.find(x), uf.find(y)
	if xNode.root != yNode.root {
		xRoot, yRoot := uf.nodes[xNode.root], uf.nodes[yNode.root]
		if yRoot.rank < xRoot.rank {
			xNode, yNode = yNode, xNode
			xRoot, yRoot = yRoot, xRoot
			x, y = y, x
			value = 1 / value
		}
		xRoot.root = yNode.root
		xRoot.weight = value * yNode.weight / xNode.weight
		xRoot.rank += yRoot.rank
		yRoot.rank += xRoot.rank
	}
}

func (uf *UnionFind) calculateQuery(query []string) float64 {
	xNode, yNode := uf.find(query[0]), uf.find(query[1])
	// operands not found
	if xNode == nil || yNode == nil {
		return -1
	}
	// no path from x to y
	if xNode.root != yNode.root {
		return -1
	}
	if query[0] == query[1] {
		return 1
	}
	return xNode.weight / yNode.weight
}