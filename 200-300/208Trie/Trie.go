package main

import (
	"fmt"
)

func main() {
	t := Constructor()
	t.Insert("apple")
	search := t.StartsWith("app")
	// search1 := t.Search("apple")
	fmt.Println("search ", search)
	// fmt.Println("search ", search1)
}

type Node struct {
	End  bool
	Next map[byte]*Node
}

func NewNode() *Node {
	return &Node{
		End:  false,
		Next: map[byte]*Node{},
	}
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{Root: NewNode()}
}

func (t *Trie) Insert(word string) {
	t.add(t.Root, []byte(word), 0)
}

func (t *Trie) add(node *Node, word []byte, index int) {
	if len(word) <= index {
		node.End = true
		return
	}
	c := word[index]
	if node.Next[c] == nil {
		node.Next[c] = NewNode()
	}
	t.add(node.Next[c], word, index+1)
}

func (t *Trie) find(node *Node, word []byte, index int) bool {
	if len(word) <= index {
		return node.End
	}
	c := word[index]
	if node.Next[c] == nil {
		return false
	}
	return t.find(node.Next[c], word, index+1)
}
func (t *Trie) findPre(node *Node, word []byte, index int) bool {
	if len(word) <= index {
		return true
	}
	c := word[index]
	if node.Next[c] == nil {
		return false
	}
	return t.findPre(node.Next[c], word, index+1)
}

func (t *Trie) Search(word string) bool {
	return t.find(t.Root, []byte(word), 0)
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.findPre(t.Root, []byte(prefix), 0)
}
