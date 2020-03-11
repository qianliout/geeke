package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{10, 10, 1, 10, 10, 10}
	res := findMin2(nums)
	fmt.Println("res is ", res)
}

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
请找出其中最小的元素。
注意数组中可能存在重复的元素。
示例 1：
输入: [1,3,5]
输出: 1
示例 2：
输入: [2,2,2,0,1]
输出: 0
说明：
    这道题是 寻找旋转排序数组中的最小值 的延伸题目。
    允许重复会影响算法的时间复杂度吗？会如何影响，为什么？
*/
//[]int{10, 1, 10, 10, 10}这种情况是种特例，要特别考虑
func findMin(nums []int) int {
	if len(nums) == 0 {
		return math.MaxInt64
	}
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	mid := left + (right-left)/2

	if nums[left] < nums[mid] {
		return minInt(nums[left], findMin(nums[mid+1:]))
	} else if nums[mid] < nums[right] {
		return minInt(findMin(nums[:mid]), nums[mid])
	} else if nums[mid] == nums[right] {
		return minInt(findMin(nums[:right])) // 这里注意，不能使用right-1，因为是右开区间，
	} else if nums[mid] == nums[left] && mid < len(nums)-1 {
		return minInt(findMin(nums[left+1:]))
	}

	return nums[mid]
}

func findMin2(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

func minInt(nums ...int) int {
	var mini int = math.MaxInt64
	for _, i2 := range nums {
		if i2 < mini {
			mini = i2
		}
	}
	return mini
}
