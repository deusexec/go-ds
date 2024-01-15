package queue

import "testing"

func Test_Enqueue(t *testing.T) {
	queue := New[int]()
	n := 1
	queue.Enqueue(n)
	if queue.Length() != 1 {
		t.Error("Queue length should be equals 1")
	}
}

func Test_Dequeue(t *testing.T) {
	queue := New[int]()
	n := 1
	queue.Enqueue(n)
	item := queue.Dequeue()
	if queue.Length() != 0 {
		t.Error("Queue length should be equals 0")
	}
	if item != n {
		t.Error("Queue elements should match")
	}
}

func Test_Peek(t *testing.T) {
	queue := New[int]()
	n := 1
	queue.Enqueue(n)
	if queue.Peek() != n {
		t.Error("Queue elements should match")
	}
}

func Test_IsEmpty(t *testing.T) {
	queue := New[int]()
	if queue.IsEmpty() != true {
		t.Error("Queue should be empty")
	}
	queue.Enqueue(1)
	if queue.IsEmpty() == true {
		t.Error("Queue should not be empty")
	}
}

func Test_Length(t *testing.T) {
	queue := New[int]()
	if queue.Length() != 0 {
		t.Error("Queue length should be equals 0")
	}
	queue.Enqueue(1)
	if queue.Length() != 1 {
		t.Error("Queue length should be equals 1")
	}
	queue.Dequeue()
	if queue.Length() != 0 {
		t.Error("Queue length should be equals 0")
	}
}
