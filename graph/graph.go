package graph

import (
	"cmp"
)

type graph[K cmp.Ordered, T any] struct {
	nodes      map[K][]*Node[K, T]
	nodesCount int
	edgesCount int
}

func NewGraph[K cmp.Ordered, T any](list ...*Node[K, T]) *graph[K, T] {
	nodes := make(map[K][]*Node[K, T], len(list))
	for _, n := range list {
		if _, ok := nodes[n.key]; ok {
			continue
		}
		nodes[n.key] = []*Node[K, T]{}
	}
	return &graph[K, T]{nodes, len(list), 0}
}

func (g *graph[K, T]) AddNode(n *Node[K, T]) {
	if _, ok := g.nodes[n.key]; ok {
		return
	}
	g.nodes[n.key] = []*Node[K, T]{}
	g.nodesCount++
}

func (g *graph[K, T]) AddEdge(src, dst *Node[K, T]) {
	// same here, add check...
	g.nodes[src.key] = append(g.nodes[src.key], dst)
	g.edgesCount++
}

func (g *graph[K, T]) NodeCount() int {
	return g.nodesCount
}

func (g *graph[K, T]) EdgeCount() int {
	return g.edgesCount
}
