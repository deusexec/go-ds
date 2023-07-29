package queue

import "testing"

func TestEnqueue(t *testing.T) {
	var queue Queue
	n := 1
	queue.Enqueue(n)
	if queue.Length() != 1 {
		t.Error("Queue length should be equals 1")
	}
}

func TestDequeue(t *testing.T) {
	var queue Queue
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

func TestPeek(t *testing.T) {
	var queue Queue
	n := 1
	queue.Enqueue(n)
	if queue.Peek() != n {
		t.Error("Queue elements should match")
	}
}

func TestIsEmpty(t *testing.T) {
	var queue Queue
	if queue.IsEmpty() != true {
		t.Error("Queue should be empty")
	}
	queue.Enqueue(1)
	if queue.IsEmpty() == true {
		t.Error("Queue should not be empty")
	}
}

func TestLength(t *testing.T) {
	var queue Queue
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
