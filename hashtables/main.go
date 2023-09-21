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
	key  string
	next *bucketNode
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
