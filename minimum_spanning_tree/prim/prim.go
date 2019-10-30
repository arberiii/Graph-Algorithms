package prim

import (
	"math"

	"github.com/arberiii/Graph-Algorithms/graph"
)

type edge struct {
	u int
	v int
}

func Prim(g graph.Graph, w map[edge]float64) graph.Graph {
	tree := graph.NewGraph()
	for v := range g {
		tree[v] = []int{}
		break
	}
	for len(tree) != len(g) {
		e := lightest(tree, g, w)
		tree[e.u] = append(tree[e.u], e.v)
		tree[e.v] = append(tree[e.v], e.u)
	}
	return tree
}

func lightest(t, g graph.Graph, w map[edge]float64) edge {
	var ret edge
	weight := math.Inf(1)
	for x := range t {
		for y := range g[x] {
			if notIn(y, t) {
				if w[edge{u: x, v: y}] < weight {
					weight = w[edge{u: x, v: y}]
					ret = edge{u: x, v: y}
				}
			}
		}
	}
	return ret
}

func notIn(v int, t graph.Graph) bool {
	for k := range t {
		if k == v {
			return false
		}
	}
	return true
}
