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

func printBFS(tree *node) {
	queue := queue.New()
	var temp *node

	if tree != nil {
		queue.Enqueue(tree)
	}

	for !queue.IsEmpty() {
		item := queue.Dequeue()
		temp = item.(*node)

		fmt.Print(temp, " ")

		if temp.left != nil {
			queue.Enqueue(temp.left)
		}
		if temp.right != nil {
			queue.Enqueue(temp.right)
		}
	}
}

func printDFS(tree *node) {
	stack := stack.New()
	var temp *node

	if tree != nil {
		stack.Push(tree)
	}

	for !stack.IsEmpty() {
		item := stack.Pop()
		temp = item.(*node)

		fmt.Print(temp, " ")

		if temp.right != nil {
			stack.Push(temp.right)
		}
		if temp.left != nil {
			stack.Push(temp.left)
		}
	}
}
