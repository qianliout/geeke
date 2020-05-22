package main

import (
	"fmt"
)

func main() {
	nums := []int{5,7,7,8,8,10}
	target := 6
	res := searchRange2(nums, target)
	fmt.Println("res is ", res)

}

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
你的算法时间复杂度必须是 O(log n) 级别。
如果数组中不存在目标值，返回 [-1, -1]。
示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/

func searchRange(nums []int, target int) []int {
	left, right, start, end := -1, -1, 0, len(nums)-1

	if len(nums) == 0 {
		return []int{left, right}
	}

	// 先确定左边界
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			end = mid - 1
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}
	// 检查越界
	if start >= len(nums) || nums[start] != target {
		left = -1
	} else {
		left = start
	}
	// 再确定右边界
	start, end = 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		}
	}
	// 检查越界
	if end < 0 || nums[end] != target {
		right = -1
	} else {
		right = end
	}
	return []int{left, right}
}

func searchRange2(nums []int, target int) []int {
	left, right := -1, -1
	if len(nums) == 0 {
		return []int{left, right}
	}
	length := len(nums) - 1
	start, end := 0, length

	// 先确实左边界
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			// 下一轮的搜索区间是[mid+1,end]
			start = mid + 1
		} else {
			// 搜索区间是[start,mid]
			end = mid
		}
	}
	//  判断是否存在目标元素
	if nums[start] != target {
		return []int{-1, -1}
	}

	left = start

	// 先确实右边界
	start, end = 0, length
	for start < end {
		mid := start + (end-start+1)/2 // 因为下达的start=mid,所以要向上取整
		if nums[mid] > target {
			// 下一轮的搜索区间是 [start,mid-1]
			end = mid - 1
		} else {
			// 搜索区间是[mid,end]
			start = mid // 如果是这种,就一定得向上取整
		}
	}
	right = start
	return []int{left, right}
}
