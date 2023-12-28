package queue

import "errors"

var (
	ErrorQueueIsEmpty = errors.New("queue is empty")
)

type Interface interface {
	Enqueue(item any)
	Dequeue() any
	Peek() any
	Length() int
	IsEmpty() bool
}

type queue struct {
	items []any
}

func New() *queue {
	return new(queue)
}

func (q *queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() any {
	if q.IsEmpty() {
		panic(ErrorQueueIsEmpty)
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *queue) Peek() any {
	if q.IsEmpty() {
		panic(ErrorQueueIsEmpty)
	}
	return q.items[0]
}

func (q *queue) Length() int {
	return len(q.items)
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}
