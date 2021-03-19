package main

import (
	"fmt"

	. "outback/leetcode/common"
)

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

/*
   给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
   示例 1：
   输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
   输出：6
   解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
*/
func maxSubArray(nums []int) int {
	dp := make(map[int]map[int]int)
	for i, v := range nums {
		dp[i] = make(map[int]int)
		dp[i][i] = v
	}
	ans := dp[0][0]
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			dp[i][j] = Max(
				dp[i][j+1]+nums[j],
				dp[i-1][j]+nums[i],
				dp[i-1][j+1]+nums[i]+nums[j],
			)

			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

func maxSubArray2(nums []int) int {
	currSum, maxSum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = Max(nums[i], currSum+nums[i])
		maxSum = Max(currSum, maxSum)
	}
	return maxSum
}
func maxSubArray3(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = Max(dp[i-1]+nums[i], nums[i])
		result = Max(dp[i], result)
	}
	return result
}
