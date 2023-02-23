package main

import (
	"fmt"
)

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 0)
	lRUCache.Put(2, 2)
	lRUCache.Get(1)
	lRUCache.Put(3, 3)
	fmt.Println(lRUCache.Get(2)) // -1
	lRUCache.Put(4, 4)
	fmt.Println(lRUCache.Get(1)) // -1
	fmt.Println(lRUCache.Get(3)) // 3
	fmt.Println(lRUCache.Get(4)) // 4

}

/*
LRU (最近最少使用) 缓存,不止要考虑最少，还要考虑最近,是两个条件
*/
type LRUCache struct {
	dataMap map[int]*DoublyLinkedList // 数据
	data    *DoublyLinkedList

	capacity int
	length   int
}

// 只能使用双向链表
func Constructor(capacity int) LRUCache {
	return LRUCache{
		dataMap:  map[int]*DoublyLinkedList{},
		data:     nil,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if pre, ok := this.dataMap[key]; ok {
		val := pre.Val
		node := &DoublyLinkedList{Key: key, Val: val}
		this.data = this.data.Remove(pre)
		this.data = this.data.PutNode(node)
		this.dataMap[key] = node
		return val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 存在
	node := &DoublyLinkedList{Key: key, Val: value}
	if pre, ok := this.dataMap[key]; ok {
		this.data = this.data.Remove(pre)
		this.data = this.data.PutNode(node)
		this.dataMap[key] = node
		return
	}
	// 不存在,且容量够
	if this.length < this.capacity {
		this.data = this.data.PutNode(node)
		this.dataMap[key] = node
		this.length++
	} else {
		// 不存在，且容量不够了，移出头部
		delete(this.dataMap, this.data.Key)
		this.data = this.data.RemoveFirstNode()
		this.data = this.data.PutNode(node)
		this.dataMap[key] = node
	}
}

type DoublyLinkedList struct {
	Pre  *DoublyLinkedList
	Next *DoublyLinkedList
	Val  int
	Key  int
}

// PutNode 在最开始添加一个节点
func (dl *DoublyLinkedList) PutNode(node *DoublyLinkedList) *DoublyLinkedList {
	if dl == nil {
		return node
	}
	dl.Pre = node
	node.Next = dl
	return node
}

// 移出一个node,如果没有这个key,就不移出
func (dl *DoublyLinkedList) Remove(node *DoublyLinkedList) *DoublyLinkedList {
	if dl == nil {
		return nil
	}
	if dl == node {
		return dl.Next
	}
	cur := dl
	for cur != nil {
		if cur == node {
			cur.Pre.Next = cur.Next
			break
		}
		cur = cur.Next
	}
	return dl
}

// 移出一个node,如果没有这个key,就不移出
func (dl *DoublyLinkedList) RemoveFirstNode() *DoublyLinkedList {
	if dl == nil {
		return dl
	}
	return dl.Next
}

// 移出一个node,如果没有这个key,就不移出
func (dl *DoublyLinkedList) RemoveLastNode() *DoublyLinkedList {
	if dl == nil || dl.Next == nil {
		return nil
	}
	cur := dl
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return dl
}

// 查看这个key是否存在
func (dl *DoublyLinkedList) ExitNode(node *DoublyLinkedList) bool {
	cur := dl

	for cur != nil {
		if cur == node {
			return true
		}
		cur = cur.Next
	}
	return false
}
