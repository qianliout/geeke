package main

import (
	. "qianliout/leetcode/common/utils"
)

func main() {

}

func trap2(height []int) int {

	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0], rightMax[len(height)-1] = height[0], height[len(height)-1]
	// 左边的最大值
	for i := 1; i < len(height); i++ {
		leftMax[i] = Max(height[i], leftMax[i-1])
	}
	// 右边的最大值
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = Max(height[i], rightMax[i+1])
	}
	// 计算
	ans := 0
	for i := 1; i < len(height); i++ {
		ans = ans + Min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}

// 从trap2中可以看到，对于leftMax,rightMax，只会在i时使用一次，是可以动态更新的
func trap3(height []int) int {
	left, right, leftMax, rightMax := 0, len(height)-1, height[0], height[len(height)-1]

	ans := 0
	for left < right {
		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
			leftMax = Max(height[left], leftMax)
		} else {
			ans += rightMax - height[right]
			right--
			rightMax = Max(height[right], rightMax)
		}
	}

	return ans
}

func trap4(height []int) int {
	stak := make([]int, 0)
	ans := 0
	for i := 0; i < len(height); i++ {
		for len(stak) > 0 && height[stak[len(stak)-1]] < height[i] {

		}

		stak = append(stak, i)
	}
	return ans
}
