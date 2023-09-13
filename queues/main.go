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

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
