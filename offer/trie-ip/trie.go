package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tr := Constructor()

	for i := 0; i < 200; i++ {
		go func(i uint8) {
			tr.Insert([]uint8{1, 0, 0, 1, 0, uint8(i), 1, 0, 1})
			// fmt.Println("insert执行完成")
		}(uint8(i))
	}
	for i := 0; i < 500; i++ {
		go func(i uint8) {
			tr.Delete([]uint8{1, 0, 0, 1, 0, 1, 1, uint8(i), uint8(i + 3)})
			// fmt.Println("Delete执行完成")
		}(uint8(i))
	}

	for i := 0; i < 20; i++ {
		go func(i uint8) {
			search := tr.Search([]uint8{1, 0, 0, 1, 0, uint8(i), 1, 0, 1})
			// fmt.Println("Search执行完成")
			fmt.Println(search)
		}(uint8(i))
	}
	time.Sleep(5 * time.Second)
}

type Node struct {
	IsEnd bool
	Lock  sync.RWMutex
	// 因为这里只会有两种值，所以可使用数组的方式进行存储，也是一种优化的方式
	Next map[uint8]*Node
}

type Trie struct {
	Root *Node
}

func NewNode() *Node {
	return &Node{Next: make(map[uint8]*Node)}
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
		root.Lock.Lock()
		defer root.Lock.Unlock()
		root.IsEnd = true
		return
	}
	c := word[index]
	root.Lock.RLock()
	defer root.Lock.RUnlock()
	nex := root.Next[c]
	if nex == nil {
		root.Lock.RUnlock() // 先释放读锁
		root.Lock.Lock()    // 再加写锁
		nex = NewNode()
		root.Next[c] = nex
		root.Lock.Unlock() // 释放写锁，这里不能用defer,因为defer是在程序退出前执行，如果用defer就在会下面写锁之后，形成死锁
		root.Lock.RLock()  // 再加读锁

	}
	add(nex, word, index+1)
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word []uint8) bool {
	return findWord(this.Root, word, 0)
}

func findWord(root *Node, word []uint8, index int) bool {
	root.Lock.RLock()
	defer root.Lock.RUnlock()
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

// 删除节点
func (this *Trie) Delete(word []uint8) {
	deleteWord(this.Root, word, 0)
}

// 对于ip掩码来说，并发量大，且更新频繁，所以这里最好不直接删除节点，只是把节点的值改了，这样可以避免map频繁的等量扩容
func deleteWord(root *Node, word []uint8, index int) {
	if len(word) == index {
		root.Lock.Lock()
		defer root.Lock.Unlock()
		root.IsEnd = false
		return
	}
	c := word[index]
	root.Lock.RLock()
	defer root.Lock.RUnlock()
	if root.Next[c] == nil {
		return
	}
	nex := root.Next[c]
	deleteWord(nex, word, index+1)
}
