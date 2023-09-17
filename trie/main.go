package main

// AlphabetSize is the number of possible characters in the trie
const alphabetSize = 26

// Node represents a single node in the Trie.
type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
}

// Trie represents the Trie data structure.
type Trie struct {
	root *Node
}

// Constructor function to create a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

// Insert will take in a char and add it to the trie
func (t *Trie) Insert(char rune) {
	wordLength := char - 'a'
	node := t.root
	if node.children[wordLength] == nil {
		node.children[wordLength] = &Node{}
	}
	node = node.children[wordLength]
	node.isEnd = true
}

// Search searches for a character in the Trie
func (t *Trie) Search(char rune) bool {
	wordLength := char - 'a'
	node := t.root
	if node.children[wordLength] == nil {
		return false
	}
	node = node.children[wordLength]
	return node.isEnd
}

func main() {
	// call the trie
}
