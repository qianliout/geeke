package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	nums := []int{3, 1, 5, 8}
	coins := maxCoins(nums)
	fmt.Println("coins is ", coins)
}

func maxCoins(nums []int) int {
	// 前后加1
	nums = append([]int{1}, append(nums, 1)...)
	// dp[i][j]表示i到j开区间的值，不包括i和j
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}
	// 初值都是0，可以不显式的指定
	// 这里i,j的循环方向，没有明白
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = Max(dp[i][j], dp[i][k]+dp[k][j]+nums[k]*nums[i]*nums[j])
			}
		}
	}

	return dp[0][len(nums)-1]
}
