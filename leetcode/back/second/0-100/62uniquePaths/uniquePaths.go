package main

import (
	"fmt"
)

func main() {
	res := uniquePaths(3, 2)
	fmt.Println(res)
}

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
问总共有多少条不同的路径？
*/

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	dp := make(map[int]map[int]int)
	// 初值
	for i := 0; i < m; i++ {
		if dp[i] == nil {
			dp[i] = make(map[int]int)
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
