package main

import (
	"container/heap"
	"fmt"
	"sort"

	. "outback/leetcode/common/commonHeap"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindow4(nums, k)
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
	deleteMap := make(map[int]int)
	maxHeap := make(MaxHeap, 0)
	res := make([]int, 0)
	i := 0
	for i < len(nums) {
		heap.Push(&maxHeap, nums[i])
		i++
		if i >= k {
			for {
				top := maxHeap[0]
				if deleteMap[top] <= 0 {
					res = append(res, top)
					break
				} else {
					p := heap.Pop(&maxHeap).(int)
					deleteMap[p]--
				}
			}
			deleteMap[nums[i-k]]++
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

// 单调栈的思想
func maxSlidingWindow4(nums []int, k int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	// 注意这里的单调整栈存放的是下标
	stack := make([]int, 0)

	for i, n := range nums {
		// 第一步,把超出范围的数据移出(注意这里有等于,因为当等于是,说明刚好有k个元素,再加入i这个元素就多了一个,所以也要移出)
		if i >= k && len(stack) > 0 && i-stack[0] >= k {
			stack = stack[1:]
		}
		// 为了把这个数加到对应的位(保持单掉递减),把多余的数据去除
		for len(stack) > 0 && nums[stack[len(stack)-1]] < n {
			stack = stack[:len(stack)-1]
		}
		// 加入值
		stack = append(stack, i)
		// 输出值
		if i >= k-1 {
			res = append(res, nums[stack[0]])
		}
	}
	return res
}
