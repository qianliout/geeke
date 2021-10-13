package main

import (
	"container/heap"
	"fmt"

	"outback/leetcode/back/common/commonHeap"
)

func main() {

	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	m.AddNum(3)
	fmt.Println(m.leftHeap)
	fmt.Println(m.rightHeap)
	fmt.Println(m.FindMedian())

}

/*
中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
例如，
[2,3,4] 的中位数是 3
[2,3] 的中位数是 (2 + 3) / 2 = 2.5
设计一个支持以下两种操作的数据结构：
    void addNum(int num) - 从数据流中添加一个整数到数据结构中。
    double findMedian() - 返回目前所有元素的中位数。
示例：
addNum(1)
addNum(2)
findMedian() -> 1.5
addNum(3)
findMedian() -> 2
进阶:
    如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
    如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
*/

/** initialize your data structure here. */

// 解法就是维护两个堆，而且不删除，相对简单

type MedianFinder struct {
	leftHeap  commonHeap.MaxHeap
	rightHeap commonHeap.MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		leftHeap:  make(commonHeap.MaxHeap, 0),
		rightHeap: make(commonHeap.MinHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 左边比右边多一个
	if len(this.leftHeap) == 0 || len(this.leftHeap) > 0 && this.leftHeap[0] >= num {
		heap.Push(&this.leftHeap, num)
	} else {
		heap.Push(&this.rightHeap, num)
	}
	// 调整

	for len(this.leftHeap)-len(this.rightHeap) > 1 {
		v := heap.Pop(&this.leftHeap).(int)
		heap.Push(&this.rightHeap, v)
	}
	for len(this.rightHeap) > len(this.leftHeap) {
		v := heap.Pop(&this.rightHeap).(int)
		heap.Push(&this.leftHeap, v)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.leftHeap)-len(this.rightHeap) == 1 {
		return float64(this.leftHeap[0])
	}
	if len(this.leftHeap) == len(this.rightHeap) && len(this.leftHeap) > 0 && len(this.rightHeap) > 0 {
		return float64(this.leftHeap[0]+this.rightHeap[0]) / 2.0
	}
	return 0
}
