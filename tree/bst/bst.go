package bst

import "cmp"

func New[T cmp.Ordered](data T) *Node[T] {
	return &Node[T]{
		data: data,
	}
}

func Search[T cmp.Ordered](node *Node[T], data T) *Node[T] {
	if node == nil {
		return nil
	} else if data < node.data {
		return Search(node.left, data)
	} else if data > node.data {
		return Search(node.right, data)
	}
	return node
}

func Insert[T cmp.Ordered](node *Node[T], data T) *Node[T] {
	if node == nil {
		return New(data)
	} else if data < node.data {
		node.left = Insert(node.left, data)
	} else if data > node.data {
		node.right = Insert(node.right, data)
	}
	return node
}
