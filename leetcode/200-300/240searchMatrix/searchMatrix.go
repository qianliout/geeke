package main

func main() {

}

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	col, row := 0, len(matrix[0])-1
	for col < len(matrix) && row >= 0 {
		if matrix[col][row] == target {
			return true
		} else if matrix[col][row] > target {
			row--
		} else if matrix[col][row] < target {
			col++
		}
	}

	return false
}
