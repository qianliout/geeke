package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {
	res := numSquares(12)
	fmt.Println("res is ", res)
}

/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
示例 1:
输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:
输入: n = 13
输出: 2
解释: 13 = 4 + 9.
*/
// dp[i] = MIN(dp[i], dp[i - j * j] + 1)
func numSquares(n int) int {
	dp := make(map[int]int)
	dp[0], dp[1], dp[2], dp[3], dp[4] = 0, 1, 2, 3, 1
	if n <= 1 {
		return dp[n]
	}
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = common.Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func numSquares2(n int) int {
	dp := make(map[int]int)
	dp[1] = 1
	if n <= 1 {
		return dp[n]
	}

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = common.Min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
