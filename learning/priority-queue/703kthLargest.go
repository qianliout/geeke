package priority

import (
	"container/heap"
	commonheap "outback/leetcode/common/heap"
	"sort"
)

/*
设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。
你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。
每次调用 KthLargest.add，返回当前数据流中第K大的元素。
示例:
int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
说明:
你可以假设 nums 的长度≥ k-1 且k ≥ 1。
*/

type KthLargest struct {
	MaxLent int
	//Numbers     []int
	MinHeap     *commonheap.IntMinHeap
	SortedSlice []int
}

func ConstructorKthLargest(k int, nums []int) KthLargest {
	sort.Ints(nums)
	if len(nums) <= k {
		return KthLargest{
			MaxLent:     k,
			MinHeap:     nil,
			SortedSlice: nums,
		}
	} else {
		return KthLargest{
			MaxLent:     k,
			MinHeap:     nil,
			SortedSlice: nums[len(nums)-k:],
		}
	}
}

func ConstructorKthLargestByHeap(k int, nums []int) *KthLargest {
	minHeap := commonheap.IntMinHeap{}
	for key, value := range nums {
		if key < k {
			minHeap.Push(value)
		} else {
			if value < minHeap.Peek().(int) {
				continue
			} else {
				minHeap.Push(value)
				minHeap.PopMini()
				heap.Init(&minHeap)
			}
		}
	}
	return &KthLargest{
		MaxLent:     k,
		MinHeap:     &minHeap,
		SortedSlice: nil,
	}

}

func (this *KthLargest) Add(val int) int {
	//return AddBySortedSlice(this, val)
	return AddByMinHeap(this, val)
}

func AddBySortedSlice(this *KthLargest, val int) int {
	if val >= this.SortedSlice[len(this.SortedSlice)-1] {
		this.SortedSlice = append(this.SortedSlice, val)
		this.SortedSlice = this.SortedSlice[1:]
		return this.SortedSlice[0]
	}
	if val <= this.SortedSlice[0] {
		return this.SortedSlice[0]
	}
	this.SortedSlice = append(this.SortedSlice, val)
	sort.Ints(this.SortedSlice)
	this.SortedSlice = this.SortedSlice[1:]
	return this.SortedSlice[0]
}
func AddByMinHeap(this *KthLargest, val int) int {
	preMin := this.MinHeap.Peek().(int)
	if val <= preMin {
		return preMin
	} else {
		this.MinHeap.Push(val)
		heap.Init(this.MinHeap)
		this.MinHeap.PopMini()
		return this.MinHeap.Peek().(int)
	}
}
