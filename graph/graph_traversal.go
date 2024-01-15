package graph

import (
	"github.com/deusexec/go-ds/queue"
	"github.com/deusexec/go-ds/stack"
)

type traversalOrder uint

const (
	BFS traversalOrder = iota
	DFS
)

func (g *graph[K, T]) Traversal(order traversalOrder, start *Node[K, T], callback func(node *Node[K, T])) {
	switch order {
	case BFS:
		g.bfs(start, callback)
	case DFS:
		g.dfs(start, callback)
	default:
		panic("unexpected traversal type")
	}
}

func (g *graph[K, T]) bfs(start *Node[K, T], callback func(node *Node[K, T])) {
	queue := queue.New[*Node[K, T]]()
	visited := make(map[*Node[K, T]]bool, len(g.nodes))

	visited[start] = true
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		node := queue.Dequeue()
		callback(node)

		edges := g.nodes[node.key]
		for _, node := range edges {
			if !visited[node] {
				visited[node] = true
				queue.Enqueue(node)
			}
		}
	}
}

func (g *graph[K, T]) dfs(start *Node[K, T], callback func(node *Node[K, T])) {
	stack := stack.New[*Node[K, T]]()
	visited := make(map[*Node[K, T]]bool, len(g.nodes))

	visited[start] = true
	stack.Push(start)

	for !stack.IsEmpty() {
		node := stack.Pop()
		callback(node)

		edges := g.nodes[node.key]
		for _, node := range edges {
			if !visited[node] {
				visited[node] = true
				stack.Push(node)
			}
		}
	}
}
