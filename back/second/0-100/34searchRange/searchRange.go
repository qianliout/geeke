package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 7, 7, 8, 8, 10}
	res := searchRange(nums, 8)
	fmt.Println("res is ", res)
}

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。
示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1

	// 寻找左测
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid
		}
	}

	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	start = left

	left, right = 0, len(nums)
	// 寻找右测
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}

	if left-1 >= len(nums) || nums[left-1] != target {
		return []int{-1. - 1}
	}
	end = left - 1
	return []int{start, end}
}

func searchRange2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start, end := -1, -1
	left, right := 0, len(nums)-1
	// 先找左
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	start = left

	// 再找右
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left-1 >= len(nums) || nums[left-1] != target {
		return []int{-1, -1}
	}
	end = left - 1
	return []int{start, end}

}
