package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{0}
	res := lengthOfLIS(nums)
	fmt.Println("res is ", res)
}

/*
给定一个无序的整数数组，找到其中最长上升子序列的长度。
示例:
输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
说明:
    可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
    你算法的时间复杂度应该为 O(n2) 。
*/
// 注意，这里不要求连续
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make(map[int]int)
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		n := 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				n = maxInt(dp[j]+1, n)
			}
		}
		dp[i] = n

		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}

func maxInt(nums ...int) int {
	nu := math.MinInt64
	for _, num := range nums {
		if num > nu {
			nu = num
		}
	}
	return nu
}

