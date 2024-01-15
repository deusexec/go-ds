package bst

import (
	"fmt"

	"github.com/deusexec/go-ds/queue"
	"github.com/deusexec/go-ds/stack"
)

const (
	PREORDER = iota
	INORDER
	POSTORDER
	BFS
	DFS
)

func Traversal(order uint8, node *node) {
	switch order {
	case PREORDER:
		printPreOrder(node)
	case INORDER:
		printInOrder(node)
	case POSTORDER:
		printPostOrder(node)
	case BFS:
		printBFS(node)
	case DFS:
		printDFS(node)
	}
}

func printPreOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Printf("%c ", node.value)
	printPreOrder(node.left)
	printPreOrder(node.right)
}

func printInOrder(node *node) {
	if node == nil {
		return
	}
	printInOrder(node.left)
	fmt.Printf("%c ", node.value)
	printInOrder(node.right)
}

func printPostOrder(node *node) {
	if node == nil {
		return
	}
	printPostOrder(node.left)
	printPostOrder(node.right)
	fmt.Printf("%c ", node.value)
}

func printBFS(start *node) {
	if start == nil {
		return
	}
	queue := queue.New[*node]()
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		n := queue.Dequeue()
		fmt.Print(n, " ")

		if n.left != nil {
			queue.Enqueue(n.left)
		}
		if n.right != nil {
			queue.Enqueue(n.right)
		}
	}
}

func printDFS(start *node) {
	if start == nil {
		return
	}
	stack := stack.New[*node]()
	stack.Push(start)

	for !stack.IsEmpty() {
		n := stack.Pop()
		fmt.Print(n, " ")

		if n.right != nil {
			stack.Push(n.right)
		}
		if n.left != nil {
			stack.Push(n.left)
		}
	}
}
