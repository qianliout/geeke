package main

func main() {

}

type Node struct {
	IsWord bool
	Next   map[string]*Node
}

type Trie struct {
	Root *Node
}

func NewNode() *Node {
	return &Node{Next: make(map[string]*Node)}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Root: NewNode()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	add(this.Root, word, 0)
}

func add(root *Node, word string, index int) {
	if len(word) == index {
		root.IsWord = true
		return
	}
	c := string([]rune(word)[index])
	if root.Next[c] == nil {
		root.Next[c] = NewNode()
	}
	add(root.Next[c], word, index+1)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return findWord(this.Root, word, 0)
}

func findWord(root *Node, word string, index int) bool {
	if len(word) == index {
		return root.IsWord
	}
	c := string([]rune(word)[index])
	if root.Next[c] == nil {
		return false
	}
	return findWord(root.Next[c], word, index+1)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return findPrefix(this.Root, prefix, 0)
}

func findPrefix(root *Node, word string, index int) bool {
	if len(word) == index {
		return true
	}
	c := string([]rune(word)[index])
	if root.Next[c] == nil {
		return false
	}
	return findPrefix(root.Next[c], word, index+1)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Query(word);
 * param_3 := obj.StartsWith(prefix);
 */
