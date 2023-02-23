package main

import (
	"container/heap"
	"fmt"
)

func main() {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(words, 3))
}

/*
给一非空的单词列表，返回前 k 个出现次数最多的单词。
返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
示例 1：
输入: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
输出: ["i", "love"]
解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
    注意，按字母顺序 "i" 在 "love" 之前。
示例 2：
输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
输出: ["the", "is", "sunny", "day"]
解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
    出现次数依次为 4, 3, 2 和 1 次。
注意：
假定 k 总为有效值， 1 ≤ k ≤ 集合元素数。
输入的单词均由小写字母组成。
扩展练习：
尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
*/
// 优先队列的用法
func topKFrequent(words []string, k int) []string {
	pq := make(PriorityQueue, 0)
	mm := make(map[string]int)
	for _, w := range words {
		mm[w] += 1
	}

	for k, v := range mm {
		heap.Push(&pq, &StringItem{
			Value:    k,
			Priority: v,
		})
	}
	res := make([]string, 0)
	for k > 0 && len(pq) > 0 {
		pop := heap.Pop(&pq).(*StringItem)
		res = append(res, pop.Value)
		k--
	}
	return res
}

type StringItem struct {
	Value    string // The Key of the item; arbitrary.
	Priority int    // The Priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*StringItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Priority > pq[j].Priority {
		return true
	} else if pq[i].Priority == pq[j].Priority {
		return pq[i].Value > pq[j].Value
	}
	return false
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*StringItem)
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
