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

func Traversal[T cmp.Ordered](order traversalOrder, node *Node[T], callback func(node *Node[T])) {
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

func preorder[T cmp.Ordered](node *Node[T], callback func(node *Node[T])) {
	if node == nil {
		return
	}
	callback(node)
	preorder(node.left, callback)
	preorder(node.right, callback)
}

func inorder[T cmp.Ordered](node *Node[T], callback func(node *Node[T])) {
	if node == nil {
		return
	}
	inorder(node.left, callback)
	callback(node)
	inorder(node.right, callback)
}

func postorder[T cmp.Ordered](node *Node[T], callback func(node *Node[T])) {
	if node == nil {
		return
	}
	postorder(node.left, callback)
	postorder(node.right, callback)
	callback(node)
}

func bfs[T cmp.Ordered](start *Node[T], callback func(node *Node[T])) {
	if start == nil {
		return
	}
	queue := queue.New[*Node[T]]()
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

func dfs[T cmp.Ordered](start *Node[T], callback func(node *Node[T])) {
	if start == nil {
		return
	}
	stack := stack.New[*Node[T]]()
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
