package trie

type TrieNode struct {
	children   map[rune]*TrieNode
	isTerminal bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	trie := &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
	return trie
}

func (t *Trie) Insert(word string) *Trie {
	currentNode := t.root
	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isTerminal = true
	return t
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.isTerminal
}

func (t *Trie) Delete(word string) *Trie {
	t.root.delete(word)
	return t
}

func (n *TrieNode) delete(word string) {
	if len(word) == 0 {
		n.isTerminal = false
		return
	}

	char := rune(word[0])
	if child, ok := n.children[char]; ok {
		child.delete(word[1:])
		if len(child.children) == 0 && !child.isTerminal {
			delete(n.children, char)
		}
	}
}

func (t *Trie) PrefixSearch(prefix string) []string {
	currentNode := t.root
	for _, char := range prefix {
		if _, ok := currentNode.children[char]; !ok {
			return []string{}
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.words(prefix)
}

func (n *TrieNode) words(prefix string) []string {
	result := []string{}
	if n.isTerminal {
		result = append(result, prefix)
	}
	for char, child := range n.children {
		result = append(result, child.words(prefix+string(char))...)
	}
	return result
}
