package main

import (
	"fmt"
)

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	// heights := []int{4, 3, 5, 5, 9, 2, 8, 4, 7, 2, 3, 8, 3, 5, 4, 7, 9} // 34

	// heights := []int{2, 1, 2}
	// heights := []int{2, 0, 2}
	// heights := []int{4, 2, 0, 3, 2, 5}
	res := largestRectangleArea(heights)
	fmt.Println("res is ", res)
}

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。
以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
示例:
输入: [2,1,5,6,2,3]
输出: 10
*/

// 这道题有各种方法,但是最好的还是用单调栈的方法
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	// 这种方式,对go语言来说,内存开销大,但是前后加零,确实很巧妙,哨兵技巧
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	n := len(heights)
	maxArea := 0
	stark := make([]int, 0)
	for i := 0; i < n; i++ {
		// 出栈,并计最值
		for len(stark) > 0 && heights[stark[len(stark)-1]] > heights[i] {
			top := stark[len(stark)-1] // 这个元素就是当前可以确定元素的高度
			// 这里要注意，要先弹栈,把这个元素弹出后，stark前一个元素就表示当前这个高度可以向左扩张的地方
			stark = stark[:len(stark)-1]
			right := i - 1
			left := stark[len(stark)-1] + 1
			maxArea = Max(maxArea, (right-left+1)*heights[top])
		}
		stark = append(stark, i)
	}
	return maxArea
}

func largere(heights []int) int {
	stark := make([]int, 0)
	stark = append(stark, -1)
	maxArea := 0
	for i, _ := range heights {
		for len(stark) > 1 && heights[i] <= heights[stark[len(stark)-1]] {
			pop := stark[len(stark)-1]
			// 这里要特别注意,一定是先pop一个出来,再计算宽度,不然就会出错,
			stark = stark[:len(stark)-1]
			res := heights[pop] * (i - stark[len(stark)-1] - 1)
			if res > maxArea {
				maxArea = res
			}
		}
		stark = append(stark, i)
		// fmt.Println(stark, maxArea)
	}

	for len(stark) > 1 {
		pop := stark[len(stark)-1]
		stark = stark[:len(stark)-1]
		res := heights[pop] * (len(heights) - stark[len(stark)-1] - 1)
		if res > maxArea {
			maxArea = res
		}
	}
	return maxArea
}
