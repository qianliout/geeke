package main

import (
	"fmt"

	. "qianliout/leetcode/leetcode/common/utils"
)

func main() {
	nums := []int{1, 2}
	positive := firstMissingPositive(nums)
	fmt.Println("ans is ", positive)
}

/*
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
示例 1：
输入：nums = [1,2,0]
输出：3
*/
func firstMissingPositive(nums []int) int {
	hasOne := false
	for _, b := range nums {
		if b == 1 {
			hasOne = true
		}
	}
	if !hasOne {
		return 1
	}
	// 缺少的元素一定是1到len(nums)之间
	for i, b := range nums {
		if b <= 0 || b > len(nums) {
			nums[i] = 1
		}
	}
	for i := range nums {
		idx := Abs(nums[i]) - 1
		nums[idx] = -Abs(nums[idx])
	}
	for i, b := range nums {
		if b > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
