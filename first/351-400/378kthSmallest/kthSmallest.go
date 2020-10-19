package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2},
		{1, 3},
		//{12, 13, 15},
	}
	res := kthSmallest(matrix, 2)
	fmt.Println("res is ", res)

}

/*
给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第k小的元素。
请注意，它是排序后的第k小元素，而不是第k个元素。
示例:
matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,
返回 13。
*/

// 题目中说是它是排序后的第k小元素，而不是第k个元素,所以可能有重复
// 这种方法可以找出第k个元素，但无法满足题目中的要求,如果没有重复就可以满足
func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1
	}
	row := len(matrix[0]) // 一行有多少个
	col := len(matrix)
	left, down := row-1, 0
	num := row
	for left >= 0 && down < col {
		if num == k {
			return matrix[down][left]
		}
		if num < k {
			down++
			num += row
		}
		if num > k {
			left--
			num--
		}
	}
	return -1
}

// 这道题可以使用大顶堆，加map去重的方式
// 因为可以有重复元素，其实使用二分法其实是不好做的
func kthSmallest2(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1
	}
	rowNum := len(matrix[0])
	col, row := 0, rowNum-1
	num := rowNum
	for col < len(matrix) && row >= 0 {
		if num == k {
			return matrix[col][row]
		} else if num < k {
			col++
			num += rowNum
		} else if num > k {
			row--
			num--
		}
	}
	return -1
}
