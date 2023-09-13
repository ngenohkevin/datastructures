package main

import "fmt"

// Stack represents stack that hold a slice

type Stack struct {
	items []int
}

// push will add value at the end
func (s *Stack) Push(i int) {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	s.items = append(s.items, i)
}

// pop will remove value at the end
// and returns the removed value
func (s *Stack) Pop() int {

	if s.IsEmpty() {
		panic("Stack is empty")
	}

	l := len(s.items) - 1

	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

// Peek returns the top element without removing it
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}

	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in a stack
func (s *Stack) Size() {

}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)
}
