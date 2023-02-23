package main

import (
	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {

}

/*
给定一个三角形 triangle ，找出自顶向下的最小路径和。
每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
*/
func minimumTotal(triangle [][]int) int {
	// 从下往上更新值
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] = Min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}
