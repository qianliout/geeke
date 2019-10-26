package priority

import (
	"fmt"
	"sort"
)

/*
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
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
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小
*/

func MaxSlidingWindow(nums []int, k int) []int {
	//return MaxSlidingWindowBySortSlice(nums, k)
	return MaxSlidingWindowByDqueue(nums, k)
}

//维护已排序的队列,这种方法不可取，因为要用到排序，并且还用到深复制，时间，空间都不好
func MaxSlidingWindowBySortSlice(nums []int, k int) []int {
	if k <= 0 || len(nums) <= 0 {
		return nil
	}
	result := make([]int, 0)
	for i := 0; i <= len(nums)-k; i++ {
		so := make([]int, k)
		copy(so, nums[i:i+k])
		sort.Ints(so)
		result = append(result, so[k-1])
	}
	return result
}

// 使用双端队列
func MaxSlidingWindowByDqueue(nums []int, k int) []int {
	if len(nums) == 0 || k <= 1 {
		return nums
	}

	result := make([]int, 0)
	length := len(nums)
	dq := make([]int, 0) // 注意这里存是队列的下标，这里是关键
	for i := 0; i < length; i++ {
		fmt.Println("i:", i)
		for len(dq) > 0 && nums[dq[len(dq)-1]] < nums[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if len(dq) > 0 && i-dq[0] >= k {
			dq = dq[1:]
		}
		if i >= k-1 {
			result = append(result, dq[0])
		}
	}

	return result
}
