package main

import (
	"fmt"
	"sort"

	"outback/leetcode/common/heap"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow(nums, k)
	fmt.Println("res is ", res)
}

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。
示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]
解释:
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
*/

// 大顶推的操作(因为自己的大顶堆有问题，所以使用小顶堆来实现)
func maxSlidingWindow(nums []int, k int) []int {

	res := make([]int, 0)
	if len(nums) <= 0 {
		return res
	}

	preMinHaap := make(heap.IntMinHeap, 0)
	for i, va := range nums {
		if i < k {
			preMinHaap.Push(va)
		} else {
			preMinHaap.Pop() // 先移出最后的那个
			preMinHaap.Push(va)
		}
		minHeap := make(heap.IntMinHeap, k)
		copy(minHeap, preMinHaap)
		fmt.Println("1,", minHeap)
		if i >= k-1 {
			heap.InitMin(&minHeap)
			res = append(res, minHeap.PeekLast().(int))
		}
	}
	return res
}

// 使用双端队列的方法

// 使用维护一个已排序的方法，因为这里要用到深复制，所以要超出时间限制，小数据量可以使用
func maxSlidingWindow2(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	dqueue := make([]int, 0)
	for i, va := range nums {
		if i >= k {
			dqueue = dqueue[1:]
		}
		dqueue = append(dqueue, va)

		if i >= k-1 {
			kq := make([]int, k)
			copy(kq, dqueue)
			sort.Ints(kq)
			res = append(res, kq[len(kq)-1])
		}
	}
	return res
}
func maxSlidingWindow3(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	dq := make([]int, 0) // 这里存下标,这样可以可以直接到nums中根据下标找到相对应的值，
	for i, va := range nums {
		// 第一步，把不在窗口范围内的值移出
		if i >= k && len(dq) > 0 && i-dq[0] >= k {
			dq = dq[1:]
		}
		// 第二步，把新加的值放到指定的位置
		for len(dq) > 0 && nums[dq[len(dq)-1]] <= va {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		// 第三步，把结果返回
		if i >= k-1 {
			res = append(res, nums[dq[0]])
		}
	}
	return res
}
