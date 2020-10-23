package main

import (
	"fmt"
	"math"

	. "outback/leetcode/common"
)

func main() {
	coins := []int{2}
	res := coinChange2(coins, 4)
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
	// sort.Ints(coins)

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
	// fmt.Println(dp)
	return dp[amount]
}

func coinChange2(coins []int, amount int) int {
	dp := make(map[int]int)
	// 这里要赋初值，因为，如果你不赋初值，那么所有的值都不会比0小，下面取最小值时就会出错
	// 这就是这里最不容易考虑的点
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32 // 表示不能
	}
	dp[0] = 0 // 这个也是特殊值
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] = Min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
