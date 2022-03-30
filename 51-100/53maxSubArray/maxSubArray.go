package main

import (
	. "qianliout/leetcode/common/utils"
)

func main() {

}

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
*/
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i])
		res = Max(dp[i], res)
	}
	return res
}

func maxSubArray2(nums []int) int {
	res, curMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		curMax = Max(curMax+nums[i], nums[i])
		res = Max(curMax, res)
	}
	return res
}
