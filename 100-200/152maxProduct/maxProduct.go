package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/common/utils"
)

func main() {
	nums := []int{2, 3, -2, 4}
	ans := maxProduct(nums)
	fmt.Println("ans is ", ans)
}

/*
给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
*/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	maxDp := make([]int, len(nums))
	minDp := make([]int, len(nums))
	maxDp[0], minDp[0] = nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			maxDp[i] = Max(maxDp[i-1]*nums[i], nums[i])
			minDp[i] = Min(minDp[i-1]*nums[i], nums[i])
		} else if nums[i] < 0 {
			maxDp[i] = Max(minDp[i-1]*nums[i], nums[i])
			minDp[i] = Min(maxDp[i-1]*nums[i], nums[i])
		} else if nums[i] == 0 {
			maxDp[i] = 0
			minDp[i] = 0
		}
	}
	ans := math.MinInt

	for i := range maxDp {
		ans = Max(ans, maxDp[i])
	}

	return ans
}
