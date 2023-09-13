package main

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
