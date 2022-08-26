package main

import (
	. "qianliout/leetcode/common/utils"
)

func main() {

}

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	buy1, sell1, buy2, sell2 := -prices[0], 0, -prices[0], 0

	for i := 1; i < len(prices); i++ {
		buy1 = Max(buy1, -prices[i])
		sell1 = Max(sell1, buy1+prices[i])
		buy2 = Max(buy2, sell1-prices[i])
		sell2 = Max(sell2, buy2+prices[i])
	}
	return sell2
}
