package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 3, 5, 6}
	// tart := 8
	// res := searchInsert(nums, tart)
	// fmt.Println("res is ", res)
	fmt.Println(len([]rune("你好你在那里")))

}

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。
示例 1:
输入: [1,3,5,6], 5
输出: 2
示例 2:
输入: [1,3,5,6], 2
输出: 1
示例 3:
输入: [1,3,5,6], 7
输出: 4
示例 4:
输入: [1,3,5,6], 0
输出: 0
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right || (left <= len(nums)-2 && nums[left] <= target && nums[left+1] > target) {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	// 因为如果越界，刚好就是插入的位置
	return left
}

func searchInsert2(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	lenght := len(nums)
	if nums[lenght-1] < target {
		return lenght
	}
	left, right := 0, lenght-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			//  因为题目中说没有重复元素
			return mid
		} else if nums[mid] < target {
			//  严格小于target,一定不是解,区间是[mid+1,rigth]
			left = mid + 1   // 这种方式下,向下取整也没有问题
		} else {
			//  此时区间是 [left,mid]
			right = mid
		}
	}
	//  检查最后一个元素,如果最后一个都不符合, 则插入到right的位置,此时left刚才就是lenght
	return left
}

func searchInsert3(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	lenght := len(nums)

	left, right := 0, lenght
	//  上面没有做最后一个的验证,所以要把边界扩到最后面
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			//  因为题目中说没有重复元素
			return mid
		} else if nums[mid] < target {
			//  严格小于target,一定不是解,区间是[mid+1,rigth]
			left = mid + 1   // 这种方式下,向下取整也没有问题
		} else {
			//  此时区间是 [left,mid]
			right = mid
		}
	}
	//  检查最后一个元素,如果最后一个都不符合, 则插入到right的位置,此时left刚才就是lenght
	return left
}


func searchInsert3(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	lenght := len(nums)

	left, right := 0, lenght
	//  上面没有做最后一个的验证,所以要把边界扩到最后面
	for left < right {
		mid := left + (right-left)/2
		//  也可以不做==的判断
		if nums[mid] < target {
			//  严格小于target,一定不是解,区间是[mid+1,rigth]
			left = mid + 1   // 这种方式下,向下取整也没有问题
		} else {
			//  此时区间是 [left,mid]
			right = mid
		}
	}
	//  检查最后一个元素,如果最后一个都不符合, 则插入到right的位置,此时left刚才就是lenght
	return left
}