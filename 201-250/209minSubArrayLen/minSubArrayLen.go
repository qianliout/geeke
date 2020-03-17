package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 4, 4}
	s := 1
	res := minSubArrayLen2(s, nums)
	fmt.Println("res is ", res)
}

/*
给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组。如果不存在符合条件的连续子数组，返回 0。
示例:
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
进阶:
如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
*/
func minSubArrayLen(s int, nums []int) int {
	if len(nums) <= 0 || s <= 0 {
		return 0
	}
	ans := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		thisSum := 0
		thisAns := 0
		for j := i; j >= 0; j-- {
			thisSum += nums[j]
			thisAns++
			if thisSum >= s {
				if thisAns < ans {
					ans = thisAns
				}
				break
			}
		}
	}
	// 都没有找到
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// 双指针，也是滑动窗口的方法
func minSubArrayLen2(s int, nums []int) int {
	if len(nums) <= 0 || s <= 0 {
		return 0
	}
	left, right, sum, ans := 0, 0, 0, math.MaxInt64
	for right < len(nums) {
		sum += nums[right]
		right++

		for sum >= s {
			if (right - left) < ans {
				ans = right - left
			}
			sum -= nums[left]
			left++
			// 如果放在left++ 后面，就要加1，注意
			//if (right - left + 1) < ans {
			//	ans = right - left + 1
			//}
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
