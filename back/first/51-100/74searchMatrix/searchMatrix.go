package main

import (
	"fmt"
)

func main() {
	//matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}, {25, 30, 34, 50}}
	matrix := [][]int{{1}, {3}, {23, 30, 34, 50}, {25, 30, 34, 50}}
	res := searchMatrix(matrix, 0)
	fmt.Println("res is ", res)
}

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
    每行中的整数从左到右按升序排列。
    每行的第一个整数大于前一行的最后一个整数。
示例 1:
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:
输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	col := findx(matrix, target)
	fmt.Println("col is ", col)
	return findy(matrix, col, target)
}

func findx(matrix [][]int, target int) int {
	left, right := 0, len(matrix)-1

	for left <= right || (left < right && matrix[left][0] <= target && target < matrix[right][0] && right-left == 1) {
		mid := left + (right-left)/2
		//fmt.Println("left is ", left, "right is ", right, "mid is ", mid)
		if matrix[mid][0] == target {
			return mid
		} else if matrix[mid][0] <= target {
			left = mid + 1
		} else if matrix[mid][0] > target {
			right = mid - 1
		}
	}
	if left > len(matrix)-1 {
		return len(matrix) - 1
	}
	return left - 1
}

func findy(matrix [][]int, col, target int) bool {
	if col < 0 {
		return false
	}
	left, right := 0, len(matrix[col])-1

	for left <= right {
		mid := left + (right-left)/2
		//fmt.Println("left is ", left, "right is ", right, "mid is ", mid)
		if matrix[col][mid] == target {
			return true
		} else if matrix[col][mid] <= target {
			left = mid + 1
		} else if matrix[col][mid] > target {
			right = mid - 1
		}
	}
	if left > len(matrix)-1 || (left-1 == right && matrix[col][left-1] != target) {
		return false
	}
	return true
}
