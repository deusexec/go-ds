package graph

import (
	"cmp"
	"fmt"
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
			panic(fmt.Sprintf("node %v already presents in the graph's with a key %v", n, n.key))
		}
		nodes[n.key] = []*Node[K, T]{}
	}
	return &graph[K, T]{nodes, len(list), 0}
}

func (g *graph[K, T]) AddNode(n *Node[K, T]) {
	g.nodes[n.key] = []*Node[K, T]{}
	g.nodesCount++
}

func (g *graph[K, T]) AddEdge(src, dst *Node[K, T]) {
	g.nodes[src.key] = append(g.nodes[src.key], dst)
	g.edgesCount++
}

func (g *graph[K, T]) NodeCount() int {
	return g.nodesCount
}

func (g *graph[K, T]) EdgeCount() int {
	return g.edgesCount
}
