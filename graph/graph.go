package graph

type Graph map[int][]int

func NewGraph() Graph {
	g := make(Graph)
	return g
}
