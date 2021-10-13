package main

func main() {

}

/*
包含整数的二维矩阵 M 表示一个图片的灰度。你需要设计一个平滑器来让每一个单元的灰度成为平均灰度 (向下舍入) ，
平均灰度的计算是周围的8个单元和它本身的值求平均，如果周围的单元格不足八个，则尽可能多的利用它们。
示例 1:
输入:
[[1,1,1],
 [1,0,1],
 [1,1,1]]
输出:
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]
解释:
对于点 (0,0), (0,2), (2,0), (2,2): 平均(3/4) = 平均(0.75) = 0
对于点 (0,1), (1,0), (1,2), (2,1): 平均(5/6) = 平均(0.83333333) = 0
对于点 (1,1): 平均(8/9) = 平均(0.88888889) = 0
*/
func imageSmoother(M [][]int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(M); i++ {
		res = append(res, make([]int, len(M[i])))
	}

	if len(M) == 0 || len(M[0]) == 0 {
		return res
	}
	// 赋值
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {
			res[i][j] = help(M, i, j)

		}
	}
	return res
}

func help(M [][]int, x, y int) int {
	dx := []int{-1, 1, 0, 0, 1, 1, -1, -1}
	dy := []int{0, 0, -1, 1, 1, -1, 1, -1}
	all, n := M[x][y], 1
	for i := 0; i < 8; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx < 0 || ny < 0 || nx >= len(M) || ny >= len(M[0]) {
			continue
		} else {
			all += M[nx][ny]
			n++
		}
	}
	return all / n
}
