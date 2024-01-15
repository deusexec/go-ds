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
	ei := g.findEdge(src, dst)
	if ei != -1 {
		return
	}
	g.nodes[src.key] = append(g.nodes[src.key], dst)
	g.edgesCount++
}

func (g *graph[K, T]) DeleteNode(node *Node[K, T]) {
	if _, ok := g.nodes[node.key]; !ok {
		return
	}
	for key, edges := range g.nodes {
		for i, n := range edges {
			if node.key == n.key {
				g.nodes[key] = append(g.nodes[key][:i], g.nodes[key][i+1:]...)
				g.edgesCount--
			}
		}
	}
	g.edgesCount = g.edgesCount - len(g.nodes[node.key])
	g.nodesCount--

	delete(g.nodes, node.key)
}

func (g *graph[K, T]) DeleteEdge(src, dst *Node[K, T]) {
	ei := g.findEdge(src, dst)
	if ei == -1 {
		return
	}
	edges := g.nodes[src.key]
	g.nodes[src.key] = append(edges[:ei], edges[ei+1:]...)
	g.edgesCount--
}

func (g *graph[K, T]) NodeCount() int {
	return g.nodesCount
}

func (g *graph[K, T]) EdgeCount() int {
	return g.edgesCount
}
