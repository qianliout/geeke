package main

func main() {

}

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素
*/
func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}
	left, right, up, below := 0, len(matrix[0])-1, 0, len(matrix)-1
	all := len(matrix) * len(matrix[0])
	for {
		// 先左
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}
		up++
		// 再向下
		for i := up; i <= below; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 向右
		for i := right; i >= left; i-- {
			res = append(res, matrix[below][i])
		}
		below--
		// 向上
		for i := below; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		// 这里写退出循环代码很重要
		if len(res) >= all {
			break
		}
	}
	return res[:all]
}
