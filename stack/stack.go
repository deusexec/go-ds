package stack

import "errors"

var (
	ErrorStackIsEmpty = errors.New("stack is empty")
)

type Interface interface {
	Push(item any)
	Pop() any
	Peek() any
	Length() int
	IsEmpty() bool
}

type stack struct {
	items []any
}

func New() *stack {
	return new(stack)
}

func (s *stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() any {
	if s.IsEmpty() {
		panic(ErrorStackIsEmpty)
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item
}

func (s *stack) Peek() any {
	if s.IsEmpty() {
		panic(ErrorStackIsEmpty)
	}
	return s.items[len(s.items)-1]
}

func (s *stack) Length() int {
	return len(s.items)
}

func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}
