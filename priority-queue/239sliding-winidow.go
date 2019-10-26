package priority

import "sort"

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
提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

进阶：
你能在线性时间复杂度内解决此题吗？
*/

func MaxSlidingWindow(nums []int, k int) []int {
	return MaxSlidingWindowBySortedSlice(nums, k)
}

// 方法一，维护一个已排序的数组
func MaxSlidingWindowBySortedSlice(nums []int, k int) []int {
	maxList := make([]int, 0)
	result := make([]int, 0)
	for key, value := range nums {
		if key < k-1 {
			maxList = append(maxList, value)
		} else {
			maxList = append(maxList, value)
			sort.Ints(maxList)
			if len(maxList) > k {
				maxList = maxList[1:]
			}
			result = append(result, maxList[k-1])
		}
	}
	return result
}

// 方法二，使用双端队列，在go语言中可以使用slice实现
func MaxSlidingWindowByQueue(nums []int, k int) []int {
}
