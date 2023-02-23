package main

import (
	"fmt"
	"math"
	"sort"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	closest := threeSumClosest(nums, 1)
	fmt.Println("closest is ", closest)
}

/*
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。
示例 1：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/
func threeSumClosest(nums []int, target int) int {
	closet, ans := math.MaxInt64, math.MaxInt64

	sort.Ints(nums)
	for i := range nums {
		left, right := i+1, len(nums)-1
		for left < right {
			tag := nums[i] + nums[left] + nums[right]
			if tag > target {
				right--
			} else if tag < target {
				left++
			} else {
				return tag
			}
			if AbsSub(tag, target) < closet {
				closet = AbsSub(tag, target)
				ans = tag
			}
		}
	}
	return ans
}
