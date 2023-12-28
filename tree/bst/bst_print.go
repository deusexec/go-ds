package bst

import "fmt"

const (
	PREORDER = iota
	INORDER
	POSTORDER
)

func Print(order uint8, node *node) {
	switch order {
	case PREORDER:
		printPreOrder(node)
	case INORDER:
		printInOrder(node)
	case POSTORDER:
		printPostOrder(node)
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
