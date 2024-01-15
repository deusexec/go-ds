package stack

import "testing"

func Test_Push(t *testing.T) {
	stack := New[int]()
	stack.Push(1)
	if stack.Length() != 1 {
		t.Error("Stack should contain 1 element")
	}
}

func Test_Pop(t *testing.T) {
	stack := New[int]()
	n := 1
	stack.Push(n)
	item := stack.Pop()
	if item != n {
		t.Error("Pop returns different element")
	}
}

func Test_Peek(t *testing.T) {
	stack := New[int]()
	n := 1
	stack.Push(n)
	if stack.Peek() != n {
		t.Error("Stack peek elements should be equals")
	}
}

func Test_IsEmpty(t *testing.T) {
	stack := New[int]()
	if stack.IsEmpty() != true {
		t.Error("Stack should be empty")
	}
	stack.Push(1)
	if stack.IsEmpty() == true {
		t.Error("Stack should not be empty")
	}
}

func Test_Length(t *testing.T) {
	stack := New[int]()
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
