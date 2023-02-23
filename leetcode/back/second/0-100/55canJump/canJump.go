package main

import (
	"fmt"

	"qianliout/leetcode/leetcode/common/utils"
)

func main() {
	num := []int{2, 0, 0}
	res := canJump(num)
	fmt.Println("res is ", res)
}

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。
示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
示例 2:
输入: [3,2,1,0,4]
输出: false
解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/
func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	// 指针指向i时，dp这个小人走到的位置，如如小i,说明已不能走到i了，
	dp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if dp[i-1] < i {
				return false
			}
			dp[i] = utils.Max(nums[i]+i, dp[i-1])
		} else {
			dp[i] = nums[i] + i
		}
	}

	return true
}

func canJump2(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	step := 0 // 表示还能前进多少步
	for i := 0; i < len(nums)-1; i++ {
		step = utils.Max(step-1, nums[i])
		if step == 0 {
			return false
		}
	}

	return true
}
