package graph

import "testing"

func Test_NewGraph(t *testing.T) {
	g := NewGraph[int, string]()

	if g.nodesCount != 0 {
		t.Error("Node count should be equals 0")
	}
	if g.edgesCount != 0 {
		t.Error("Edge count should be equals 0")
	}
	if len(g.nodes) != 0 {
		t.Error("Node list should be empty")
	}
}

func Test_AddNode(t *testing.T) {
	g := NewGraph[int, string]()

	nodeA := NewNode[int, string](1, "nodeA")
	nodeB := NewNode[int, string](2, "nodeB")

	g.AddNode(nodeA)
	g.AddNode(nodeB)

	expectedNodeCount := 2

	if g.NodeCount() != expectedNodeCount || len(g.nodes) != expectedNodeCount {
		t.Errorf("Node count should be equals %d, actual %d \n", expectedNodeCount, g.NodeCount())
	}
}

func Test_AddEdge(t *testing.T) {
	g := NewGraph[int, string]()

	nodeA := NewNode[int, string](1, "nodeA")
	nodeB := NewNode[int, string](2, "nodeB")

	g.AddNode(nodeA)
	g.AddNode(nodeB)

	g.AddEdge(nodeA, nodeB)

	expectedNodeCount := 2
	expectedEdgeCount := 1

	if g.NodeCount() != expectedNodeCount || len(g.nodes) != expectedNodeCount {
		t.Errorf("Node count should be equals %d, actual %d \n", expectedNodeCount, g.NodeCount())
	}
	if g.EdgeCount() != expectedEdgeCount {
		t.Errorf("Edge count should be equals %d, actual %d \n", expectedEdgeCount, g.EdgeCount())
	}
}

func Test_DeleteNode(t *testing.T) {
	g := NewGraph[int, string]()

	nodeA := NewNode[int, string](1, "nodeA")
	nodeB := NewNode[int, string](2, "nodeB")

	g.AddNode(nodeA)
	g.AddNode(nodeB)

	g.DeleteNode(nodeA)
	if g.NodeCount() != 1 || len(g.nodes) != 1 {
		t.Errorf("Node count should be equals %d, actual %d \n", 1, g.NodeCount())
	}

	g.DeleteNode(nodeB)
	if g.NodeCount() != 0 || len(g.nodes) != 0 {
		t.Errorf("Node count should be equals %d, actual %d \n", 1, g.NodeCount())
	}
}

func Test_DeleteEdge(t *testing.T) {
	g := NewGraph[int, string]()

	nodeA := NewNode[int, string](1, "nodeA")
	nodeB := NewNode[int, string](2, "nodeB")
	nodeC := NewNode[int, string](3, "nodeC")
	nodeD := NewNode[int, string](4, "nodeD")

	g.AddNode(nodeA)
	g.AddNode(nodeB)
	g.AddNode(nodeC)
	g.AddNode(nodeD)

	g.AddEdge(nodeA, nodeB)
	g.AddEdge(nodeC, nodeD)

	g.DeleteEdge(nodeA, nodeB)
	if g.EdgeCount() != 1 {
		t.Errorf("Edge count should be equals %d, actual %d \n", 1, g.EdgeCount())
	}

	g.DeleteEdge(nodeC, nodeD)
	if g.EdgeCount() != 0 {
		t.Errorf("Edge count should be equals %d, actual %d \n", 0, g.EdgeCount())
	}
}
