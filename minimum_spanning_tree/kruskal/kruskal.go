package kruskal

import "github.com/arberiii/Graph-Algorithms/graph"

type edge struct {
	u      int
	v      int
	weight float64
}

// we suppose that the weight are sorted
func Kruskal(g graph.Graph, w []*edge) graph.Graph {
	tree := graph.NewGraph()
	components := make(map[int]int)
	for v := range g {
		tree[v] = []int{}
		components[v] = v
	}

	for i := range w {
		if !hasCycle(w[i], components) {
			tree[w[i].u] = append(tree[w[i].v])
			for key, val := range components {
				if val == w[i].v {
					components[key] = w[i].u
				}
			}
			tree[w[i].v] = append(tree[w[i].u])
		}
	}
	return tree
}

func hasCycle(e *edge, c map[int]int) bool {
	return c[e.u] == c[e.v]
}
