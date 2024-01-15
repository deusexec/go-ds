package stack

import "errors"

var (
	errorStackIsEmpty = errors.New("stack is empty")
)

type stack[T any] struct {
	items []T
}

func New[T any]() *stack[T] {
	return new(stack[T])
}

func (s *stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *stack[T]) Pop() T {
	if s.IsEmpty() {
		panic(errorStackIsEmpty)
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item
}

func (s *stack[T]) Peek() T {
	if s.IsEmpty() {
		panic(errorStackIsEmpty)
	}
	return s.items[len(s.items)-1]
}

func (s *stack[T]) Length() int {
	return len(s.items)
}

func (s *stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
