package main

import (
	"math"

	. "outback/leetcode/common"
)

func main() {

}

func jump(nums []int) int {
	dp := make(map[int]int)
	dp[0], dp[1] = 0, 1
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		min := math.MaxInt64
		for j := i - 1; j >= 0; j-- {
			if j+nums[j] >= i && dp[j] < min {
				min = dp[j]
			}
		}
		dp[i] = min + 1
	}
	return dp[len(nums)-1]
}

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。
示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/
func canJump2(nums []int) bool {
	dp := make(map[int]bool)
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

func canJump(nums []int) bool {
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		step = Max(step-1, nums[i])
		if step == 0 {
			return false
		}
	}
	return true
}
