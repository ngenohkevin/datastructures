package main

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

type Node struct {
	children [26]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// init trie will create a new trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}

	return result
}

func main() {

}
