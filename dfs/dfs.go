package dfs

import (
	"github.com/arberiii/Graph-Algorithms/graph"
)

func DFS(s int, g graph.Graph) graph.Graph {
	color := make(map[int]string)
	for k := range g {
		color[k] = "white"
	}
	tree := graph.NewGraph()
	dfsFromVertex(s, tree, g, color)
	return tree
}

func dfsFromVertex(u int, tree graph.Graph, g graph.Graph, color map[int]string) {
	color[u] = "grey"
	for _, v := range g[u] {
		if color[v] == "white" {
			tree[v] = append(tree[v], u)
			tree[u] = append(tree[u], v)
			dfsFromVertex(v, tree, g, color)
		}
	}
	color[u] = "black"
}
