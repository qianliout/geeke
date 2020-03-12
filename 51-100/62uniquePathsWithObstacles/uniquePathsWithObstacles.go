package main

import (
	"fmt"
)

func main() {

	res := uniquePaths(3, 2)
	fmt.Println("res is ", res)
}

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径
*/

func uniquePaths(m int, n int) int {
	dp := make(map[int]map[int]int, m)

	for i := 0; i <= m; i++ {
		dp[i] = make(map[int]int, 2)
	}
	// 状态转移方程
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1 // 定义初值，这里是这道题的关键
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
