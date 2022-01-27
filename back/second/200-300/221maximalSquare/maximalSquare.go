package main

import (
	"qianliout/leetcode/back/common"
)

func main() {

}

/*
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
示例:
输入:
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0
输出: 4
*/

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// dp[i][j]表示，下标为i,j选重时的最大边长
	dp := make(map[int]map[int]int)
	// 初始化并赋初值
	for i := 0; i < len(matrix); i++ {
		dp[i] = map[int]int{}
	}

	max := 0
	for col := 0; col < len(matrix); col++ {
		for row := 0; row < len(matrix[col]); row++ {
			if matrix[col][row] == '1' {
				if col == 0 || row == 0 {
					dp[col][row] = 1
				} else {
					dp[col][row] = common.Min(dp[col-1][row], dp[col][row-1], dp[col-1][row-1]) + 1
				}
				if dp[col][row] > max {
					max = dp[col][row]
				}
			} else {
				dp[col][row] = 0
			}

		}
	}
	return max * max
}
