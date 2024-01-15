package bst

import (
	"cmp"
	"fmt"
)

type Node[K cmp.Ordered, T any] struct {
	key   K
	data  T
	left  *Node[K, T]
	right *Node[K, T]
}

func (n *Node[K, T]) Key() K {
	return n.key
}

func (n *Node[K, T]) Data() T {
	return n.data
}

func (n *Node[K, T]) Left() *Node[K, T] {
	return n.left
}

func (n *Node[K, T]) Right() *Node[K, T] {
	return n.right
}

func (n *Node[K, T]) String() string {
	return fmt.Sprintf("{ %v }", n.data)
}
