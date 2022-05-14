package main

import (
	"fmt"
)

func main() {
	nums := []int{11, 13, 15, 17}
	fmt.Println(findMin(nums))
}

/*
153. 寻找旋转排序数组中的最小值
给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = left + 1
			// } else if nums[mid] < nums[left] { // 两种写法都对
		} else if nums[mid] < nums[right] {
			right = mid
		} else { // mid==right
			return nums[left]
		}
	}
	return nums[left]
}
