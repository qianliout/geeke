package main

import (
	"fmt"
	"math"
)

func main() {
	price := []int{3, 3}
	res := maxProfit(price)
	fmt.Println("res is ", res)
}

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
示例 1:
输入: [3,3,5,0,0,3,1,4]
输出: 6
解释: 在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
     随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
示例 2:
输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3:
输入: [7,6,4,3,1]
输出: 0
解释: 在这个情况下, 没有交易完成, 所以最大利润为 0。
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// dpNo[i][k] 第i天,最多允许k次交易，不持有股票
	// dpHave[i][k] 第i天，最多允许k次交易，持有股票
	k := 1
	dpNo := make(map[int]map[int]int)
	dpHave := make(map[int]map[int]int)
	// 初始化
	for i := -1; i < len(prices); i++ {
		dpNo[i] = make(map[int]int)
		dpHave[i] = make(map[int]int)
	}
	// 初值
	for i := 0; i < len(prices); i++ {
		dpNo[i][0] = 0               // 不持有股票，不交易，利润为0
		dpHave[i][0] = math.MinInt64 // 持有股票，但是不交易是不可能的，为负无穷表示
	}
	dpHave[0][1] = -prices[0]

	for i := 0; i < len(prices); i++ {
		if i-1 == -1 {
			dpNo[i][k] = 0
			dpHave[i][k] = math.MinInt64
		}
		for j := 1; j <= k; j++ {
			//  第i天不持有股票，并可以交易j次
			dpNo[i][j] = Max(dpNo[i-1][j], dpHave[i-1][j]+prices[i])
			//  第i天持有股票，并可以交易j次
			dpHave[i][j] = Max(dpHave[i-1][j], dpNo[i-1][j-1]-prices[i])
		}
	}
	return dpNo[len(prices)-1][k]
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
