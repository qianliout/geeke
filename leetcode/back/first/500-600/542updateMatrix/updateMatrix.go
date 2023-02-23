package main

import (
	"fmt"
)

func main() {
	mums := [][]int{{1, 1, 1}, {1, 1, 1}, {0, 1, 0}}
	ans := updateMatrix(mums)
	fmt.Println("ans is ", ans)
}

/*
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
两个相邻元素间的距离为 1 。
示例 1:
输入:
0 0 0
0 1 0
0 0 0
输出:
0 0 0
0 1 0
0 0 0
示例 2:
输入:
0 0 0
0 1 0
1 1 1
输出:
0 0 0
0 1 0
1 2 1
注意:
    给定矩阵的元素个数不超过 10000。
    给定矩阵中至少有一个元素是 0。
    矩阵中的元素只在四个方向上相邻: 上、下、左、右。
*/

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	queue := make([][2]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]
		x := first[0]
		y := first[1]
		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if inrange(&matrix, nx, ny) && matrix[nx][ny] == -1 {
				matrix[nx][ny] = matrix[x][y] + 1
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}
	return matrix
}

func inrange(matrix *[][]int, col, row int) bool {
	if col < 0 || row < 0 || col >= len(*matrix) || row >= len((*matrix)[0]) {
		return false
	}
	return true
}
