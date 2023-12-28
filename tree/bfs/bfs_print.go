package bfs

import (
	"fmt"

	"github.com/deusexec/go-ds/queue"
	"github.com/deusexec/go-ds/stack"
)

func BFS(tree *node) {
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

func DFS(tree *node) {
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
