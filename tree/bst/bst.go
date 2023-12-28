package bst

type BstValue byte

type BstInterface interface {
	New(value BstValue) *node
	Search(node *node, value BstValue)
	Insert(node *node, value BstValue)
}

func New(value BstValue) *node {
	return &node{
		value: value,
	}
}

func Search(node *node, value BstValue) *node {
	if node == nil {
		return nil
	} else if value < node.value {
		return Search(node.left, value)
	} else if value > node.value {
		return Search(node.right, value)
	}
	return node
}

func Insert(node *node, value BstValue) *node {
	if node == nil {
		return New(value)
	} else if value < node.value {
		node.left = Insert(node.left, value)
	} else if value > node.value {
		node.right = Insert(node.right, value)
	}
	return node
}
