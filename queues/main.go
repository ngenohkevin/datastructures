package main

type Queue struct {
	items []int
}

// Enqueue (push) adds an element to the rear of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue (Pop) removes and returns the front element from the queue.
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("Queue is empty")
	}

	front := q.items[0]
	q.items = q.items[1:]

	return front
}

// Front (Peek) returns the front element without removing it.
func (q *Queue) Front() int {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.items[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}
