package main

import (
	"fmt"

	"outback/leetcode/common/unionfind"
)

func main() {
	grid := [][]byte{
		[]byte("11010"),
		[]byte("11010"),
		[]byte("11010"),
		[]byte("00000"),
	}
	res := numIslands(grid)
	fmt.Println("nums is ", res)
}

/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上
相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。
示例 1:
输入:
11110
11010
11000
00000
输出: 1
示例 2:
输入:
11000
11000
00100
00011
链接：https://leetcode-cn.com/problems/number-of-islands
*/

// 本题使用并查集是最好的解法
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	col := len(grid)
	row := len(grid[0])

	uf := unionfind.NewUnionFind(col * row)
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			if grid[i][j] == '0' {
				uf.Count--
			}
			if grid[i][j] == '1' {
				// 上下左右合并
				if i > 0 && grid[i-1][j] == '1' {
					uf.Union(i*row+j, (i-1)*row+j)
				}

				if i < col-1 && grid[i+1][j] == '1' {
					uf.Union(i*row+j, (i+1)*row+j)
				}

				if j > 0 && grid[i][j-1] == '1' {
					uf.Union(i*row+j, i*row+j-1)
				}
				if j < row-1 && grid[i][j+1] == '1' {
					uf.Union(i*row+j, i*row+j+1)
				}
			}
		}
	}

	// 找有多少个

	return uf.Count
}

func Islands(i, j int, uf *unionfind.UnionFind, grid [][]byte) bool {
	col := len(grid)
	row := len(grid[0])

	uif := uf.Find(i*row + j)

	if i > 0 && grid[i-1][j] == '1' {
		if uf.Find((i-1)*row+j) == uif {
			return false
		}
	}

	if i < col-1 && grid[i+1][j] == '1' {
		if uif == uf.Find((i+1)*row+j) {
			return false
		}
	}

	if j > 0 && grid[i][j-1] == '1' {
		if uif == uf.Find(i*row+j-1) {
			return false
		}
	}
	if j < row-1 && grid[i][j+1] == '1' {
		if uif == uf.Find(i*row+j+1) {
			return false
		}
	}
	return true
}
