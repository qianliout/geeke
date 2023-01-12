package main

import (
	"fmt"
)

func main() {
	m := Constructor()
	m.BuildDict([]string{"hello", "hallo", "leetcode"})

	/*

		k["MagicDictionary", "buildDict", "search", "search", "search", "search"]
		[[], [["hello","hallo","leetcode"]], ["hello"], ["hallo"], ["hell"], ["leetcode"]]*/
	fmt.Println(m.Search("hello"))
	fmt.Println(m.Search("hallo"))
	fmt.Println(m.Search("hell"))
	fmt.Println(m.Search("leetcode"))
}

type Node struct {
	IsEnd bool
	Next  map[byte]*Node
}

type MagicDictionary struct {
	Root *Node
}

func Constructor() MagicDictionary {
	return MagicDictionary{Root: NewNode()}
}

func NewNode() *Node {
	return &Node{Next: make(map[byte]*Node)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := range dictionary {
		this.Root.Add(dictionary[i])
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	word := []byte(searchWord)
	for i := 0; i < len(word); i++ {
		if query(this.Root, word, 0, i) {
			return true
		}
	}
	return false
}

func query(root *Node, word []byte, idx, changedIdx int) bool {
	if root == nil {
		return false
	}
	if idx == len(word) {
		return root.IsEnd
	}
	ch := word[idx]
	if idx == changedIdx {
		for i := 'a'; i <= 'z'; i++ {
			if byte(i) == ch {
				continue
			}
			if query(root.Next[byte(i)], word, idx+1, changedIdx) {
				return true
			}
		}
		return false
	}
	return query(root.Next[ch], word, idx+1, changedIdx)
}

func (tr *Node) Add(word string) {
	add(tr, []byte(word), 0)
}

func add(node *Node, word []byte, index int) {
	if index == len(word) {
		node.IsEnd = true
		return
	}
	c := word[index]
	next := node.Next[c]
	if next == nil {
		next = NewNode()
		node.Next[c] = next
	}
	add(next, word, index+1)
}

func findWord(root *Node, word []byte, index int) bool {
	if len(word) == index {
		return root.IsEnd
	}
	c := word[index]
	if root.Next[c] == nil {
		return false
	}
	nex := root.Next[c]
	return findWord(nex, word, index+1)
}
