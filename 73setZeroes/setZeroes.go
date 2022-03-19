package main

func main() {

}

/*
给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
示例 1：
输入：matrix = [[1,1,1],[1,0,1],[1,1,1]]
输出：[[1,0,1],[0,0,0],[1,0,1]]
*/

func setZeroes(matrix [][]int) {
	fullBlood(&matrix)
}

func fullBlood(matrix *[][]int) {
	row0, col0 := false, false
	for i := 0; i < len(*matrix); i++ {
		if (*matrix)[i][0] == 0 {
			col0 = true
			break
		}
	}
	for j := 0; j < len((*matrix)[0]); j++ {
		if (*matrix)[0][j] == 0 {
			row0 = true
			break
		}
	}
	for i := 1; i < len(*matrix); i++ {
		for j := 1; j < len((*matrix)[i]); j++ {
			if (*matrix)[i][j] == 0 {
				(*matrix)[i][0], (*matrix)[0][j] = 0, 0
			}
		}
	}

	for i := 1; i < len(*matrix); i++ {
		for j := 1; j < len((*matrix)[i]); j++ {
			if (*matrix)[i][0] == 0 || (*matrix)[0][j] == 0 {
				(*matrix)[i][j] = 0
			}
		}
	}

	if col0 {
		for i := 0; i < len(*matrix); i++ {
			(*matrix)[i][0] = 0
		}
	}
	if row0 {
		for i := 0; i < len((*matrix)[0]); i++ {
			(*matrix)[0][i] = 0
		}
	}
}
