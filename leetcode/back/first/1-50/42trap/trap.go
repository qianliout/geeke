package main

import (
	"fmt"
)

func main() {

	// nums := []int{0, 3, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	nums := []int{4, 2, 3}
	res := trap(nums)
	fmt.Println("rain is ", res)
}

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
感谢 Marcos 贡献此图。
示例:
输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/
// 单调递减栈
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	stark := make([]int, 0)

	// 单调整递减栈表可能接到雨水的柱子
	stark = append(stark, -1)
	rain := 0
	for i := 0; i < n; i++ {
		// 一定是大于,比栈顶的原素都大,那说明可以计算之前可以接的水了
		for len(stark) > 1 && height[i] > height[stark[len(stark)-1]] {
			// 说明这一个加进去之后就会有比之前高的了，就可以装雨水了，所以得计算当前层能装的水
			pop := stark[len(stark)-1]
			stark = stark[:len(stark)-1]
			rain += calculateRain(height, stark[len(stark)-1], i, pop)
		}
		stark = append(stark, i)
	}
	return rain
}

// 这里只记算当前层的装水量
func calculateRain(height []int, left, right, pop int) int {
	if left == -1 {
		return 0
	}
	if right-left <= 1 {
		return 0
	}
	min := height[left]
	if height[right] < min {
		min = height[right]
	}

	return (right - left - 1) * (min - height[pop])
}

// 使用双指针,更容易理解
func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left, right := 0, len(height)-1
	lMax, rMax := height[left], height[right]
	ans := 0
	for left < right {
		if lMax < rMax {
			ans += lMax - height[left]
			left++
			lMax = Max(height[left], lMax)
		} else {
			ans += rMax - height[right]
			right--
			rMax = Max(rMax, height[right])
		}
	}
	return ans
}
