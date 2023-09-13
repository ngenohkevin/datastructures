package main

import "fmt"

// Node represents a single node in the BST.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BST represents the Binary Search Tree.
type BST struct {
	Root *Node
}

// Insert inserts a new value into the BST
func (bst *BST) Insert(value int) {
	newNode := &Node{Value: value}

	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	insertNode(bst.Root, newNode)
}

// Search searches for a value in the BST.
func (bst *BST) Search(value int) bool {
	return searchNode(bst.Root, value)
}

// Helper function to recursively search for a value in the BST.
func searchNode(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	if value < node.Value {
		return searchNode(node.Left, value)
	}
	return searchNode(node.Right, value)
}

// Helper function to recursively inset a node into the BST.
func insertNode(currentNode, newNode *Node) {
	if newNode.Value < currentNode.Value {
		if currentNode.Left == nil {
			currentNode.Left = newNode
		} else {
			insertNode(currentNode.Left, newNode)
		}
	} else {
		if currentNode.Right == nil {
			currentNode.Right = newNode
		} else {
			insertNode(currentNode.Right, newNode)
		}
	}
}

// InOrderTraversal performs an in-order traversal of the BST.
func (bst *BST) InOrderTraversal() {
	inOrder(bst.Root)
	fmt.Println()
}

// Helper function to recursively perform in-order traversal.
func inOrder(node *Node) {
	if node == nil {
		return
	}

	inOrder(node.Left)
	fmt.Printf("%d ", node.Value)
	inOrder(node.Right)
}

func main() {
	bst := BST{}

	// Insert value into the BST
	values := []int{50, 30, 70, 20, 40, 60, 80}

	for _, value := range values {
		bst.Insert(value)
	}
	// Perform in-order traversal to verify the tree structure.
	fmt.Print("In-Order Traversal: ")
	bst.InOrderTraversal()

	value := 50

	found := bst.Search(value)
	if found {
		fmt.Printf("Value %d found in the BST.\n", value)
	} else {
		fmt.Printf("Value %d not found in the BST\n", value)
	}
}
