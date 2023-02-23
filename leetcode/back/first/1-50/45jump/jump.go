package main

import (
	"fmt"
	"math"
)

func main() {
	// nums := []int{1, 3, 2}
	// nums := []int{2, 3, 1, 4, 4}
	nums := []int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}
	res := jump2(nums)
	fmt.Println("res is ", res)
}

/*
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
示例:
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
链接：https://leetcode-cn.com/problems/jump-game-ii
*/
// 方法一，dp,方法是正确的，但是会超时
func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	// dp[i]表示走到下标为i的这个元素的最小步数
	dp := make(map[int]int)
	dp[0], dp[1] = 0, 1

	for i := 1; i < length; i++ {
		min := math.MaxInt64
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] && dp[j] < min {
				min = dp[j]
			}
		}
		dp[i] = min + 1
	}
	return dp[length-1]
}

// 直接跳 （贪心算法，也是双指针法）
func jump2(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res, end, maxpos := 0, 0, 0
	// 这里的难点是len(nums)-1，因为，如果i符合条件，就会res++,如果i=len(nums)-1,就会出界
	for i := 0; i < len(nums)-1; i++ {

		if i+nums[i] > maxpos {
			maxpos = i + nums[i]
		}
		if i == end {
			end = maxpos
			res++
		}
	}

	return res
}
