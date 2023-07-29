package queue

import "errors"

var (
	ErrorQueueIsEmpty = errors.New("Queue is empty")
)

type QueueInterface interface {
	Enqueue(item any)
	Dequeue() any
	Peek() any
	Length() int
	IsEmpty() bool
}

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		panic(ErrorQueueIsEmpty)
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Peek() any {
	if q.IsEmpty() {
		panic(ErrorQueueIsEmpty)
	}
	return q.items[0]
}

func (q *Queue) Length() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
