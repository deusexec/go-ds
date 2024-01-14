package graph

import (
	"cmp"
	"fmt"

	"github.com/deusexec/go-ds/queue"
	"github.com/deusexec/go-ds/stack"
)

type TraversalType int
type TraversalCallback[K cmp.Ordered, T any] func(node *Node[K, T])

const (
	BFS TraversalType = iota
	DFS
)

func (g *graph[K, T]) Traversal(t TraversalType, start *Node[K, T], callback TraversalCallback[K, T]) {
	switch t {
	case BFS:
		g.bfs(start, callback)
	case DFS:
		g.dfs(start, callback)
	default:
		panic("unexpected traversal type")
	}
}

func (g *graph[K, T]) bfs(start *Node[K, T], callback TraversalCallback[K, T]) {
	queue := queue.New()
	visited := make(map[*Node[K, T]]bool, len(g.nodes))

	visited[start] = true
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		node := queue.Dequeue().(*Node[K, T])
		callback(node)

		edges := g.nodes[node.key]
		for _, node := range edges {
			if !visited[node] {
				visited[node] = true
				queue.Enqueue(node)
			}
		}
	}

	fmt.Println()
}

func (g *graph[K, T]) dfs(start *Node[K, T], callback TraversalCallback[K, T]) {
	stack := stack.New()
	visited := make(map[*Node[K, T]]bool, len(g.nodes))

	visited[start] = true
	stack.Push(start)

	for !stack.IsEmpty() {
		node := stack.Pop().(*Node[K, T])
		callback(node)

		edges := g.nodes[node.key]
		for _, node := range edges {
			if !visited[node] {
				visited[node] = true
				stack.Push(node)
			}
		}
	}

	fmt.Println()
}
