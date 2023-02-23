package main

import (
	"fmt"
	. "qianliout/leetcode/common/unionfind"
)

func main() {
	//grid := [][]int{
	//	{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
	//	{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	//	{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	//}

	grid := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 1, 1},
	}
	//res := maxAreaOfIsland(grid)
	res := bfs(grid)
	fmt.Println("res is ", res)

}

/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
示例 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。
示例 2:
[[0,0,0,0,0,0,0,0]]
*/

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	ans := 0
	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid[col]); row++ {
			if grid[col][row] == 1 {
				num := dfs(&grid, col, row)
				if ans < num {
					ans = num
				}
			}
		}
	}
	return ans
}

// dfs 计算的是grid[col][row]=1时，他的面积是多少
func dfs(grid *[][]int, col, row int) int {
	if col < 0 || row < 0 || col >= len(*grid) || row >= len((*grid)[col]) || (*grid)[col][row] == 0 {
		return 0
	}
	num := 1
	// 沉岛思想
	(*grid)[col][row] = 0
	num += dfs(grid, col-1, row)
	num += dfs(grid, col+1, row)
	num += dfs(grid, col, row+1)
	num += dfs(grid, col, row-1)
	return num
}

// 注意，这种方法是对的，但是如果是全一的情况，且数据量比较大时，会把内存用完
func bfs(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	num := 0

	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid[col]); row++ {
			if grid[col][row] == 1 {
				stark := make([][2]int, 0)
				thisNum := 0
				stark = append(stark, [2]int{col, row})
				for len(stark) > 0 {
					s := stark[0]
					if grid[s[0]][s[1]] == 1 {
						thisNum++
						grid[s[0]][s[1]] = 0
					}
					stark = stark[1:]
					// 加数据
					if s[0]-1 >= 0 && grid[s[0]-1][s[1]] == 1 {
						stark = append(stark, [2]int{s[0] - 1, s[1]})
					}
					if s[0]+1 < len(grid) && grid[s[0]+1][s[1]] == 1 {
						stark = append(stark, [2]int{s[0] + 1, s[1]})
					}
					if s[1]-1 >= 0 && grid[s[0]][s[1]-1] == 1 {
						stark = append(stark, [2]int{s[0], s[1] - 1})
					}
					if s[1]+1 < len(grid[col]) && grid[s[0]][s[1]+1] == 1 {
						stark = append(stark, [2]int{s[0], s[1] + 1})
					}
				}
				if thisNum > num {
					num = thisNum
				}
			}
		}
	}
	return num
}
