package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12}}
	// {13, 14, 15, 16}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
示例 1:
输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:
输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	up, below, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	all := len(matrix) * len(matrix[0])

	for {
		// 向左到最后
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++

		// 再向下
		for i := up; i <= below; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// 再向右
		for i := right; i >= left; i-- {
			res = append(res, matrix[below][i])
		}
		below--

		// 再向上
		for i := below; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++

		if len(res) >= all {
			break
		}
	}
	return res[:all]
}
