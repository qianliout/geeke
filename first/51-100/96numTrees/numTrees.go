package main

import (
	"fmt"
)

func main() {
	fmt.Println(numTrees(3))
}

/*
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
示例:
输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

// 有了95的理解就简单的多了，用dp,其实也是一个有顺序的01背包问题，因为有顺序，所以target在外层循环
// 内层循环是物品数，物品数就是1到i个物品
func numTrees(n int) int {
	dp := make([]int, n+3)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ { // right
		for j := 1; j <= i; j++ { // left
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
