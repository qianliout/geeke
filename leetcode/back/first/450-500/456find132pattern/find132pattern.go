package main

import ()

func main() {

}

/*
给定一个整数序列：a1, a2, ..., an，一个132模式的子序列 ai, aj, ak 被定义为：当 i < j < k 时，ai < ak < aj。
设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。
注意：n 的值小于15000。
示例1:
输入: [1, 2, 3, 4]
输出: False
解释: 序列中不存在132模式的子序列。
示例 2:
输入: [3, 1, 4, 2]
输出: True
解释: 序列中有 1 个132模式的子序列： [1, 4, 2].
示例 3:
输入: [-1, 3, 2, 0]
输出: True
解释: 序列中有 3 个132模式的的子序列: [-1, 3, 2], [-1, 3, 0] 和 [-1, 2, 0].
*/

// 解法在：https://leetcode-cn.com/problems/132-pattern/solution/zhan-jie-fa-chao-xiang-xi-ti-jie-by-siyy/
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	// 先维护前面
	length := len(nums)
	min := nums[0]
	minStark := make([]int, length)
	for i := 0; i < length; i++ {
		if nums[i] < min {
			min = nums[i]
		}
		minStark[i] = min
	}

	// 再维护后面
	stark := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		if nums[i] > minStark[i] {
			for len(stark) > 0 && stark[len(stark)-1] <= minStark[i] {
				stark = stark[:len(stark)-1]
			}
			if len(stark) > 0 && stark[len(stark)-1] < nums[i] {
				return true
			}
			stark = append(stark, nums[i])
		}
	}
	return false
}
