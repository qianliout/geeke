package main

import (
	"fmt"

	. "qianliout/leetcode/common/utils"
)

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(nums)
	fmt.Println("area is ", area)
}

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
*/
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	// 哨兵技巧，这个要学会
	heights = append(append([]int{0}, heights...), 0)
	stark := make([]int, 0)
	ans := 0
	for i := 0; i < len(heights); i++ {
		// 单调递增
		for len(stark) > 0 && heights[stark[len(stark)-1]] > heights[i] {
			// 单调递增 说明：栈顶的一定是比之前没有入栈的元素都小,直到栈顶的前一个元素
			top := heights[stark[len(stark)-1]]

			stark = stark[:len(stark)-1]
			right := i - 1
			left := stark[len(stark)-1]
			ans = Max(ans, top*(right-left))
		}
		stark = append(stark, i)
	}
	return ans
}
