package main

import (
	"fmt"
)

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}

/*
给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。
我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
示例 1：
输入：
nums = [1, 7, 3, 6, 5, 6]
输出：3
解释：
索引 3 (nums[3] = 6) 的左侧数之和 (1 + 7 + 3 = 11)，与右侧数之和 (5 + 6 = 11) 相等。
同时, 3 也是第一个符合要求的中心索引。
示例 2：
输入：
nums = [1, 2, 3]
输出：-1
解释：
数组中不存在满足此条件的中心索引。
*/
// 暴力法
// 这里有注意的是，左侧和右侧的元素也都可能为空
func pivotIndex(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	left := 0

	for i := 0; i < len(nums); i++ {
		if i-1 >= 0 {
			left += nums[i-1]
		}
		right := sum - nums[i] - left
		if left == right {
			return i
		}
	}
	return -1
}
