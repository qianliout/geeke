package main

import (
	"container/heap"
	"fmt"
)

func main() {
	lfu := Constructor(3)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	lfu.Put(3, 3)
	lfu.Put(4, 4)
	fmt.Printf("FrequencyMap %#v\n", lfu.FrequencyMap)
	fmt.Printf("valueMap %#v\n", lfu.ValueMap)
	fmt.Printf("current %#v\n", lfu.Current)

	//lfu.Get(4)
	////time.Sleep(time.Second)
	//lfu.Get(3)
	////time.Sleep(time.Second)
	//lfu.Get(2)
	////time.Sleep(time.Second)
	//lfu.Get(1)
	//lfu.PutNode(5, 5)
	////lfu.Get(4)
	//
	//fmt.Println(lfu.Get(1))
	//fmt.Println(lfu.Get(2))
	//fmt.Println(lfu.Get(3))
	//fmt.Println(lfu.Get(4))
	//fmt.Println(lfu.Get(5))
	//lfu.PutNode(3, 3)
	//fmt.Println(lfu.Get(2))
	//fmt.Println(lfu.Get(3))
	//lfu.PutNode(4, 4)
	//fmt.Println(lfu.Get(1))
	//fmt.Println(lfu.Get(3))
	//fmt.Println(lfu.Get(4))

}

/*
请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。它应该支持以下操作：get 和 put。
get(key) - 如果键存在于缓存中，则获取键的值（总是正数），否则返回 -1。
put(key, value) - 如果键已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量时，
则应该在插入新项之前，使最不经常使用的项无效。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除最久未使用的键。
「项的使用次数」就是自插入该项以来对其调用 get 和 put 函数的次数之和。使用次数会在对应项被移除后置为 0 。
进阶：
你是否可以在 O(1) 时间复杂度内执行两项操作？
示例：
LFUCache cache = new LFUCache( 2 capacity (缓存容量))
cache.put(1, 1)
cache.put(2, 2)
cache.get(1) // 返回 1
cache.put(3, 3) // 去除 key 2
cache.get(2) // 返回 -1 (未找到key 2)
cache.get(3) // 返回 3
cache.put(4, 4) // 去除 key 1
cache.get(1) // 返回 -1 (未找到 key 1)
cache.get(3) // 返回 3
cache.get(4) // 返回 4
*/
//type LFUCache struct {
//	Lfu     map[int]int
//	Pq      PriorityQueue
//	Cap     int
//	Current int
//}
//
//func Constructor(capacity int) LFUCache {
//	lfu := LFUCache{
//		Lfu:     make(map[int]int),
//		Pq:      make(PriorityQueue, 0),
//		Cap:     capacity,
//		Current: 0,
//	}
//	return lfu
//}
//
//func (this *LFUCache) Get(key int) int {
//	if this.Cap <= 0 {
//		return -1
//	}
//	// 取值
//	v, ok := this.Lfu[key]
//	if !ok {
//		return -1
//	}
//	// 查到值了
//	// 更新Pq
//	for j, i := range this.Pq {
//		if i.Key == key {
//			r := heap.Remove(&this.Pq, j).(*LFUItem)
//			r.Priority.Frequency += 1
//			r.Priority.LastTime = time.Now().Unix()
//			heap.Push(&this.Pq, r)
//			break
//		}
//	}
//	return v
//}
//
//func (this *LFUCache) PutNode(key int, value int) {
//	if this.Cap <= 0 {
//		return
//	}
//	// 第一种情况,加的老值
//	v := this.Get(key)
//	if v > -1 {
//		this.Lfu[key] = value
//		// 更新Pq
//		for j, i := range this.Pq {
//			if i.Key == key {
//				heap.Remove(&this.Pq, j)
//				i.Priority.Frequency += 1
//				i.Priority.LastTime = time.Now().Unix()
//				heap.Push(&this.Pq, i)
//				break
//			}
//		}
//		return
//	}
//	// 第二种方式,加的新值,但是还有空间
//	if this.Current < this.Cap {
//		this.Lfu[key] = value
//		r := LFUItem{
//			Key: key,
//			Priority: PriorityItem{
//				LastTime:  time.Now().Unix(),
//				Frequency: 1,
//			},
//		}
//		heap.Push(&this.Pq, &r)
//		this.Current += 1
//		return
//	}
//	// 第三种情况,加的新值,已没有空间了
//	// 先弄一个出去
//	r := heap.Remove(&this.Pq, 0).(*LFUItem)
//	delete(this.Lfu, r.Key)
//	this.Lfu[key] = value
//	newR := LFUItem{
//		Key: key,
//		Priority: PriorityItem{
//			LastTime:  time.Now().Unix(),
//			Frequency: 1,
//		},
//	}
//	heap.Push(&this.Pq, &newR)
//}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.PutNode(key,value);
 */

type PriorityItem struct {
	LastTime  int64
	Frequency int
}

type LFUItem struct {
	Key      int          // The Key of the item; arbitrary.
	Priority PriorityItem // The Priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*LFUItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	if pq[i].Priority.Frequency > pq[j].Priority.Frequency {
		return false
	} else if pq[i].Priority.Frequency < pq[j].Priority.Frequency {
		return true
	} else {
		return pq[i].Priority.LastTime < pq[j].Priority.LastTime // 最后使用的排在最前面些
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*LFUItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the Priority and Key of an LFUItem in the queue.
func (pq *PriorityQueue) update(item *LFUItem, value int, priority PriorityItem) {
	item.Key = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
