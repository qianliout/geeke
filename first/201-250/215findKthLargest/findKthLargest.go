package main

import (
	"container/heap"
	"fmt"
	"sort"
	. "outback/leetcode/common/commonHeap"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//nums := []int{2, 1}
	res := findKthLargest3(nums, 2)
	fmt.Println("res is ", res)

}

/*
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
示例 1:
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
示例 2:
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
*/

// 方法一，排序
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return 0
	}
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 方法二，小顶堆
func findKthLargest2(nums []int, k int) int {
	minHeap := make(MinHeap, 0)
	for _, num := range nums {
		if len(minHeap) < k {
			heap.Push(&minHeap, num)
		} else {
			if num > minHeap[0] {
				heap.Pop(&minHeap)
				heap.Push(&minHeap, num)
			}
		}
	}
	return minHeap[0]
}

func findKthLargest3(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return 0
	}

	k = len(nums) - k // 转化为第k小的问题
	low, hight := 0, len(nums)-1
	for {
		n := partion(&nums, low, hight)
		if k == n {
			return nums[k]
		} else if n < k {
			low = n + 1

		} else if n > k {
			hight = n - 1
		}
	}
}

func partion(nums *[]int, low, hight int) int {
	i := low
	pivot := (*nums)[hight] // 这里可以随机选一个点做为pivot，会在极端的情况下加快速度（比如，已经是排好序的）
	for j := i; j < hight; j++ {
		if (*nums)[j] < pivot {
			(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
			i++
		}
	}
	(*nums)[i], (*nums)[hight] = (*nums)[hight], (*nums)[i]
	return i
}
