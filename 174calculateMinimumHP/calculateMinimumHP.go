package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	res := calculateMinimumHP(dungeon)
	fmt.Println("res is ", res)
}

/*
一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，
他必须穿过地下城并通过对抗恶魔来拯救公主。
骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为 0），
要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
为了尽快到达公主，骑士决定每次只向右或向下移动一步。
编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。
例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。
-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)
*/
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	col, row := len(dungeon)-1, len(dungeon[0])-1
	// 到i,j时必须保有的血量(最少为0)
	dp := make([][]int, len(dungeon))
	for i := range dp {
		dp[i] = make([]int, len(dungeon[i]))
	}
	// 初值
	dp[col][row] = Max(0, -dungeon[col][row])
	for i := col - 1; i >= 0; i-- {
		dp[i][row] = Max(0, dp[i+1][row]-dungeon[i][row])
	}
	for j := row - 1; j >= 0; j-- {
		dp[col][j] = Max(0, dp[col][j+1]-dungeon[col][j])
	}
	// 遍历
	for i := col - 1; i >= 0; i-- {
		for j := row - 1; j >= 0; j-- {
			dp[i][j] = Max(0, Min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])

		}
	}
	// 至少要有1的血量，才能不死
	return dp[0][0] + 1
}
