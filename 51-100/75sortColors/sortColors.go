package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	// nums := []int{2}
	sortColors(nums)
	fmt.Println(nums)
}

/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
必须在不使用库的sort函数的情况下解决这个问题。
*/
func sortColors(nums []int) {
	red, blue, i := 0, len(nums)-1, 0
	for i < blue {
		if nums[i] == 2 {
			nums[i], nums[blue] = nums[blue], nums[i]
			blue--
		} else if nums[i] == 0 {
			nums[i], nums[red] = nums[red], nums[i]
			red++
			i++ // 这一步是这个题目的关键点
		} else {
			i++
		}
	}
}
