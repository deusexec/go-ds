package graph

import (
	"cmp"
	"fmt"
)

type Node[K cmp.Ordered, T any] struct {
	key  K
	data T
	next *Node[K, T]
}

func NewNode[K cmp.Ordered, T any](key K, data T) *Node[K, T] {
	return &Node[K, T]{key, data, nil}
}

func (n *Node[K, T]) Key() K {
	return n.key
}

func (n *Node[K, T]) Data() T {
	return n.data
}

func (n *Node[K, T]) Next() *Node[K, T] {
	return n.next
}

func (n *Node[K, T]) String() string {
	return fmt.Sprintf("[%v] ", n.key)
}
