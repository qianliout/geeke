package main

import (
	"fmt"

	"outback/leetcode/common/heap"
)

func main() {
	num := []int{5, 3, 1, 1, 1, 3, 73, 1}
	res := topKFrequent(num, 2)
	fmt.Println("res is ", res)
}

/*
347. 前 K 个高频元素
给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:
输入: nums = [1], k = 1
输出: [1]
说明：
    你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
    你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
*/
func topKFrequent(nums []int, k int) []int {
	ans := make([]int, 0)
	if len(nums) == 0 || len(nums) < k {
		return ans
	}
	m := make(map[int]int)
	for _, v := range nums {
		m[v] += 1
	}
	fmt.Println("map is ", m)
	// 再使用小顶堆把第k大的值找出来(其实是优先队列的用法)
	minHeap := heap.IntMinHeap{}
	for n, v := range m {
		if len(minHeap) < k {
			minHeap.Push(n)
		} else {
			if v > m[minHeap.Peek().(int)] {
				minHeap.PopMini()
				minHeap.Push(n)
			}
		}
		heap.InitMin(&minHeap) // 这里会出错，因为他是以数值进行重排的,如果想有用的话就要重写Less方法
		fmt.Println("heap is ", minHeap)
	}
	for len(minHeap) > 0 {
		ans = append(ans, minHeap.PopMini().(int))
	}
	return ans
}
