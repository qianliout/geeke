package main

import (
	"fmt"
	"math"
)

func main() {
	nums := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := minPathSum(nums)
	fmt.Println("res is ", res)
}

/*
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	dp := make(map[int]map[int]int)
	// 初始化map
	for i := 0; i < len(grid); i++ {
		dp[i] = make(map[int]int)
	}
	// 初值
	dp[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < len(grid[0]); i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	// 状态转移方程
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = MinInt(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func MinInt(nums ...int) int {
	ans := math.MaxInt64
	for _, num := range nums {
		if num < ans {
			ans = num
		}
	}
	return ans
}
