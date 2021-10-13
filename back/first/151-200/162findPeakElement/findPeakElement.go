package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 2, 1}
	res := findPeakElement(nums)
	fmt.Println("res is ", res)
}

/*
峰值元素是指其值大于左右相邻值的元素。
给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞。
示例 1:
输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。
示例 2:
输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
*/
func findPeakElement(nums []int) int {
	//return find1(nums)
	return find2(nums, 0, len(nums)-1)
	//return find1(nums)
}

func find1(nums []int) int {
	nums = append([]int{math.MinInt64}, nums...)
	nums = append(nums, math.MinInt64)

	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i - 1
		}
	}
	return -1
}

// 因为返回任意解，所以可以用递归的方式，如果返回全部解，就要想另一种方法了
func find2(nums []int, left, right int) int {

	if left == right {
		return left
	}
	mid := left + (right-right)/2

	if nums[mid] > nums[mid+1] {
		return find2(nums, left, mid)
	} else {
		return find2(nums, mid+1, right)
	}
}

// 因为返回一个解就行,这里为什么会对,有些不懂呢,题目中不是说大于左右相临的元素吗
func find3(nums []int) int {
	for i := 0; i <= len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

func find4(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	for i := 0; i <= len(nums)-1; i++ {
		if i == 0 {
			if nums[i] > nums[i+1] {
				return i
			}
		} else if i == len(nums)-1 {
			if nums[i] > nums[i-1] {
				return i
			}
		} else {
			if nums[i] > nums[i+1] && nums[i-1] < nums[i] {
				return i
			}
		}
	}
	return len(nums) - 1
}
