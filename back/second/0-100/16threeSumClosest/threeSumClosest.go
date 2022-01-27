package main

import (
	"fmt"
	"math"
	"sort"

	"qianliout/leetcode/back/common"
)

func main() {
	num := []int{-1, 2, 1, -4}
	res := threeSumClosest(num, 1)
	fmt.Println("res is ", res)
}

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
提示：
    3 <= nums.length <= 10^3
    -10^3 <= nums[i] <= 10^3
    -10^4 <= target <= 10^4
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := 0
	closest := math.MaxInt64
	for i, n := range nums {
		left, right := i+1, len(nums)-1
		for left < right {
			c := common.AbsSub(n+nums[left]+nums[right], target)
			s := n + nums[left] + nums[right]
			if c < closest {
				closest = c
				ans = n + nums[left] + nums[right]
			}

			if s < target {
				left++
			} else if s > target {
				right--
			} else {
				break
			}
		}
	}
	return ans
}
