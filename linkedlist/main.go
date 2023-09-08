package main

import "fmt"

type node struct {
	next *node
	data int
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n

	l.head.next = second
	l.length++

}

// delete the node by value
func (l *linkedList) deleteByValue(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	current := l.head
	var prev *node

	for current != nil && current.data != value {
		prev = current
		current = current.next
	}

	if current != nil {
		prev.next = current.next
		l.length--
	}
}
func (l *linkedList) display() {
	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	mylist := linkedList{}

	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)

	mylist.deleteByValue(18)

	mylist.display()
}
