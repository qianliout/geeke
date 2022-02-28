package main

import (
	"fmt"
	"math"

	. "qianliout/leetcode/common/utils"
)

func main() {
	nums := []int{2, 1}
	jm := jump(nums)
	fmt.Println("jump is ", jm)
}

/*
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例 2:
输入: nums = [2,3,0,1,4]
输出: 2
*/

func jump2(nums []int) int {
	// dp[i]表示到i最少需要多少步
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		dp[i] = math.MaxInt64
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				dp[i] = Min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[len(nums)-1]
}

// 贪心算法
func jump3(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res, end, maxPos := 0, 0, 0
	// 这里的难点是len(nums)-1，因为，如果i符合条件，就会res++,如果i=len(nums)-1,就会出界
	// 题目已经说明了，保证可以跳到最后一个位置
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > maxPos {
			maxPos = i + nums[i]
		}
		if i == end {
			end = maxPos
			res++
		}
	}
	return res
}

func jump(nums []int) int {
	res, end, maxPos := 0, 0, 0
	// res 表示跳跃的步数
	// end 表示上一次跳跃时的最后位置
	// 表示到i这个位置时，能跳跃到的最大位置
	// 题目中已经说明确保可以跳跃到最后，所以当i走到最后时，res++的话就会超过，所以i只能走到倒数第二
	for i := 0; i < len(nums)-1; i++ {
		maxPos = Max(maxPos, i+nums[i])
		if end == i {
			end = maxPos
			res++
		}
	}
	return res
}
