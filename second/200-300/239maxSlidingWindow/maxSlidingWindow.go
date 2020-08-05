package main

import (
	"container/heap"
	"fmt"

	. "outback/leetcode/common/commonHeap"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	res := maxSlidingWindow2(nums, 3)
	fmt.Println("res is ", res)

}

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。
进阶：
你能在线性时间复杂度内解决此题吗？
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
示：
    1 <= nums.length <= 10^5
    -10^4 <= nums[i] <= 10^4
    1 <= k <= nums.length
*/

// 单调栈
func maxSlidingWindow(nums []int, k int) []int {
	return MaxHeap{}
}

// 使用堆
func maxSlidingWindow2(nums []int, k int) []int {
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
