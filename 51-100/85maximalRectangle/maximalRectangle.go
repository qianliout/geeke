package main

import (
	"fmt"
)

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	res := maximalRectangle(matrix)
	fmt.Println("res is ", res)
}

/*
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
示例:
输入:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
输出: 6
*/

// 按层带入84就可以了
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	row := len(matrix[0])
	col := len(matrix)
	maxArea := 0
	for i := 0; i < col; i++ {
		//h := CalculatedHeight(matrix, row, i)
		res := largestRectangleArea(CalculatedHeight(matrix, row, i))
		//fmt.Println("i is ", i, res, h)
		if res > maxArea {
			maxArea = res
		}
	}
	return maxArea
}

func CalculatedHeight(matrix [][]byte, row, col int) []int {
	res := make([]int, row)

	for j := 0; j < row; j++ {
		if matrix[col][j] == '0' {
			res[j] = 0
		} else {
			h := 0
			for i := col; i >= 0; i-- {
				if matrix[i][j] == '1' {
					h++
				} else {
					break
				}
			}
			res[j] = h
		}
	}
	return res
}

func largestRectangleArea(heights []int) int {
	stark := make([]int, 0)
	stark = append(stark, -1)
	maxArea := 0
	for i, _ := range heights {
		for len(stark) > 1 && heights[i] <= heights[stark[len(stark)-1]] {
			pop := stark[len(stark)-1]
			//这里要特别注意,一定是先pop一个出来,再计算宽度,不然就会出错,
			stark = stark[:len(stark)-1]
			res := heights[pop] * (i - stark[len(stark)-1] - 1)
			if res > maxArea {
				maxArea = res
			}
		}
		stark = append(stark, i)
		//fmt.Println(stark, maxArea)
	}

	for len(stark) > 1 {
		pop := stark[len(stark)-1]
		stark = stark[:len(stark)-1]
		res := heights[pop] * (len(heights) - stark[len(stark)-1] - 1)
		if res > maxArea {
			maxArea = res
		}
	}
	return maxArea
}
