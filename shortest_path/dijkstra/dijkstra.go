package dijkstra

import (
	"math"

	"github.com/arberiii/Graph-Algorithms/graph"
)

func Dijasktra(s int, g graph.Graph, w map[int][]float64) map[int]float64 {
	c := make(map[int]float64)
	for v := range g {
		c[v] = math.Inf(1)
	}
	c[s] = 0
	var S []int
	for len(S) != len(g) {
		x := minimum(S, g, c)
		S = append(S, x)
		for i := range g[x] {
			if notInS(g[x][i], S) {
				if c[x]+w[x][i] < c[g[x][i]] {
					c[g[x][i]] = c[x] + w[x][i]
				}
			}
		}
	}

	return c
}

func minimum(S []int, g graph.Graph, c map[int]float64) int {
	ret := 0
	m := math.Inf(1)
	for v := range g {
		if notInS(v, S) {
			if c[v] < m {
				m = c[v]
				ret = v
			}
		}
	}
	return ret
}

func notInS(v int, S []int) bool {
	for _, k := range S {
		if k == v {
			return false
		}
	}
	return true
}
