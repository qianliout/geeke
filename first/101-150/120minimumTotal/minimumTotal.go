package main

import (
	"fmt"

	. "outback/leetcode/common"
)

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		//{6, 5, 7},
		//{4, 1, 8, 3},
	}

	res := minimumTotal(triangle)
	fmt.Println("res is ", res)
}

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
例如，给定三角形：
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
说明：
如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
*/
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || (len(triangle) == 1 && len(triangle[0]) == 0) {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	// 从底从上找
	for i := len(triangle) - 2; i >= 0; i-- {
		row := len(triangle[i])
		for j := 0; j < row; j++ {
			triangle[i][j] = Min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	//fmt.Println("triahgd is ", triangle)
	return triangle[0][0]
}
