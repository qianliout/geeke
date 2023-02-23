package main

import (
	"fmt"
)

func main() {
	res := generateMatrix(1)
	fmt.Println("res is ", res)
}

/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
示例:
输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

func generateMatrix(n int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}
	if n == 0 {
		return res
	}
	if n == 1 {
		res[0][0] = 1
		return res
	}

	up, down, left, right := 0, n-1, 0, n-1
	level, i := 1, 1
	for level < n {
		// 向左
		for j := left; j <= right; j++ {
			res[up][j] = i
			i++
		}
		up++

		// 再向下
		for j := up; j <= down; j++ {
			res[j][right] = i
			i++
		}
		right--

		// 再向右
		for j := right; j >= left; j-- {
			res[down][j] = i
			i++
		}
		down--
		// 再向上
		for j := down; j >= up; j-- {
			res[j][left] = i
			i++
		}
		left++

		level++
	}
	return res
}
