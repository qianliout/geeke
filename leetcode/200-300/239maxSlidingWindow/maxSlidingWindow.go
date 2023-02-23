package main

import (
	"container/heap"
	"fmt"

	. "qianliout/leetcode/leetcode/common/commonHeap"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	window := maxSlidingWindow(nums, 3)
	fmt.Println("window is ", window)

}

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
*/

// 方案一,优先队列
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]*IntItem, 0)
	qu := PriorityQueue(res)
	for i := 0; i < k-1; i++ {
		heap.Push(&qu, &IntItem{Value: i, Priority: nums[i]})
	}

	ans := make([]int, 0)

	for i := k - 1; i < len(nums); i++ {
		heap.Push(&qu, &IntItem{Value: i, Priority: nums[i]})
		// 有k个值了，就得取值了
		for qu[0].Value <= i-k {
			heap.Pop(&qu)
		}
		ans = append(ans, qu[0].Priority)
	}
	return ans
}

// 单调stark
func maxSlidingWindow2(nums []int, k int) []int {
	stark, ans := make([]int, 0), make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// 移出多余数据
		for len(stark) > 0 && i-stark[0] >= k {
			stark = stark[1:]
		}
		// 维护stark,单调递减
		for len(stark) > 0 && nums[stark[len(stark)-1]] < nums[i] {
			stark = stark[0 : len(stark)-1]
		}
		stark = append(stark, i)
		// 取值
		if i >= k-1 {
			ans = append(ans, nums[stark[0]])
		}
	}
	return ans
}
