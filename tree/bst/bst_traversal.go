package bst

import (
	"cmp"

	"github.com/deusexec/go-ds/queue"
	"github.com/deusexec/go-ds/stack"
)

type traversalOrder uint

const (
	PREORDER traversalOrder = iota
	INORDER
	POSTORDER
	BFS
	DFS
)

func Traversal[K cmp.Ordered, T any](order traversalOrder, node *Node[K, T], callback func(node *Node[K, T])) {
	switch order {
	case PREORDER:
		preorder(node, callback)
	case INORDER:
		inorder(node, callback)
	case POSTORDER:
		postorder(node, callback)
	case BFS:
		bfs(node, callback)
	case DFS:
		dfs(node, callback)
	}
}

func preorder[K cmp.Ordered, T any](node *Node[K, T], callback func(node *Node[K, T])) {
	if node == nil {
		return
	}
	callback(node)
	preorder(node.left, callback)
	preorder(node.right, callback)
}

func inorder[K cmp.Ordered, T any](node *Node[K, T], callback func(node *Node[K, T])) {
	if node == nil {
		return
	}
	inorder(node.left, callback)
	callback(node)
	inorder(node.right, callback)
}

func postorder[K cmp.Ordered, T any](node *Node[K, T], callback func(node *Node[K, T])) {
	if node == nil {
		return
	}
	postorder(node.left, callback)
	postorder(node.right, callback)
	callback(node)
}

func bfs[K cmp.Ordered, T any](start *Node[K, T], callback func(node *Node[K, T])) {
	if start == nil {
		return
	}
	queue := queue.New[*Node[K, T]]()
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		n := queue.Dequeue()
		callback(n)

		if n.left != nil {
			queue.Enqueue(n.left)
		}
		if n.right != nil {
			queue.Enqueue(n.right)
		}
	}
}

func dfs[K cmp.Ordered, T any](start *Node[K, T], callback func(node *Node[K, T])) {
	if start == nil {
		return
	}
	stack := stack.New[*Node[K, T]]()
	stack.Push(start)

	for !stack.IsEmpty() {
		n := stack.Pop()
		callback(n)

		if n.right != nil {
			stack.Push(n.right)
		}
		if n.left != nil {
			stack.Push(n.left)
		}
	}
}
