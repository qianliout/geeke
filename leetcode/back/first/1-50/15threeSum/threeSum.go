package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 0, 0, 0}
	res := threeSum(nums)
	fmt.Println("res is ", res)

}

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func threeSum(nums []int) [][]int {
	return Sum(nums, 0)
}

func Sum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > target {
			break
		}
		if i >= 1 && nums[i] == nums[i-1] {
			continue //防止重复,还是不能防止重复,这种情况下就还是重复的，{0,0,0,0},这种情况去重会很麻烦，
		}
		exitMap := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			_, exit := exitMap[nums[j]]
			if !exit {
				exitMap[target-nums[i]-nums[j]] = 1
			} else {
				res = append(res, []int{nums[i], nums[j], target - nums[i] - nums[j]})
			}
		}
	}
	return res
}

func Sum2(nums []int, target int) [][]int {
	res := make([][]int, 0)

	if len(nums) <= 0 {
		return res
	}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			s := nums[i] + nums[left] + nums[right]
			if s > target {
				right--
			} else if s < target {
				left++
			} else if s == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}
	return res
}
