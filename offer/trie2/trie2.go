package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tr := Constructor()

	for i := 0; i < 5; i++ {
		go func(i uint8) {
			tr.Insert([]uint8{1, 0, 0, 1, 0, uint8(i), 1, 0})
		}(uint8(i))
	}
	for i := 0; i < 5; i++ {
		go func(i uint8) {
			tr.Delete([]uint8{1, 0, 0, 1, 0, 1, 1, uint8(i)})
		}(uint8(i))
	}

	for i := 0; i < 5; i++ {
		go func(i uint8) {
			search := tr.Search([]uint8{1, 0, 0, 1, 0, 1, 1, uint8(i)})
			fmt.Println(search)
		}(uint8(i))
	}
	time.Sleep(2 * time.Second)

}

type Node struct {
	IsEnd bool
	Lock  sync.RWMutex

	Next sync.Map
}

type Trie struct {
	Root *Node
}

func NewNode() *Node {
	return &Node{Next: sync.Map{}}
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Root: NewNode()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word []uint8) {
	add(this.Root, word, 0)
}

func add(root *Node, word []uint8, index int) {
	if len(word) == index {
		root.IsEnd = true
		return
	}
	c := word[index]
	_, ok := root.Next.Load(c)
	if !ok {
		root.Next.Store(c, NewNode())
	}
	nex, ok := root.Next.Load(c)
	if ok {
		if next, ok := nex.(*Node); ok {
			add(next, word, index+1)
		} else {
			panic("构建trie出错")
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word []uint8) bool {
	return findWord(this.Root, word, 0)
}

func findWord(root *Node, word []uint8, index int) bool {
	if len(word) == index {
		return root.IsEnd
	}
	c := word[index]
	load, ok := root.Next.Load(c)
	if !ok {
		return false
	}

	if next, ok := load.(*Node); ok {
		return findWord(next, word, index+1)
	} else {
		panic("查找trie出错")
	}
}

// 删除节点
func (this *Trie) Delete(word []uint8) {
	deleteWord(this.Root, word, 0)
}

// 对于ip掩码来说，并发量大，且更新频繁，所以这里最好不直接删除节点，只是把节点的值改了，这样可以避免map频繁的等量扩容(rehash)
func deleteWord(root *Node, word []uint8, index int) {
	if len(word) == index {
		root.Lock.RLock()
		defer root.Lock.Unlock()
		root.IsEnd = false
		return
	}
	c := word[index]
	load, ok := root.Next.Load(c)
	if !ok {
		return
	}

	if next, ok := load.(*Node); ok {
		deleteWord(next, word, index+1)
	} else {
		panic("删除trie出错")
	}
}
