package bfs

import "github.com/arberiii/Graph-Algorithms/graph"

func BFS(g graph.Graph, s int) graph.Graph {
	tree := graph.NewGraph()
	var queue []int
	visited := make(map[int]bool)

	for k := range g {
		visited[k] = false
	}
	queue = add(queue, s)
	tree[s] = []int{}
	for len(queue) != 0 {
		u := delete(queue)
		for v := range g[u] {
			if visited[v] == false {
				add(queue, v)
				visited[v] = true
				tree[u] = append(tree[u], v)
			}
		}
	}
	return tree
}

func add(q []int, v int) []int {
	q = append(q, v)
	return q
}

func delete(q []int) int {
	v := q[0]
	q = q[1:]
	return v
}
