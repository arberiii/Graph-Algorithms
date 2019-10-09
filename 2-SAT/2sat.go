package sat

import (
	"errors"

	"github.com/arberiii/Graph-Algorithms/bfs"
	"github.com/arberiii/Graph-Algorithms/graph"
)

type clause struct {
	x int
	y int
}

func Sat2(n int, clauses []clause) (map[int]bool, error) {
	ret := make(map[int]bool)
	determined := make(map[int]bool)
	for i := 1; i < n+1; i++ {
		determined[i] = false
	}
	digraph := formGraph(n, clauses)
	S := make(map[int][]int)
	for v := range digraph {
		b := bfs.BFS(v, digraph)
		for vertices := range b {
			S[v] = append(S[v], vertices)
		}
	}
	for i := 1; i < n+1; i++ {
		if xInS(i, S) {
			return nil, errors.New("NO")
		}
	}

	for len(determined) != 0 {
		for k := range determined {
			if !xCInS(k, S) {
				for _, v := range S[k] {
					ret[v] = true
					ret[-v] = false
					delete(determined, abs(v))
				}
			} else {
				for _, v := range S[-k] {
					ret[v] = true
					ret[-v] = false
					delete(determined, abs(v))
				}
			}
		}
	}

	for i := 1; i < n+1; i++ {
		delete(ret, -i)
	}
	return ret, nil
}

func formGraph(n int, clauses []clause) graph.Graph {
	g := graph.NewGraph()
	// vertex set
	for i := 1; i < n+1; i++ {
		g[i] = []int{}
		g[-i] = []int{}
	}

	// adding directed edges
	for _, c := range clauses {
		g[-c.x] = append(g[-c.x], c.y)
		g[-c.y] = append(g[-c.y], c.x)
	}

	return g
}

func xInS(x int, S map[int][]int) bool {
	b1 := false
	for _, v := range S[x] {
		if v == -x {
			b1 = true
		}
	}
	b2 := false
	for _, v := range S[-x] {
		if v == x {
			b2 = true
		}
	}
	return b1 && b2
}

func xCInS(x int, S map[int][]int) bool {
	for _, v := range S[x] {
		if v == -x {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
