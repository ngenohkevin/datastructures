package main

// HashTable represents a basic hash table with key-value pairs.
type HashTable struct {
	table map[int]*bucket
}

// bucket represents a linked list of key-value pairs.
