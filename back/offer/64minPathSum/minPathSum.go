package main

import (
	"outback/leetcode/back/common"
)

func main() {

}

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
*/

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dp := make(map[int]map[int]int)

	// 初始化加初值
	for i := 0; i < len(grid); i++ {
		dp[i] = make(map[int]int)
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 0; j < len(grid[0]); j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = common.Max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}
