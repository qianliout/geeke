package main

import (
	"container/heap"
	"fmt"

	. "qianliout/leetcode/leetcode/common/commonHeap"
)

func main() {
	// nums := []int{3, 2, 1, 5, 6, 4}
	nums := []int{-1, 2, 0}
	largest := findKthLargest(nums, 2)
	fmt.Println("largest:", largest)
}

/*
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
*/
func findKthLargest(nums []int, k int) int {
	minHeap := make(MinHeap, 0)
	for i := range nums {
		if i < k {
			heap.Push(&minHeap, nums[i])
			continue
		}
		top := minHeap[0]
		if nums[i] < top {
			continue
		}
		heap.Pop(&minHeap)
		heap.Push(&minHeap, nums[i])
	}
	return minHeap[0]
}
