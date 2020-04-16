package main

import (
	"fmt"
	"math"
)

type args struct {
	points [][]int // 参数结构
}

func main() {
	tests := []struct {
		name string // 测试名称
		args args   // 传入参数
		want int    // 预期数据
	}{
		{name: "没有点", args: args{points: [][]int{}}, want: 0},
		{name: "只有一个点", args: args{points: [][]int{
			[]int{1, 1},
		}}, want: 1},
		{name: "只有两个点", args: args{points: [][]int{
			[]int{1, 1},
			[]int{0, 0},
		}}, want: 2},
		{name: "有两个(1,1)的点", args: args{points: [][]int{
			[]int{1, 1},
			[]int{0, 0},
			[]int{1, 1},
		}}, want: 3},
		{name: "有两个点存在重复的情况(1,1), (0,0)", args: args{points: [][]int{
			[]int{1, 1},
			[]int{0, 0},
			[]int{1, 1},
			[]int{0, 0},
			[]int{2, 3},
		}}, want: 4},
		{name: "正常测试1", args: args{points: [][]int{
			[]int{1, 1},
			[]int{2, 2},
			[]int{3, 3},
		}}, want: 3},
		{name: "正常测试2", args: args{points: [][]int{
			[]int{1, 1},
			[]int{3, 2},
			[]int{5, 3},
			[]int{4, 1},
			[]int{2, 3},
			[]int{1, 4},
		}}, want: 4},
		{name: "三个点相同的点", args: args{points: [][]int{
			[]int{1, 1},
			[]int{1, 1},
			[]int{1, 1},
		}}, want: 3},
		{name: "斜率为0", args: args{points: [][]int{
			[]int{2, 3},
			[]int{3, 3},
			[]int{-5, 3},
			[]int{2, 1},
		}}, want: 3},
		{name: "斜率为无穷大", args: args{points: [][]int{
			[]int{3, -1},
			[]int{3, 2},
			[]int{3, 1},
			[]int{2, 1},
		}}, want: 3},
		{name: "k除不尽", args: args{points: [][]int{
			[]int{84, 250},
			[]int{0, 0},
			[]int{1, 0},
			[]int{0, -70},
			[]int{0, -70},
			[]int{1, -1},
			[]int{21, 10},
			[]int{42, 90},
			[]int{-42, -230},
		}}, want: 6},
		{name: "浮点数精度丢失", args: args{points: [][]int{
			[]int{94911152, 94911151},
			[]int{0, 0},
			[]int{94911151, 94911150},
		}}, want: 2},
	}
	for _, t := range tests {
		res := maxPoints(t.args.points)
		if res != t.want {
			fmt.Println("want ", t.name, "but ", res)
		}
	}
}

/*
   给定一个二维平面，平面上有 n 个点，求最多有多少个点在同一条直线上。
   示例 1:
   输入: [[1,1],[2,2],[3,3]]
   输出: 3
   解释:
   ^
   |
   |        o
   |     o
   |  o
   +------------->
   0  1  2  3  4

   示例 2:

   输入: [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
   输出: 4
   解释:
   ^
   |
   |  o
   |     o        o
   |        o
   |  o        o
   +------------------->
   0  1  2  3  4  5  6
*/
func maxPoints(points [][]int) int {
	if len(points) <= 1 {
		return len(points)
	}
	res := math.MinInt64
	resMap := make(map[float32]int)
	for i := 0; i < len(points); i++ {
		resMap[float32(i)] = 1
		thisMax := 1
		for j := i + 1; j < len(points); j++ {
			slope := float32(points[i][1]-points[j][1]) / float32(points[i][0]-points[j][1])
			if points[i][1] == points[j][1] || points[i][0] == points[j][1] {
				resMap[math.MaxFloat32] += 1
				if resMap[math.MaxFloat32] > thisMax {
					thisMax = resMap[math.MaxFloat32]
				}
				continue
			}

		}
	}
}
