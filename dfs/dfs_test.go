package dfs

import (
	"testing"

	"github.com/arberiii/Graph-Algorithms/graph"
)

func TestDFS(t *testing.T) {
	g := graph.NewGraph()
	g = map[int][]int{
		0:  {1, 2, 3},
		1:  {0, 4},
		2:  {0, 3, 4},
		3:  {0, 2, 6, 7},
		4:  {1, 2, 8},
		5:  {2, 8},
		6:  {3, 9, 10},
		7:  {3, 10},
		8:  {4, 5},
		9:  {6, 10},
		10: {6, 7},
	}
	tree := DFS(0, g)
	length := 0
	for i := range tree {
		length += len(tree[i])
	}
	// n-1 edges, but degree should be 2*(n-1)
	if length != 20 {
		t.Fatalf("length of tree should be 20, but got %d", length)
	}

}
