package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 8
	res := search(nums, target)
	fmt.Println("res is ", res)
}

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。
示例 1:
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:
输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

// 二分法
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { //为什么要用等于。因为可能left==mid
			if nums[left] <= target && nums[mid] >= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] { // 右边升序
			if nums[mid] <= target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	if left > length-1 || nums[left] != target {
		return -1
	}
	return left
}
