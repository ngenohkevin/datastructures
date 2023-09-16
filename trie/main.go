package main

// AlphabetSize is the number of possible characters in the trie
const alphabetSize = 26

// TrieNode represents a single node in the Trie.
type TrieNode struct {
	children [alphabetSize]*TrieNode
	isEnd    bool
}

// Trie represents the Trie data structure.
type Trie struct {
	root *TrieNode
}

// Constructor function to create a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

// Insert will take in a char and add it to the trie
func (t *Trie) Insert(char rune) {
	charIndex := char - 'a'
	node := t.root
	if node.children[charIndex] == nil {
		node.children[charIndex] = &TrieNode{}
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
