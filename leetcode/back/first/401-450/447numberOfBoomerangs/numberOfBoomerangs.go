package main

import (
	"fmt"
)

func main() {
	points := [][]int{{1, 0}, {0, 0}, {2, 0}}

	res := numberOfBoomerangs(points)
	fmt.Println("res is ", res)
}

/*
给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j
之间的距离和 i 和 k 之间的距离相等（需要考虑元组的顺序）。
找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
示例:
输入:
[[0,0],[1,0],[2,0]]
输出:
2
解释:
两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
*/

func numberOfBoomerangs(points [][]int) int {
	ans := 0
	for i := 0; i < len(points); i++ {
		mp := make(map[int]int)
		for j := 0; j < len(points); j++ {
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			mp[dx*dx+dy*dy]++
		}
		for _, v := range mp {
			// 这里没有理解，到底是为什么呢
			ans += v * (v - 1)
		}
	}

	return ans
}
