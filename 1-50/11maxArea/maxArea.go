package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := maxArea(nums)
	fmt.Println(res)
}

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
(i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。
图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/
// 这种题典型的双指针
func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	start, end := 0, len(height)-1
	res := math.MinInt64
	for start < end {
		v := miniInt(height[start], height[end]) * (end - start)
		if v > res {
			res = v
		}
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return res
}

func miniInt(nums ...int) int {
	res := math.MaxInt64
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}
