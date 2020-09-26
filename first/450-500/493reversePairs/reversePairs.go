package main

import (
	"fmt"
)

func main() {
	nums := []int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}
	res := reversePairs(nums)
	fmt.Println("res is ", res)
}

/*
给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
你需要返回给定数组中的重要翻转对的数量。
示例 1:
输入: [1,3,2,3,1]
输出: 2
示例 2:
输入: [2,4,3,5,1]
输出: 3
注意:
    给定数组的长度不会超过50000。
    输入数组中的所有数字都在32位整数的表示范围内。
*/
func reversePairs(nums []int) int {
	// return Violence(nums)
	return Violence(nums)
}

// 首先暴力法,可以加一下些mem技术，但是当没有重复数据且数据量很大时，仍然会超时
func Violence(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > 2*nums[j] {
				ans++
			}
		}
	}
	return ans
}

func useDP(nums []int) int {
	return 0
}
