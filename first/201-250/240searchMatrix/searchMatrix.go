package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}}
	res := searchMatrix3(matrix, 4)

	//nums := []int{1, 3}
	//res := binersearch(nums, 2)

	fmt.Println("res is ", res)
}

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
    每行的元素从左到右升序排列。
    每列的元素从上到下升序排列。
示例:
现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target = 5，返回 true。
给定 target = 20，返回 false
链接：https://leetcode-cn.com/problems/search-a-2d-matrix-ii
*/

// 暴力法
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 因为是有序的，所以可以使用二分搜索
func searchMatrix2(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	//只是对行进行二分
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) == 0 {
			continue
		}
		if matrix[i][0] > target {
			return false
		}
		left, right := 0, len(matrix[i])-1
		for left <= right {
			mid := left + (right-left)/2
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] > target {
				right = mid - 1
			} else if matrix[i][mid] < target {
				left = mid + 1
			}
		}
	}
	return false
}

// 从右上角往下找
func searchMatrix3(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	left := len(matrix[0]) - 1 // 往左走
	down := 0

	for left >= 0 && down <= len(matrix)-1 {
		if matrix[down][left] == target {
			return true
		} else if matrix[down][left] > target {
			left--
		} else if matrix[down][left] < target {
			down++
		}
	}
	return false
}
