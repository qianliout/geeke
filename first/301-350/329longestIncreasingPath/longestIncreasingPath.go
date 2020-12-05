package main

import (
	"fmt"
	. "outback/leetcode/common"
)

func main() {
	// nums := [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {19, 18, 17, 16, 15, 14, 13, 12, 11, 10}, {20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, {39, 38, 37, 36, 35, 34, 33, 32, 31, 30}, {40, 41, 42, 43, 44, 45, 46, 47, 48, 49}, {59, 58, 57, 56, 55, 54, 53, 52, 51, 50}, {60, 61, 62, 63, 64, 65, 66, 67, 68, 69}, {79, 78, 77, 76, 75, 74, 73, 72, 71, 70}, {80, 81, 82, 83, 84, 85, 86, 87, 88, 89}, {99, 98, 97, 96, 95, 94, 93, 92, 91, 90}, {100, 101, 102, 103, 104, 105, 106, 107, 108, 109}, {119, 118, 117, 116, 115, 114, 113, 112, 111, 110}, {120, 121, 122, 123, 124, 125, 126, 127, 128, 129}, {139, 138, 137, 136, 135, 134, 133, 132, 131, 130}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	nums := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}

	res := longestIncreasingPath(nums)

	fmt.Println("res is ", res)
}

/*
   给定一个整数矩阵，找出最长递增路径的长度。
   对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。
   示例 1:
   输入: nums =
   [
     [9,9,4],
     [6,6,8],
     [2,1,1]
   ]
   输出: 4
   解释: 最长递增路径为 [1, 2, 6, 9]。
   示例 2:
   输入: nums =
   [
     [3,4,5],
     [3,2,6],
     [2,2,1]
   ]
   输出: 4
   解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
*/

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	col := len(matrix)
	// maxPath 表示i,j这个点所到到的最大长度，记忆化
	maxPath := make(map[int]map[int]int)

	max := 1
	for c := 0; c < col; c++ {
		for r := 0; r < len(matrix[c]); r++ {
			dfs(matrix, maxPath, c, r, &max)
		}
	}
	return max
}

func dfs(matrix [][]int, used map[int]map[int]int, c, r int, max *int) int {
	dir := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	if used[c] == nil {
		used[c] = make(map[int]int)
	}
	// 这一步,一是记录值，二是保证不重复
	if used[c][r] > 0 {
		return used[c][r]
	}
	used[c][r] = 1
	for i := 0; i < 4; i++ {
		newC := c + dir[i][0]
		newR := r + dir[i][1]

		if inRange(matrix, newC, newR) && matrix[newC][newR] > matrix[c][r] {
			used[c][r] = Max(used[c][r], dfs(matrix, used, newC, newR, max)+1)
		}
	}
	*max = Max(*max, used[c][r])
	return used[c][r]
}

func inRange(matric [][]int, c, r int) bool {
	if c < 0 || r < 0 || c >= len(matric) || r >= len(matric[0]) {
		return false
	}
	return true
}
