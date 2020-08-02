package main

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
}

/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。
示例 1:
输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2:
输入:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
*/
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	// 最简单的方式,新建另一个矩阵
	othrer := make([][]int, len(matrix))

	for i, v1 := range matrix {
		for j, v2 := range v1 {
			if othrer[i] == nil {
				othrer[i] = make([]int, len(matrix[0]))
			}
			othrer[i][j] = v2
		}
	}
	for i, v1 := range othrer {
		for j, v2 := range v1 {
			if v2 == 0 {
				full(&matrix, i, j)
			}
		}
	}
}

func full(matrix *[][]int, col, row int) {
	for i := 0; i < len((*matrix)[col]); i++ {
		(*matrix)[col][i] = 0
	}
	for i := 0; i < len(*matrix); i++ {
		(*matrix)[i][row] = 0
	}
}
