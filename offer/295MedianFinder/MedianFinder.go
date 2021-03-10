package main

import (
	"container/heap"
)

func main() {
	
}

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	
	return MedianFinder{
		minHeap: make([]int, 0),
		maxHeap: make([]int, 0),
	}
	
}

// 保证左边最多只比右边多一个,左边是大，右边是小
func (this *MedianFinder) AddNum(num int) {
	if len(this.minHeap) > 0 && num > this.minHeap[0] {
		heap.Push(&this.minHeap, num)
		if len(this.minHeap) > len(this.maxHeap) {
			pop := heap.Pop(&this.minHeap).(int)
			heap.Push(&this.maxHeap, pop)
		}
	} else {
		heap.Push(&this.maxHeap, num)
		if len(this.maxHeap)-len(this.minHeap) > 1 {
			pop := heap.Pop(&this.maxHeap).(int)
			heap.Push(&this.minHeap, pop)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxHeap) == 0 {
		return 0
	}
	if len(this.minHeap) == 0 {
		return float64(this.maxHeap[0])
	}
	if len(this.maxHeap)-len(this.minHeap) == 1 {
		return float64(this.maxHeap[0])
	}
	return float64(this.minHeap[0]+this.maxHeap[0]) / 2
}

/*
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
例如，
[2,3,4] 的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
void addNum(int num) - 从数据流中添加一个整数到数据结构中。
double findMedian() - 返回目前所有元素的中位数。
*/

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

type MinHeap []int

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
