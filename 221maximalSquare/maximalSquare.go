package main

import (
	. "qianliout/leetcode/common/utils"
)

func main() {

}

func maximalSquare(matrix [][]byte) int {
	ans := 0
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ans
	}
	// dp[i][j]表示，下标为i,j选重时的最大边长
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[i]))
	}

	// 状态转移
	for col := 0; col < len(matrix); col++ {
		for row := 0; row < len(matrix[col]); row++ {
			if matrix[col][row] == '1' {
				// 初值
				if col == 0 || row == 0 {
					dp[col][row] = 1
				} else {
					dp[col][row] = Min(dp[col-1][row], dp[col][row-1], dp[col-1][row-1]) + 1
				}
			}
			ans = Max(ans, dp[col][row])
		}
	}
	return ans * ans
}
