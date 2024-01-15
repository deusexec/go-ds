package bst

import (
	"cmp"
	"fmt"
)

type Node[T cmp.Ordered] struct {
	data  T
	left  *Node[T]
	right *Node[T]
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func (n *Node[T]) Data() T {
	return n.data
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("{ %v }", n.data)
}
