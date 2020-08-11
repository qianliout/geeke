package main

import (
	"container/heap"
	"fmt"

	heap2 "outback/leetcode/common/commonHeap"
)

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())

	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
	obj.AddNum(0)
	fmt.Println(obj.FindMedian())
	//obj.AddNum(3)

	//fmt.Println(obj.FindMedian())
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
// 使用数组实现，插入排序 fixme 插入排序的代码没有写对，要要自己的模板
//type MedianFinder struct {
//	Data []int
//}

/** initialize your data structure here. */
//func Constructor() MedianFinder {
//	m := new(MedianFinder)
//	m.Data = make([]int, 0)
//
//	return *m
//}
//
//func (this *MedianFinder) AddNum(num int) {
//	n := len(this.Data) - 1
//	left, right := 0, n
//	index := 0
//	if len(this.Data) > 0 {
//		if num < this.Data[left] {
//			index = 0
//		} else if num > this.Data[right] {
//			index = right + 1
//		} else {
//			for left <= right {
//				mid := left + (right-left)/2
//				if this.Data[mid] == num {
//					index = mid + 1
//					break
//				} else if mid-left == 1 && this.Data[left] < num && this.Data[mid] > num {
//					index = left + 1
//					break
//				} else if right-mid == 1 && this.Data[mid] < num && this.Data[right] > 0 {
//					index = mid + 1
//					break
//				} else if this.Data[mid] < num {
//					left = mid + 1
//				} else if this.Data[mid] > num {
//					right = mid - 1
//				}
//			}
//		}
//	}
//	values := append([]int{}, this.Data[:index]...)
//	values = append(values, num)
//	values = append(values, this.Data[index:]...)
//	this.Data = values
//}
//
//func (this *MedianFinder) FindMedian() float64 {
//	n := len(this.Data)
//	if n == 0 {
//		return 0
//	} else if n&1 == 1 {
//		mid := (n - 1) / 2
//		return float64(this.Data[mid])
//	} else {
//		left, right := (n-1)/2, n/2
//		return float64(this.Data[left]+this.Data[right]) / 2
//	}
//}

// 使用两个堆,这个代码可以优化，因为对于大顶堆index=0是最大值，对于小顶堆index=0就是最小值
type MedianFinder struct {
	LeftData  heap2.MaxHeap
	RightData heap2.MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	m := new(MedianFinder)
	m.LeftData = make(heap2.MaxHeap, 0)
	m.RightData = make(heap2.MinHeap, 0)
	return *m
}
func (this *MedianFinder) FindMedian() float64 {
	if len(this.RightData) > len(this.LeftData) {
		return float64(this.RightData[0])
	} else {
		return float64(this.LeftData[0]+this.RightData[0]) / 2.0
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 原则上右比左多一个，所以首先加到右边
	if len(this.RightData) == 0 || this.RightData[0] <= num {
		heap.Push(&this.RightData, num)
	} else {
		heap.Push(&this.LeftData, num)
	}

	// 调整
	if len(this.RightData)-len(this.LeftData) > 1 {
		v := heap.Pop(&this.RightData)
		heap.Push(&this.LeftData, v)
	} else if len(this.LeftData)-len(this.RightData) > 0 {
		v := heap.Pop(&this.LeftData)
		heap.Push(&this.RightData, v)
	}
}
