package main

import (
	"fmt"
	"math"
	"sort"

	"qianliout/leetcode/back/common"
)

func main() {
	ans := threeSumClosest([]int{-1, 2, 1, -4, 3}, 3)
	fmt.Println(ans)
}

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与
target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans, sub := math.MaxInt64, math.MaxInt64

	for i := 0; i <= len(nums)-3; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return sum
			}

			if sub > common.AbsSub(target, sum) {
				sub = common.AbsSub(target, sum)
				ans = sum
			}
		}
	}
	return ans
}
