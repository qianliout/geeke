package main

import (
	"fmt"

	common2 "qianliout/leetcode/common"
)

func main() {
	nums := []int{2, 3, 0, 4, -2, -3}
	res := maxProduct(nums)
	fmt.Println("res is ", res)
}

/*
 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
 示例 1:
 输入: [2,3,-2,4]
 输出: 6
 解释: 子数组 [2,3] 有最大乘积 6。
 示例 2:
 输入: [-2,0,-1]
 输出: 0
 解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
*/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxDp := make(map[int]int)
	minDp := make(map[int]int)
	maxDp[0], minDp[0] = nums[0], nums[0]
	res := maxDp[0]
	for i := 1; i < len(nums); i++ {
		maxDp[i] = common2.Max(nums[i], nums[i]*maxDp[i-1], nums[i]*minDp[i-1])
		if maxDp[i] > res {
			res = maxDp[i]
		}
		minDp[i] = common2.Min(nums[i], nums[i]*maxDp[i-1], nums[i]*minDp[i-1])
	}
	return res
}
