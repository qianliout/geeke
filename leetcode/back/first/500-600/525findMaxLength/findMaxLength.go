package main

import (
	"fmt"

	"qianliout/leetcode/leetcode/common/utils"
)

func main() {
	nums := []int{0, 0, 1, 0, 0, 1, 1}
	res := findMaxLength(nums)
	fmt.Println("res is ", res)
}

/*
给定一个二进制数组, 找到含有相同数量的 0 和 1 的最长连续子数组（的长度）。
示例 1:
输入: [0,1]
输出: 2
说明: [0, 1] 是具有相同数量0和1的最长连续子数组。
示例 2:
输入: [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
注意: 给定的二进制数组的长度不会超过50000。
*/

// fixme not right
func findMaxLength(nums []int) int {

	ans := 0
	dp := make(map[int]int)
	dp[0] = -1
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j+i] == 0 {
				dp[j] = dp[j] - 1
			} else {
				dp[j] = dp[j] + 1
			}

			if j-i+1 > ans {
				ans = j - i + 1
			}
		}
	}

	return ans
}

// 这道题用到的知识点就是前缀和，类似这种要找相同数量的题目其实都可以用到前缀和，用一个count变量来计算出现的总次数，
// 当出现1则++，出现0则--，相同count数之间的连续数组即为题意所指区间，用一个hashmap来保存count数的位置即可
func find(nums []int) int {
	ans := 0
	preSum := 0
	hash := make(map[int]int)
	hash[0] = -1
	for i, n := range nums {
		if n == 1 {
			preSum += 1
		} else {
			preSum += -1
		}
		idx, ok := hash[preSum]
		if ok {
			ans = utils.Max(ans, i-idx)
		} else {
			hash[preSum] = i
		}
	}

	return ans
}
