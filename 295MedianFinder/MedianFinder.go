package main

import (
	"container/heap"
	"fmt"

	. "qianliout/leetcode/common/commonHeap"
)

func main() {
	mid := Constructor()
	mid.AddNum(1)
	mid.AddNum(2)
	mid.AddNum(3)
	median := mid.FindMedian()
	fmt.Println("mdeian is ", median)
}

type MedianFinder struct {
	Left  MaxHeap
	Right MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		Left:  make([]int, 0),
		Right: make([]int, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 先加左边，左边最多比右边多一个
	heap.Push(&this.Left, num)
	// 调整
	for (len(this.Left) > 0 && len(this.Right) > 0 && this.Left[0] > this.Right[0]) || (len(this.Left)-len(this.Right) > 1) {
		pop := heap.Pop(&this.Left).(int)
		heap.Push(&this.Right, pop)
	}
	for len(this.Right) > 0 && len(this.Left)-len(this.Right) < 0 {
		pop := heap.Pop(&this.Right).(int)
		heap.Push(&this.Left, pop)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	fmt.Println("left is ", this.Left)
	fmt.Println("right is ", this.Right)
	if len(this.Left) == 0 {
		return 0
	}
	if len(this.Left)-len(this.Right) == 1 {
		return float64(this.Left[0])
	}
	if len(this.Right) > 0 {
		return float64(this.Left[0]+this.Right[0]) / 2
	}
	return 0
}
