package main

import (
	"fmt"
)

func main() {
	coins := []int{186, 419, 83, 408}
	res := coinChange(coins, 6249)
	fmt.Println("res is ", res)
}

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回 -1。
示例 1:
输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:
输入: coins = [2], amount = 3
输出: -1
说明:
你可以认为每种硬币的数量是无限的。
*/

func coinChange(coins []int, amount int) int {
	dp := make(map[int]int)
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	dp[0] = 0
	//sort.Ints(coins)

	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, v := range coins {
			if i < v || dp[i-v] == -1 {
				continue
			}
			count := dp[i-v] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	//fmt.Println(dp)
	return dp[amount]
}
