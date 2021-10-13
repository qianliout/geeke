package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {
	res := integerBreak(4)
	fmt.Println("res is ", res)
}

/*
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
示例 1:
输入: 2
输出: 1
解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2:
输入: 10
输出: 36
解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
*/
// 典型的dp
func integerBreak(n int) int {
	dp := make(map[int]int)
	if n <= 3 {
		return n - 1
	}
	dp[0], dp[1], dp[2], dp[3] = 0, 1, 2, 3

	for i := 4; i <= n; i++ {
		dp[i] = i - 1 // 最小是 1*(i-1)
		for j := 1; i-j > 0; j++ {
			dp[i] = common.Max(dp[i], dp[i-j]*dp[j])
		}
	}
	return dp[n]
}
