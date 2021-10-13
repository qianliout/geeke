package main

func main() {

}

/*
给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
示例 :
输入:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]
输出: 16
解释: 它的周长是下面图片中的 16 个黄色的边：
*/

// 题很简单，每个1都有4个边，但是，只要有另一个1和他相连，就会少一个边，所以直接计算就可以了
func islandPerimeter(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				ans += 4 - around(grid, i, j)
			}
		}
	}
	return ans
}

func around(grid [][]int, i, j int) int {
	a := 0
	isrange := func(i, j int) bool {
		if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) {
			return false
		}
		return true
	}
	if isrange(i-1, j) && grid[i-1][j] == 1 {
		a++
	}
	if isrange(i+1, j) && grid[i+1][j] == 1 {
		a++
	}
	if isrange(i, j+1) && grid[i][j+1] == 1 {
		a++
	}
	if isrange(i, j-1) && grid[i][j-1] == 1 {
		a++
	}
	return a
}
