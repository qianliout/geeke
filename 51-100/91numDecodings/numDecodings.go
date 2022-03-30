package main

import (
	"fmt"
)

func main() {
	nms := numDecodings("28")
	fmt.Println("mns is ", nms)
}

/* 一条包含字母 A-Z 的消息通过以下映射进行了 编码 ： 'A' -> "1" 'B' -> "2"
...
'Z' -> "26"
要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
"AAJF" ，将消息分组为 (1 1 10 6)
"KJF" ，将消息分组为 (11 10 6)
注意，消息不能分组为  (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
题目数据保证答案肯定是一个 32 位 的整数。
*/
func numDecodings(s string) int {
	nums := make([]int, len(s))
	for i, ch := range s {
		nums[i] = int(ch) - 48
	}
	return decode(nums)
}

func decode(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp[i]表示，nums[i-1]时解码数量
	dp := make([]int, len(nums)+1)
	if nums[0] == 0 {
		return 0
	}
	// 初值是这一道题目的关键,但还是没有能理解
	dp[0], dp[1] = 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			if nums[i-1] == 1 || nums[i-1] == 2 {
				dp[i+1] = dp[i-1]
			} else {
				return 0
			}
		} else if nums[i-1] == 1 || (nums[i-1] == 2 && nums[i] >= 1 && nums[i] <= 6) {
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}

	return dp[len(nums)]
}
