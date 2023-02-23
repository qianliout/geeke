package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}
	res := countBattleships2(board)
	fmt.Println(res)
}

/*
给定一个二维的甲板， 请计算其中有多少艘战舰。 战舰用 'X'表示，空位用 '.'表示。 你需要遵守以下规则：
    给你一个有效的甲板，仅由战舰或者空位组成。
    战舰只能水平或者垂直放置。换句话说,战舰只能由 1xN (1 行, N 列)组成，或者 Nx1 (N 行, 1 列)组成，其中N可以是任意大小。
    两艘战舰之间至少有一个水平或垂直的空位分隔 - 即没有相邻的战舰。
示例 :
X..X
...X
...X
在上面的甲板中有2艘战舰。
无效样例 :
...X
XXXX
...X
你不会收到这样的无效甲板 - 因为战舰之间至少会有一个空位将它们分开。
进阶:
你可以用一次扫描算法，只使用O(1)额外空间，并且不修改甲板的值来解决这个问题吗？
*/

// 先用沉岛法
func countBattleships(board [][]byte) int {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}
	used := make(map[[2]int]bool)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				used[[2]int{i, j}] = true

				dfs(&board, i+1, j, &used)
				dfs(&board, i-1, j, &used)
				dfs(&board, i, j+1, &used)
				dfs(&board, i, j-1, &used)
			}
		}
	}
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				res++
			}
		}
	}
	return res
}

func dfs(board *[][]byte, i, j int, used *map[[2]int]bool) {
	if !inrange(board, i, j) || (*board)[i][j] == '.' || (*used)[[2]int{i, j}] {
		return
	}
	(*board)[i][j] = '.'
	(*used)[[2]int{i, j}] = true
	dfs(board, i+1, j, used)
	dfs(board, i-1, j, used)
	dfs(board, i, j+1, used)
	dfs(board, i, j-1, used)
}

func inrange(board *[][]byte, i, j int) bool {
	if i >= len(*board) || j >= len((*board)[0]) || i < 0 || j < 0 {
		return false
	}
	return true
}

// 这道题使用沉岛感觉大材小用了，因为我题目中：
// 1,你不会收到这样的无效甲板
// 战舰只能水平或者垂直放置。换句话说,战舰只能由 1xN (1 行, N 列)组成，或者 Nx1 (N 行, 1 列)组成，其中N可以是任意大小。
// 两艘战舰之间至少有一个水平或垂直的空位分隔 - 即没有相邻的战舰。

func countBattleships2(board [][]byte) int {
	if len(board) == 0 || len(board[0]) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X' {
					//说明这一行都不可能有第二个战机了
					continue
				}
				res++
			}
		}
	}
	return res
}
