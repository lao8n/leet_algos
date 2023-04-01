package main

import (
	"fmt"
)

// package data_structures

type graph struct {
	graph      map[int]*node
	oper_table map[int]float64
	grad_table map[int]float64
}

type node struct {
	v         int
	operation func([]float64) float64
	bpop      func(float64) float64
	inputs    []*node
	outputs   []*node
}

func (g *graph) forward_pass(n *node) float64 {
	// tabulated value
	if op, ok := g.oper_table[n.v]; ok {
		return op
	}

	args := []float64{}
	for _, i := range n.inputs {
		args = append(args, g.forward_pass(i))
	}

	result := n.operation(args)
	g.oper_table[n.v] = result
	return result
}

func (g *graph) backward_pass(n *node) float64 {
	// tabulated value
	if gr, ok := g.grad_table[n.v]; ok {
		return gr
	}

	//
	n_gr := 0.0
	for _, o := range n.outputs {
		o_gr := n.bpop(g.oper_table[n.v]) * g.backward_pass(o)
		n_gr += o_gr
	}
	g.grad_table[n.v] = n_gr
	return n_gr
}

func main() {
	input1 := node{
		v:         1,
		operation: func(_ []float64) float64 { return 3 },
		bpop:      func(x float64) float64 { return x },
	}
	input2 := node{
		v:         2,
		operation: func(_ []float64) float64 { return -4 },
		bpop:      func(x float64) float64 { return x },
	}
	double := node{
		v:         3,
		operation: func(x []float64) float64 { return x[0] * 2 },
		bpop:      func(x float64) float64 { return 2 },
	}
	addOne := node{
		v:         4,
		operation: func(x []float64) float64 { return x[0] + 1 },
		bpop:      func(x float64) float64 { return 1 },
	}
	sum := node{
		v: 5,
		operation: func(x []float64) float64 {
			sum := 0.0
			for i := 0; i < len(x); i++ {
				sum += x[i]
			}
			return sum
		},
		bpop: func(x float64) float64 { return 1 },
	}
	inverse := node{
		v:         6,
		operation: func(x []float64) float64 { return 1 / x[0] },
		bpop:      func(x float64) float64 { return -1 / (x * x) },
	}
	loss := node{
		v:         7,
		operation: func(x []float64) float64 { return x[0] - 0.7 }, // simulating a loss function
	}
	// inputs & outputs
	input1.outputs = []*node{&double}
	input2.outputs = []*node{&addOne}
	double.inputs = []*node{&input1}
	double.outputs = []*node{&sum}
	addOne.inputs = []*node{&input2}
	addOne.outputs = []*node{&sum}
	sum.inputs = []*node{&double, &addOne}
	sum.outputs = []*node{&inverse}
	inverse.inputs = []*node{&sum}
	inverse.outputs = []*node{&loss}
	loss.inputs = []*node{&inverse}

	// graph setup
	g_nodes := make(map[int]*node)
	g_nodes[1] = &input1
	g_nodes[2] = &input2
	g_nodes[3] = &double
	g_nodes[4] = &addOne
	g_nodes[5] = &sum
	g_nodes[6] = &inverse
	g_nodes[7] = &loss
	g := graph{
		graph: g_nodes,
	}

	// forward pass
	num_epochs := 10
	step_size := 0.1
	losses := []float64{}
	for i := 0; i < num_epochs; i++ {
		g.oper_table = make(map[int]float64)
		g.grad_table = make(map[int]float64)

		for _, node := range g.graph {
			g.forward_pass(node)
		}
		losses = append(losses, g.oper_table[7])
		fmt.Printf("forward %v\n", g.oper_table)

		g.grad_table[7] = g.oper_table[7] //simulating loss function gradient
		// backward pass
		for _, node := range g.graph {
			g.backward_pass(node)
		}
		fmt.Printf("backward %v\n", g.grad_table)

		// update weights
		i1 := g.oper_table[1] - step_size*g.grad_table[1]
		g.graph[1].operation = func(_ []float64) float64 { return i1 }
		i2 := g.oper_table[2] - step_size*g.grad_table[2]
		g.graph[2].operation = func(_ []float64) float64 { return i2 }
		fmt.Println(g.oper_table[1], step_size*g.grad_table[1], i1)
	}
	fmt.Printf("%v\n", losses)
}
