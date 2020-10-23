package main

import (
	"fmt"
)

func main() {
	coins := []int{1, 2, 5}
	res := change(5, coins)
	fmt.Println(res)
}

/*
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。
示例 1:
输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
*/
func change(amount int, coins []int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}
	dp := make(map[int]int)
	dp[0] = 1
	for _, c := range coins {
		for i := 1; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}

	return dp[amount]
}

// 完全背包问题
func change2(amount int, coins []int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
