package main

const ArraySize = 7

// HashTable represents a basic hash table with key-value pairs.
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket represents a linked list of key-value pairs.
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
}
