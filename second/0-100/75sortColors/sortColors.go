package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 0, 2}
	sortColors(nums)
	fmt.Println(nums)
}

/*
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
注意:
不能使用代码库中的排序函数来解决这道题。
示例:
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
*/

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	left, right := 0, len(nums)-1
	cur := 0
	for cur <= right {
		if nums[cur] == 2 {
			nums[cur], nums[right] = nums[right], nums[cur]
			right--
		} else if nums[cur] == 0 {
			nums[cur], nums[left] = nums[left], nums[cur]
			left++
			cur++
		} else {
			cur++
		}
	}
}
