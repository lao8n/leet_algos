// package main
package data_structures

import (
	"fmt"
)

// adapted from algorithm 6.6 from bengio's deep learning txtbook p213
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
	}
	fmt.Printf("losses %v\n", losses)
}

// forward map[1:3 2:-4 3:6 4:-3 5:3 6:0.3333333333333333 7:-0.36666666666666664]
// backward map[1:19.799999999999997 2:-13.2 3:6.6 4:3.3 5:3.3 6:3.3 7:-0.36666666666666664]

// forward map[1:1.0200000000000002 2:-2.6799999999999997 3:2.0400000000000005 4:-1.6799999999999997 5:0.36000000000000076 6:2.777777777777772 7:2.077777777777772]
// backward map[1:-0.5493312000000009 2:0.721670400000001 3:-0.5385600000000008 4:-0.2692800000000004 5:-0.2692800000000004 6:-0.2692800000000004 7:2.077777777777772]

// forward map[1:1.0749331200000003 2:-2.75216704 3:2.1498662400000006 4:-1.7521670399999998 5:0.3976992000000008 6:2.514463192282001 7:1.8144631922820011]
// backward map[1:-0.6169770892384993 2:0.7898277473473586 3:-0.5739678848471048 4:-0.2869839424235524 5:-0.2869839424235524 6:-0.2869839424235524 7:1.8144631922820011]

// forward map[1:1.1366308289238503 2:-2.8311498147347356 3:2.2732616578477005 4:-1.8311498147347356 5:0.4421118431129649 6:2.2618710979531214 7:1.5618710979531214]
// backward map[1:-0.6939991090862233 2:0.8643155715632834 3:-0.6105756516768898 4:-0.3052878258384449 5:-0.3052878258384449 6:-0.3052878258384449 7:1.5618710979531214]

// forward map[1:1.2060307398324726 2:-2.917581371891064 3:2.412061479664945 4:-1.917581371891064 5:0.4944801077738812 6:2.022326043613641 7:1.3223260436136413]
// backward map[1:-0.779874239542203 2:0.9433202979644221 3:-0.6466454077700655 4:-0.32332270388503276 5:-0.32332270388503276 6:-0.32332270388503276 7:1.3223260436136413]

// forward map[1:1.2840181637866928 2:-3.011913401687506 3:2.5680363275733855 4:-2.011913401687506 5:0.5561229258858793 6:1.798163595588231 7:1.098163595588231]
// backward map[1:-0.8721873905307265 2:1.0229422621932442 3:-0.6792640595975389 4:-0.33963202979876944 5:-0.33963202979876944 6:-0.33963202979876944 7:1.098163595588231]

// forward map[1:1.3712369028397653 2:-3.1142076279068305 3:2.7424738056795306 4:-2.1142076279068305 5:0.6282661777727001 6:1.5916820535288931 7:0.8916820535288932]
// backward map[1:-0.9652501436439805 2:1.0960868081762494 3:-0.7039266093590351 4:-0.35196330467951753 5:-0.35196330467951753 6:-0.35196330467951753 7:0.8916820535288932]

// forward map[1:1.4677619172041634 2:-3.2238163087244556 3:2.9355238344083268 4:-2.2238163087244556 5:0.7117075256838712 6:1.4050715552559487 7:0.7050715552559488]
// backward map[1:-1.0483877106166646 2:1.1513479671792617 3:-0.714276408406661 4:-0.3571382042033305 5:-0.3571382042033305 6:-0.3571382042033305 7:0.7050715552559488]

// forward map[1:1.5726006882658299 2:-3.3389511054423817 3:3.1452013765316598 4:-2.3389511054423817 5:0.806250271089278 6:1.2403096604841546 7:0.5403096604841546]
// backward map[1:-1.104665872130284 2:1.172715159803618 3:-0.7024452426944081 4:-0.35122262134720406 5:-0.35122262134720406 6:-0.35122262134720406 7:0.5403096604841546]

// forward map[1:1.6830672754788583 2:-3.4562226214227434 3:3.3661345509577165 4:-2.4562226214227434 5:0.9099119295349731 6:1.0990074616463903 7:0.39900746164639034]
// backward map[1:-1.1120164371691437 2:1.1417774029396472 3:-0.6607082517558652 4:-0.3303541258779326 5:-0.3303541258779326 6:-0.3303541258779326 7:0.39900746164639034]

// losses [-0.36666666666666664 2.077777777777772 1.8144631922820011 1.5618710979531214 1.3223260436136413 1.098163595588231 0.8916820535288932 0.7050715552559488 0.5403096604841546 0.39900746164639034]
