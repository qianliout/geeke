package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max := maxSubArray2(nums)
	fmt.Println("max is ", max)
}

// 因为是连续，所以可以使用贪心算法
func maxSubArray(nums []int) int {
	currSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = int(math.Max(float64(nums[i]), float64(currSum+nums[i])))
		maxSum = int(math.Max(float64(maxSum), float64(currSum)))
	}
	return maxSum
}

// 使用动态规划的方法
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+nums[i]), float64(nums[i])))
		result = int(math.Max(float64(result), float64(dp[i])))
	}
	return result
}
