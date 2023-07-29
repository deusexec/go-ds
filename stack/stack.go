package stack

import "errors"

var (
	ErrorStackIsEmpty = errors.New("Stack is empty")
)

type Interface interface {
	Push(item any)
	Pop() any
	Peek() any
	Length() int
	IsEmpty() bool
}

type Stack struct {
	items []any
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		panic(ErrorStackIsEmpty)
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		panic(ErrorStackIsEmpty)
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) Length() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
