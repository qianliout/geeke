package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {
	price := []int{1, 2, 3, 0, 2}
	res := maxProfit(price)
	fmt.Println("res is ", res)
}

/*
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
示例:
输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
*/

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buyDp := make(map[int]int)  // 表示当天有股票
	saleDp := make(map[int]int) // 表示当天无股票
	// 初值
	length := len(prices)
	buyDp[0] = -prices[0]
	dp3 := make(map[int]int)

	for i := 1; i < length; i++ {
		// 今天持股,昨天持股,今天继续,昨天冷冻期,今天买入
		buyDp[i] = common.Max(dp3[i-1]-prices[i], buyDp[i-1])
		// 今天不持股,昨天不持股,今天继续,昨天持股,今天买出
		saleDp[i] = common.Max(saleDp[i-1], buyDp[i-1]+prices[i])
		// 今天冷冻期
		dp3[i] = common.Max(saleDp[i-1], dp3[i-1])
	}
	return saleDp[length-1]
}
