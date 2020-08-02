package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16}}
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
	exitMap := make(map[string]bool)
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		c := first[0]
		r := first[1]
		exitMap[fmt.Sprintf("%d_%d", c, r)] = true
		res = append(res, matrix[c][r])
		isTrue, i, j := next(matrix, c, r, &exitMap)
		if !isTrue {
			break
		}
		queue = append(queue, [2]int{i, j})

	}
	return res
}

func next(matrix [][]int, i, j int, exitMap *map[string]bool) (bool, int, int) {
	col := len(matrix)
	row := len(matrix[0])

	// 向上
	if (j-i < 0 || (*exitMap)[fmt.Sprintf("%d_%d", i, j-1)]) && (i < col) {
		if !(*exitMap)[fmt.Sprintf("%d_%d", i-1, j)] {
			return true, i - 1, j
		} else {
			return false, 0, 0
		}
	}


	
	if i-1 < col && i-1 >= 0 && j < row && !(*exitMap)[fmt.Sprintf("%d_%d", i-1, j)] {
		return true, i - 1, j
	}

	if j+1 < row && i < col && !(*exitMap)[fmt.Sprintf("%d_%d", i, j+1)] {
		return true, i, j + 1
	}

	if i+1 < col && j < row && !(*exitMap)[fmt.Sprintf("%d_%d", i+1, j)] {
		return true, i + 1, j
	}

	if j-1 < row && j-1 >= 0 && i < col && !(*exitMap)[fmt.Sprintf("%d_%d", i, j-1)] {
		return true, i, j - 1
	}

	return false, 0, 0
}
