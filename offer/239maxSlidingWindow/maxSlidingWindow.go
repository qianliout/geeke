package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ans := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7, 1, 1, 1, 1, 1 - 1}, 2)
	ans2 := maxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7, 1, 1, 1, 1, 1 - 1}, 2)
	fmt.Println("ans is ", ans, ans2)
}

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。
*/

// 方法一，大顶堆，延迟删除
func maxSlidingWindow(nums []int, k int) []int {
	maxHeap := make(MaxHeap, 0)
	deletMap := make(map[int]int)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			heap.Push(&maxHeap, nums[i])
		} else {
			heap.Push(&maxHeap, nums[i])
			// fmt.Println(maxHeap)
			for {
				val := maxHeap[0]
				if deletMap[val] == 0 {
					ans = append(ans, val)
					break
				} else {
					heap.Pop(&maxHeap)
					deletMap[val]--
				}
			}

			deletMap[nums[i-k+1]]++
		}

	}
	return ans
}

// 用单调栈的方式,递减栈，最大值放在第一个
func maxSlidingWindow2(nums []int, k int) []int {
	stark := make([]int, 0)
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for i >= k && len(stark) > 0 && i-stark[0] >= k {
			stark = stark[1:]
		}
		for len(stark) > 0 && nums[stark[len(stark)-1]] <= nums[i] {
			stark = stark[:len(stark)-1]
		}
		stark = append(stark, i)
		if i >= k-1 {
			ans = append(ans, nums[stark[0]])
		}
	}

	return ans
}

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]

}

func (m *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}
