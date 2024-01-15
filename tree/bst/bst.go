package bst

import "cmp"

func New[K cmp.Ordered, T any](key K, data T) *Node[K, T] {
	return &Node[K, T]{
		key:  key,
		data: data,
	}
}

func Search[K cmp.Ordered, T any](src, dst *Node[K, T]) *Node[K, T] {
	if src == nil {
		return nil
	} else if dst.key < src.key {
		return Search(src.left, dst)
	} else if dst.key > src.key {
		return Search(src.right, dst)
	}
	return src
}

func Insert[K cmp.Ordered, T any](src, dst *Node[K, T]) *Node[K, T] {
	if src == nil {
		return New(dst.key, dst.data)
	} else if dst.key < src.key {
		src.left = Insert(src.left, dst)
	} else if dst.key > src.key {
		src.right = Insert(src.right, dst)
	}
	return src
}
