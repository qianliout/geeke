package main

import (
	"fmt"
)

func main() {
	res := climbStairs(20)
	fmt.Println("res is ", res)
}

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。
示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
链接：https://leetcode-cn.com/problems/climbing-stairs
*/
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	dp := make(map[int]int)
	dp[0], dp[1], dp[2] = 1, 2, 3
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
