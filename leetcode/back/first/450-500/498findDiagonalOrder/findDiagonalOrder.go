package main

import (
	"fmt"
)

func main() {
	matr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := findDiagonalOrder(matr)
	fmt.Println("res is ", res)
}

/*
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。
示例:
输入:
[

	[ 1, 2, 3 ],
	[ 4, 5, 6 ],
	[ 7, 8, 9 ]

]

输出:  [1,2,4,7,5,3,6,8,9]
解释:
说明:

	给定矩阵中的元素总数不会超过 100000 。
*/
func findDiagonalOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	row, col := 0, 0
	for i := 0; i < len(matrix)+len(matrix[0])-1; i++ {
		if i%2 == 0 {
			for col >= 0 && row < len(matrix[0]) {
				res = append(res, matrix[col][row])
				col--
				row++
			}
			if row < len(matrix[0]) {
				col++
			} else {
				col += 2
				row--
			}

		} else {
			for row >= 0 && col < len(matrix) {
				res = append(res, matrix[col][row])
				col++
				row--
			}
			if col < len(matrix) {
				row++
			} else {
				row += 2
				col--
			}

		}
	}
	return res
}
func nex(row, col int, up bool) (int, int) {
	if up {
		return row + 1, col - 1
	}
	return row - 1, col + 1
}

func inRange(matrix *[][]int, row, col int) bool {
	if row < 0 || row >= len((*matrix)[0]) || col < 0 || col >= len(*matrix) {
		return false
	}
	return true
}
