package main

import (
	"fmt"

	"qianliout/leetcode/back/common"
)

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

/*
给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。
你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。
返回获得利润的最大值。
注意：这里的一笔交易指买入持有并卖出股票的整个过程，每笔交易你只需要为支付一次手续费。
示例 1:
输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
输出: 8
解释: 能够达到的最大利润:
在此处买入 prices[0] = 1
在此处卖出 prices[3] = 8
在此处买入 prices[4] = 4
在此处卖出 prices[5] = 9
总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
*/
func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	// dp1[i]表示持有该股票的
	dp1 := make([]int, len(prices))
	// 不持有
	dp2 := make([]int, len(prices))
	// 初值
	dp1[0] = -prices[0] - fee
	dp2[0] = 0

	for i := 1; i < len(prices); i++ {
		dp1[i] = common.Max(dp1[i-1], dp2[i-1]-prices[i]-fee)
		dp2[i] = common.Max(dp2[i-1], dp1[i-1]+prices[i])
	}
	return dp2[len(prices)-1]
}
