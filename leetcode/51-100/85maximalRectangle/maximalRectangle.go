package main

import (
	"fmt"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	area := maximalRectangle(matrix)
	fmt.Println("area is ", area)
	fmt.Println(1 << 7)
	fmt.Println(1 << 8)

	fmt.Println(ExistFlag(128, 7))
	fmt.Println(ExistFlag(0, 0))
}

func ExistFlag(value uint64, flag int64) bool {
	return (value>>flag)&1 == 1
}

/*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
*/
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	for i := 0; i < len(matrix); i++ {

		res := largestRectangleArea(CalculatedHeight(matrix, i))
		maxArea = Max(maxArea, res)
	}
	return maxArea
}

func CalculatedHeight(matrix [][]byte, col int) []int {
	res := make([]int, 0)
	for j := 0; j < len(matrix[0]); j++ {
		ans := 0
		for i := col; i >= 0; i-- {
			if matrix[i][j] == '1' {
				ans++
			} else {
				break
			}
		}
		res = append(res, ans)
	}
	return res
}

func largestRectangleArea(heights []int) int {
	heights = append(append([]int{0}, heights...), 0)
	ans, stark := 0, make([]int, 0)
	for i := 0; i < len(heights); i++ {
		// 单调递增栈
		for len(stark) > 0 && heights[stark[len(stark)-1]] > heights[i] {
			top := heights[stark[len(stark)-1]]
			stark = stark[:len(stark)-1]
			right, left := i-1, stark[len(stark)-1]
			ans = Max(ans, top*(right-left))
		}
		stark = append(stark, i)
	}
	return ans
}
