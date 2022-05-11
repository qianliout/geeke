package main

import (
	"fmt"

	"outback/leetcode/back/common/listnode"
)

func main() {
	lru := Constructor(1)
	lru.Put(2, 1)
	fmt.Println(lru.Get(2))

	lru.Put(3, 2)
	lru.Put(4, 2)
	lru.Put(5, 2)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
	lru.Put(3, 3)
	lru.Put(4, 1)
	fmt.Println(lru.Get(2))
}

/*
 运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。
 获取数据 get(key) - 如果关键字 (key) 存在于缓存中，则获取关键字的值（总是正数），否则返回 -1。
 写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字/值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
 进阶:
 你是否可以在 O(1) 时间复杂度内完成这两种操作？
 示例:
 LRUCache cache = new LRUCache( 2 缓存容量)
cache.put(1, 1)
cache.put(2, 2)
cache.get(1)    // 返回  1
cache.put(3, 3) // 该操作会使得关键字 2 作废
cache.get(2)    // 返回 -1 (未找到)
cache.put(4, 4) // 该操作会使得关键字 1 作废
cache.get(1)    // 返回 -1 (未找到)
cache.get(3)    // 返回  3
cache.get(4)    // 返回  4
*/

// 如查要O(1)的操作,那就只能是 hash表和双向链表

type LRUCache struct {
	CacheMap  map[int]*listnode.DoubleLinkedNode
	CacheList *listnode.DoubleLinkedNode
	Capacity  int
	Size      int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		CacheMap: make(map[int]*listnode.DoubleLinkedNode),
		Capacity: capacity,
		Size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.CacheMap[key]
	if !ok {
		return -1
	}
	val := v.Val
	this.Put(key, val)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	node := listnode.NewDoubleLinkedNode(key, value)
	if this.CacheList == nil {
		this.CacheList = node
		this.CacheMap[key] = node
		this.Size++
		return
	}

	n, ok := this.CacheMap[key]
	if ok {
		this.CacheList = this.CacheList.Remove(n)
		this.CacheList = this.CacheList.AddFirst(node)
		n.Val = value
	} else {
		this.CacheList = this.CacheList.AddFirst(node)
		this.Size++
		if this.Capacity < this.Size {
			last := this.CacheList.PopLast()
			delete(this.CacheMap, last.Key)
			this.Size--
		}
	}
	this.CacheMap[key] = node
	this.CacheList.Print("PutNode ")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.PutNode(key,value);
 */
