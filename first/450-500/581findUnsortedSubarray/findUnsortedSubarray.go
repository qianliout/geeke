package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findUnsortedSubarray1([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println(findUnsortedSubarray1([]int{1, 2, 3, 4}))
	fmt.Println(findUnsortedSubarray1([]int{1, 2, 3, 3, 3}))
}

/*
给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
你找到的子数组应是最短的，请输出它的长度。
示例 1:
输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
说明 :
输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。
*/
// https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/solution/si-lu-qing-xi-ming-liao-kan-bu-dong-bu-cun-zai-de-/
func findUnsortedSubarray(nums []int) int {
	max, min := nums[0], nums[len(nums)-1]
	start, end := 0, -1

	// 从左到右维持最大值，寻找右边界end
	for i, n := range nums {
		if n < max {
			end = i
		} else {
			max = n
		}
	}

	// 从右到左维持最小值，寻找左边界begin
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > min {
			start = i
		} else {
			min = nums[i]
		}
	}
	return end - start + 1
}

func findUnsortedSubarray1(nums []int) int {
	// max, min := nums[0], nums[len(nums)-1]

	max, min := math.MinInt64, math.MaxInt64
	start, end := 0, -1 // 这里定义初值很重要
	// 从左到右维持最大值，寻找右边界end
	for i, n := range nums {
		if n >= max { // 这里一定是大于等于
			max = n
		} else {
			end = i
		}
	}

	// 从右到左维持最小值，寻找左边界begin
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= min { // 这里一定是小于等于
			min = nums[i]
		} else {
			start = i
		}
	}
	return end - start + 1
}
