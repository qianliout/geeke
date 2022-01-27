package main

import (
	"fmt"
	"math"

	"qianliout/leetcode/back/common"
)

func main() {
	res := getMoneyAmount(10)
	fmt.Println("res is ", res)
}

/*
我们正在玩一个猜数游戏，游戏规则如下：
我从 1 到 n 之间选择一个数字，你来猜我选了哪个数字。
每次你猜错了，我都会告诉你，我选的数字比你的大了或者小了。
然而，当你猜了数字 x 并且猜错了的时候，你需要支付金额为 x 的现金。直到你猜到我选的数字，你才算赢得了这个游戏。
示例:
n = 10, 我选择了8.
第一轮: 你猜我选择的数字是5，我会告诉你，我的数字更大一些，然后你需要支付5块。
第二轮: 你猜是7，我告诉你，我的数字更大一些，你支付7块。
第三轮: 你猜是9，我告诉你，我的数字更小一些，你支付9块。
游戏结束。8 就是我选的数字。
你最终要支付 5 + 7 + 9 = 21 块钱。
*/
// https://leetcode-cn.com/problems/guess-number-higher-or-lower-ii/solution/dong-tai-gui-hua-c-you-tu-jie-by-zhang-xiao-tong-2/
func getMoneyAmount(n int) int {
	if n == 0 {
		return 0
	}
	// 定义dp
	dp := make(map[int]map[int]int)

	// 初始化
	for i := 0; i <= n; i++ {
		if dp[i] == nil {
			dp[i] = make(map[int]int)
		}
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt64
			if i == j {
				dp[i][j] = 0
			}
		}
	}
	// 计算值,先按例 (注意，所有的都是从1开始)
	for j := 2; j <= n; j++ { // 计算后一个数
		for i := j - 1; i >= 1; i-- {
			for k := i + 1; k <= j-1; k++ {
				dp[i][j] = common.Min(dp[i][j], common.Max(dp[i][k-1], dp[k+1][j])+k)
			}
			dp[i][j] = common.Min(dp[i+1][j]+i, dp[i][j], dp[i][j-1]+j)
		}
	}

	return dp[1][n]
}

func getMoneyAmount2(n int) int {

}
