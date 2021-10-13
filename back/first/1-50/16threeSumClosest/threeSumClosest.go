package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4, 0}
	target := 1
	res := threeSumClosest(nums, target)
	fmt.Println("res is ", res)
}

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/
func threeSumClosest(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}

	sort.Ints(nums)

	allClose := math.MaxInt64

	ans := 0
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			c := int(math.Abs(float64(target - s)))
			if c < allClose {
				ans = s
				allClose = c
			}
			if s < target {
				left++
			} else if s > target {
				right--
			} else {
				break
			}
		}
		if ans == target {
			break
		}
	}
	return ans
}
