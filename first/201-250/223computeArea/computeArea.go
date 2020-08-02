package main

import (
	"fmt"
	"math"
	"sort"

	. "outback/leetcode/common"
)

func main() {
	//res := computeArea(-3, 0, -2, 4, 0, -1, 9, 2)
	res := computeArea(4, -5, 5, 0, -3, -3, 3, 3)
	fmt.Println("res is ", res)
}

/*
在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。
每个矩形由其左下顶点和右上顶点坐标表示，如图所示。
Rectangle Area
示例:
输入: -3, 0, 3, 4, 0, -1, 9, 2
输出: 45
链接：https://leetcode-cn.com/problems/rectangle-area
*/
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	res := make([][]int, 0)
	res = append(res, []int{A, B})
	res = append(res, []int{C, D})
	res = append(res, []int{E, F})
	res = append(res, []int{G, H})
	return compute(res)
}

func compute(point [][]int) int {
	maxArea := int(math.Abs(float64((point[0][0] - point[1][0]) * (point[0][1] - point[1][1]))))
	maxArea += int(math.Abs(float64((point[2][0] - point[3][0]) * (point[2][1] - point[3][1]))))

	// 计算重叠区间
	if Max(point[0][0], point[1][0]) <= Min(
		point[2][0], point[3][0]) || Min(point[0][0], point[1][0]) >= Max(point[2][0], point[3][0]) || Max(
		point[0][1], point[1][1]) <= Min(
		point[2][1], point[3][1]) || Min(point[0][1], point[1][1]) >= Max(point[2][1], point[3][1]) {
		return maxArea
	}
	xs := []int{point[0][0], point[1][0], point[2][0], point[3][0]}
	ys := []int{point[0][1], point[1][1], point[2][1], point[3][1]}

	sort.Ints(xs)
	sort.Ints(ys)
	return maxArea - (xs[2]-xs[1])*(ys[2]-ys[1])
}
