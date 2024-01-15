package graph

import (
	"fmt"
	"slices"
)

func (g *graph[K, T]) findEdge(src, dst *Node[K, T]) int {
	edges, ok := g.nodes[src.key]
	if !ok {
		return -1
	}
	for i, e := range edges {
		if e.key == dst.key {
			return i
		}
	}
	return -1
}

func (g *graph[K, T]) sortedNodes() []K {
	keys := []K{}
	for k := range g.nodes {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func (g *graph[K, T]) Dump() {
	fmt.Println("===============")
	fmt.Println("==   Graph   ==")
	fmt.Println("===============")

	for _, key := range g.sortedNodes() {
		edges := g.nodes[key]
		fmt.Printf("Node: [%v] Edges: ", key)
		if len(edges) == 0 {
			fmt.Print("-")
		} else {
			for _, edge := range edges {
				fmt.Printf("[%v]->[%v] ", key, edge.key)
			}
		}
		fmt.Println()
	}

	fmt.Println("===============")
}
