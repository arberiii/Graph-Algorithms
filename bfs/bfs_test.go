package bfs

import (
	"testing"

	"github.com/arberiii/Graph-Algorithms/graph"
)

func TestBFS(t *testing.T) {
	g := graph.NewGraph()
	g = map[int][]int{
		0:  {1, 2, 3},
		1:  {0, 4},
		2:  {0, 3, 4},
		3:  {0, 6, 7},
		4:  {1, 2, 8},
		5:  {2, 8},
		6:  {3, 9, 10},
		7:  {3, 10},
		8:  {4, 5},
		9:  {6, 10},
		10: {6, 7},
	}
	tree := BFS(g, 0)
	length := 0
	for i := range tree {
		length += len(tree[i])
	}
	if length != 10 {
		t.Fatal("length of tree should be 10")
	}
}
