package main

import (
	"fmt"
)

func main() {
	matrix := generateMatrix(3)
	fmt.Println("matrix ", matrix)
}

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
*/
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	i, left, right, up, low := 1, 0, n-1, 0, n-1
	for i <= n*n {
		for j := left; j <= right; j++ {
			ans[up][j] = i
			i++
		}
		up++
		for j := up; j <= low; j++ {
			ans[j][right] = i
			i++
		}
		right--
		for j := right; j >= left; j-- {
			ans[low][j] = i
			i++
		}
		low--
		for j := low; j >= up; j-- {
			ans[j][left] = i
			i++
		}
		left++
	}
	return ans
}
