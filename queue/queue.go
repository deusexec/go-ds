package queue

import "errors"

var (
	errorQueueIsEmpty = errors.New("queue is empty")
)

type queue[T any] struct {
	items []T
}

func New[T any]() *queue[T] {
	return new(queue[T])
}

func (q *queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic(errorQueueIsEmpty)
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *queue[T]) Peek() T {
	if q.IsEmpty() {
		panic(errorQueueIsEmpty)
	}
	return q.items[0]
}

func (q *queue[T]) Length() int {
	return len(q.items)
}

func (q *queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
