package main

import (
	"fmt"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	// prices := []int{3, 2, 6, 5, 0, 3}
	// prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	prices := []int{2, 6, 8, 7, 8, 7, 9, 4, 1, 2, 4, 5, 8}
	// prices := []int{1, 2}
	profit := maxProfit(100, prices)
	fmt.Println(profit)

}

/*
给定一个整数数组 prices ，它的第 i 个元素 prices[i] 是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/
func maxProfit(k int, prices []int) int {
	if len(prices) < 2 || k <= 0 {
		return 0
	}

	k = k % len(prices)
	if k == 0 {
		k = len(prices)
	}

	// 第i天，已经交易j次
	hold, notHold := make([][]int, len(prices)), make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		hold[i], notHold[i] = make([]int, k+1), make([]int, k+1)
	}

	// 初值,j从1开始是一个技巧
	for j := 1; j <= k; j++ {
		// 在第一天，不管交易多少次都是
		hold[0][j] = -prices[0]
	}

	// 循环
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			notHold[i][j] = Max(notHold[i-1][j], hold[i-1][j]+prices[i])
			hold[i][j] = Max(hold[i-1][j], notHold[i-1][j-1]-prices[i])
		}
	}

	return Max(notHold[len(prices)-1]...)
}

func maxProfitGreedy(prices []int) int {
	res := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			res = res + prices[i+1] - prices[i]
		}
	}
	return res
}
