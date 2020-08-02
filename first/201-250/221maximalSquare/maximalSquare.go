package main

import (
	. "outback/leetcode/common"
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
	// dp[i][j]表示第i行,第j列选择时边长的最大值 ,注意,这里不是面积
	// 初值
	dp := make(map[int]map[int]int)
	// 初始化
	for i := 0; i <= len(matrix); i++ {
		dp[i] = make(map[int]int)
	}
	// 初值
	result := 0
	//初始化base case
	for col := 0; col < len(matrix); col++ {
		if matrix[col][0] == '1' {
			dp[col][0] = 1
			if dp[col][0] > result {
				result = dp[col][0]
			}
		}
	}

	for row := 0; row < len(matrix[0]); row++ {
		if matrix[0][row] == '1' {
			dp[0][row] = 1
			if dp[0][row] > result {
				result = dp[0][row]
			}
		}
	}
	// 状态转移方程
	for col := 1; col < len(matrix); col++ {
		for row := 1; row < len(matrix[col]); row++ {
			if matrix[col][row] == '1' {
				dp[col][row] = Min(dp[col-1][row], dp[col][row-1], dp[col-1][row-1]) + 1
			}
			if dp[col][row] > result {
				result = dp[col][row]
			}
		}
	}
	return result * result
}
