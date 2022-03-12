package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	nums := []int{2, 0, 1, 1, 0}
	jump := canJump(nums)
	fmt.Println("jump ", jump)
}

/*
给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标。
示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标
*/
func canJump(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums); i++ {
		maxPos = Max(i+nums[i], maxPos)
		if maxPos <= i && i < len(nums)-1 {
			return false
		}
	}
	return true
}
