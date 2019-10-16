package shortest_path_dag

import (
	"math"

	"github.com/arberiii/Graph-Algorithms/graph"
)

// it is assumed that we have a topological order, graph is adjacency list with N^-! and we have weight
func DAG(g, w graph.Graph, order []int) graph.Graph {
	d := make(map[int]float64)
	pi := make(map[int]int)
	for i := 0; i < len(order); i++ {
		d[order[i]] = math.Inf(1)
		pi[order[i]] = 0
	}
	d[order[0]] = 0
	for i := 1; i < len(order)+1; i++ {
		for c, j := range g[order[i]] {
			if d[j]+float64(w[order[i]][c]) < d[order[i]] {
				d[order[i]] = d[j] + float64(w[order[i]][c])
				pi[order[i]] = j
			}
		}
	}
	ret := graph.NewGraph()

	for _, i := range order {
		ret[i] = []int{pi[i]}
	}

	return ret
}
