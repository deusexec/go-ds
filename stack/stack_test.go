package stack

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	var stack Stack
	stack.Push(1)
	if stack.Length() != 1 {
		t.Error("Stack should contain 1 element")
	}
}

func TestStackPop(t *testing.T) {
	var stack Stack
	n := 1
	stack.Push(n)
	item := stack.Pop()
	if item != n {
		t.Error("Pop returns different element")
	}
}

func TestStackPeek(t *testing.T) {
	var stack Stack
	n := 1
	stack.Push(n)
	if stack.Peek() != n {
		t.Error("Stack peek elements should be equals")
	}
}

func TestStackIsEmpty(t *testing.T) {
	var stack Stack
	if stack.IsEmpty() != true {
		t.Error("Stack should be empty")
	}
	stack.Push(1)
	if stack.IsEmpty() == true {
		t.Error("Stack should not be empty")
	}
}

func TestStackLength(t *testing.T) {
	var stack Stack
	if stack.Length() != 0 {
		t.Error("Length should be equals 0")
	}
	stack.Push(1)
	if stack.Length() != 1 {
		t.Error("Length should be equals 1")
	}
	stack.Pop()
	if stack.Length() != 0 {
		t.Error("Length should be equals 0")
	}
}
