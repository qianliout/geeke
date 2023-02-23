package main

import (
	"fmt"
	"math"
)

func main() {

	nums := [][]int{
		{-2, -2, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	res := calculateMinimumHP(nums)
	fmt.Println("res is ", res)

}

/*
一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）
最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。
骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。
有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；
其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。
为了尽快到达公主，骑士决定每次只向右或向下移动一步。
编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。
例如，考虑到如下布局的地下城，如果骑士遵循最佳路径 右 -> 右 -> 下 -> 下，则骑士的初始健康点数至少为 7。
-2 (K) 	-3 	3
-5 	-10 	1
10 	30 	-5 (P)
说明:
    骑士的健康点数没有上限。
    任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。
*/

func calculateMinimumHP(dungeon [][]int) int {
	mem := make(map[[2]int]int)

	min := dfsWithMen(dungeon, 0, 0, &mem)
	return min
}

// 表示从[c,r]出发到终点所需要的最小血量，这个定义要好好理解
// 因为没有加记忆化，所以会超时，
func dfs(dungeon [][]int, c, r int) int {
	// 到达终点了
	if c == len(dungeon)-1 && r == len(dungeon[c])-1 {
		return Max(1, 1-dungeon[c][r])
	}

	// 到最下边了，只能往右走了
	if c+1 >= len(dungeon) {
		return Max(1, dfs(dungeon, c, r+1)-dungeon[c][r])
	}
	if r+1 >= len(dungeon[c]) {
		return Max(1, dfs(dungeon, c+1, r)-dungeon[c][r])
	}

	return Max(
		1,
		Min(dfs(dungeon, c+1, r), dfs(dungeon, c, r+1))-dungeon[c][r],
	)
}

func inRange(dungeon [][]int, c, r int) bool {
	if c < 0 || r < 0 || c >= len(dungeon) || r >= len(dungeon[0]) {
		return false
	}
	return true
}

func Min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func dfsWithMen(dungeon [][]int, c, r int, mem *map[[2]int]int) int {
	if (*mem)[[2]int{c, r}] > 0 {
		return (*mem)[[2]int{c, r}]
	}

	// 到达终点了
	if c == len(dungeon)-1 && r == len(dungeon[c])-1 {
		return Max(1, 1-dungeon[c][r])
	}

	var m int
	if c+1 >= len(dungeon) {
		m = Max(1, dfsWithMen(dungeon, c, r+1, mem)-dungeon[c][r])
	} else if r+1 >= len(dungeon[c]) {
		m = Max(1, dfsWithMen(dungeon, c+1, r, mem)-dungeon[c][r])
	} else {
		m = Max(
			1,
			Min(dfsWithMen(dungeon, c+1, r, mem), dfsWithMen(dungeon, c, r+1, mem))-dungeon[c][r],
		)
	}
	(*mem)[[2]int{c, r}] = m
	return m
}
