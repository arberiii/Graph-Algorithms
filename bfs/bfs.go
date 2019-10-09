package bfs

import (
	"github.com/arberiii/Graph-Algorithms/graph"
)

type queue struct {
	l []int
}

func BFS(s int, g graph.Graph) graph.Graph {
	tree := graph.NewGraph()
	q := newQueue()
	visited := make(map[int]bool)

	for k := range g {
		visited[k] = false
	}
	q.add(s)
	visited[s] = true
	tree[s] = []int{}
	for len(q.l) != 0 {
		u := q.deleteV()
		for _, v := range g[u] {
			if visited[v] == false {
				q.add(v)
				visited[v] = true
				tree[u] = append(tree[u], v)
				tree[v] = append(tree[v], u)
			}
		}
	}
	return tree
}

func (q *queue) add(v int) {
	q.l = append(q.l, v)
}

func (q *queue) deleteV() int {
	u := q.l[0]
	q.l = q.l[1:]
	return u
}

func newQueue() queue {
	var q queue
	return q
}
