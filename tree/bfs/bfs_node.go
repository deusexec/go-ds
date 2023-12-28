package bfs

import "fmt"

type BstNodeInterface interface {
	Left() *node
	Right() *node
	Value() BstValue
}

type node struct {
	value BstValue
	left  *node
	right *node
}

func (n *node) Left() *node {
	return n.left
}

func (n *node) Right() *node {
	return n.right
}

func (n *node) Value() BstValue {
	return n.value
}

func (n *node) String() string {
	return fmt.Sprintf("%d", n.value)
}
