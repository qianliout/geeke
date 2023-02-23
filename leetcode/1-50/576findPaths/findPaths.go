package main

import (
	"fmt"
)

/*
给你一个大小为 m x n 的网格和一个球。球的起始坐标为 [startRow, startColumn] 。你可以将球移到在四个方向上相邻的单元格内（可以穿过网格边界到达网格之外）。你 最多 可以移动 maxMove 次球。
给你五个整数 m、n、maxMove、startRow 以及 startColumn ，找出并返回可以将球移出边界的路径数量。因为答案可能非常大，返回对 109 + 7 取余 后的结果。
*/
//  https://leetcode-cn.com/problems/out-of-boundary-paths/solution/gong-shui-san-xie-yi-ti-shuang-jie-ji-yi-asrz/

func main() {
	// m = 1, n = 3, maxMove = 3, startRow = 0, startColumn = 1
	res := findPaths(8, 7, 16, 1, 5)
	fmt.Println("res is ", res)
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	way := NewWay(m, n, maxMove)
	return way.dfs(startRow, startColumn, maxMove)
}

type way struct {
	m       int
	n       int
	maxMove int
	walked  map[int]map[int]map[int]int // 三维数组
	MOD     int
	dri     [][]int
}

func NewWay(m int, n int, maxMove int) *way {
	walked := make(map[int]map[int]map[int]int)
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			for k := 0; k <= maxMove; k++ {
				if walked[i] == nil {
					walked[i] = make(map[int]map[int]int)
				}
				if walked[i][j] == nil {
					walked[i][j] = make(map[int]int)
				}
				walked[i][j][k] = -1
			}
		}
	}
	return &way{
		m:       m,
		n:       n,
		maxMove: maxMove,
		walked:  walked,
		MOD:     1e9 + 7,
		dri:     [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}},
	}
}

func (w *way) dfs(startRow, startColumn, k int) int {
	if startRow >= w.m || startColumn >= w.n || startColumn < 0 || startRow < 0 {
		return 1
	}
	if k == 0 {
		return 0
	}
	if w.walked[startRow][startColumn][k] != -1 {
		return w.walked[startRow][startColumn][k]
	}
	ans := 0
	for i := range w.dri {
		ans += w.dfs(startRow+w.dri[i][0], startColumn+w.dri[i][1], k-1)
		ans = ans % w.MOD
	}
	w.walked[startRow][startColumn][k] = ans
	return ans
}
