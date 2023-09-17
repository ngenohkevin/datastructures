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
	charIndex := char - 'a'
	node := t.root
	if node.children[charIndex] == nil {
		node.children[charIndex] = &Node{}
	}
	node = node.children[charIndex]
	node.isEnd = true
}

// Search searches for a character in the Trie
func (t *Trie) Search(char rune) bool {
	charIndex := char - 'a'
	node := t.root
	if node.children[charIndex] == nil {
		return false
	}
	node = node.children[charIndex]
	return node.isEnd
}

func main() {
	// call the trie
}
