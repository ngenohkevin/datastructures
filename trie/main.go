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

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)

	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func main() {

}
