package main

import (
	"math"

	"qianliout/leetcode/common/utils"
)

func main() {

}

/*
给定一个非空二维矩阵 matrix 和一个整数 k，找到这个矩阵内部不大于 k 的最大矩形和。
示例:
输入: matrix = [[1,0,1],
               [0,-2,3]], k = 2
输出: 2
解释: 矩形区域 [[0, 1], [-2, 3]] 的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。
说明：
    矩阵内的矩形区域面积必须大于 0。
    如果行数远大于列数，你将如何解答呢？
*/

// 四层循环

type item struct {
	i int
	j int
}

// 这种方式可以得到正确的值，但是会超时
func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	max := math.MinInt64
	rows := len(matrix[0])
	cols := len(matrix)

	for i1 := 1; i1 <= cols; i1++ {
		for j1 := 1; j1 <= rows; j1++ {
			dp := make(map[[2]int]int)
			dp[[2]int{i1, j1}] = matrix[i1-1][j1-1]
			for i2 := i1; i2 <= cols; i2++ {
				for j2 := j1; j2 <= rows; j2++ {
					dp[[2]int{i2, j2}] = dp[[2]int{i2 - 1, j2}] + dp[[2]int{i2, j2 - 1}] - dp[[2]int{i2 - 1, j2 - 1}] + matrix[i2-1][j2-1]
					if dp[[2]int{i2, j2}] > max && dp[[2]int{i2, j2}] <= k {
						max = dp[[2]int{i2, j2}]
					}
				}
			}
		}
	}
	return max
}

func maxSumSubmatrix1(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	max := math.MinInt64
	rows := len(matrix[0])
	cols := len(matrix)

	for lr := 0; lr < rows; lr++ { // 固定左边界
		for rr := lr; rr < rows; rr++ { // 探寻右边界
			rowSum := make([]int, cols)
			for c := 0; c < cols; c++ {
				rowSum[c] += matrix[c][rr]
			}
			max = utils.Max(max, dpmax(rowSum, k))
		}
	}
	return max
}

func dpmax(sumRow []int, k int) int {
	max := math.MinInt64
	for i := 0; i < len(sumRow); i++ {
		sum := 0
		for j := i; j < len(sumRow); j++ {
			sum += sumRow[j]
			if sum > max && sum <= k {
				max = sum
			}
		}
	}
	return max
}
