package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 5
	res := search1(nums, target)
	fmt.Println("res is ", res)
}

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。
编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。
示例 1:
输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
示例 2:
输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false
进阶:
    这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
    这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
*/
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1

	mid := left + (right-left)/2
	if nums[mid] == target || nums[left] == target || nums[right] == target {
		return true
	}
	if nums[mid] > nums[right] {
		if nums[mid] > target {
			return search(nums[left:mid], target) || search(nums[mid+1:], target)
		} else if nums[mid] < target {
			return search(nums[mid+1:], target)
		}
	} else if nums[mid] < nums[right] {
		if nums[mid] > target {
			return search(nums[:mid], target)
		} else if nums[mid] < target {
			return search(nums[mid+1:], target) || search(nums[left:mid], target)
		}
	} else if nums[mid] == nums[right] {
		return search(nums[left:right], target)
	}
	return false
}

func search1(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target || nums[left] == target || nums[right] == target {
			return true
		} else if nums[mid] == nums[right] {
			right--
			continue
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] > target {
				left = mid + 1
			} else {
				right = mid - 1 // 注意的是，这里right = mid 或者right = mid 提交都是对的
			}
		} else if nums[mid] > nums[right] {
			if nums[left] < target && nums[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return false
}
