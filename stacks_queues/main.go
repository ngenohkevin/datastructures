package main

import "fmt"

// Stack represents stack that hold a slice

type Stack struct {
	items []int
}

// push
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// pop

func main() {
	myStack := Stack{}
	fmt.Println(myStack)

	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
}
