package main

import "fmt"

// Stack represents stack that hold a slice

type Stack struct {
	items []int
}

// push will add value at the end
func (s *Stack) Push(i int) {

	s.items = append(s.items, i)
}

// pop will remove value at the end
// and returns the removed value
func (s *Stack) Pop() int {

	if s.IsEmpty() {
		panic("Stack is empty")
	}

	l := len(s.items) - 1

	top := s.items[l]
	s.items = s.items[:l]
	return top
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
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	myStack := Stack{}

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Push(400)
	myStack.Push(50)

	fmt.Println("Full stack :", myStack)
	fmt.Println("Peek :", myStack.Peek())
	fmt.Println("Pop :", myStack.Pop())

	fmt.Println("Full stack after pop:", myStack)

	// Check if the stack is empty
	fmt.Println("Is Empty:", myStack.IsEmpty())

	// Get the size of the stack
	fmt.Println("Size :", myStack.Size())
}
