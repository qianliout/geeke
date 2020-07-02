package main

import (
	"fmt"
	"sort"
)

func main() {
	//nums := []int{1, 1, 2, 1, 2, 2, 1}
	//nums := []int{1, 5, 1, 1, 6, 4}
	nums := []int{1, 3, 2, 2, 3, 1}
	wiggleSort2(nums)
}

/*
给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。
示例 1:
输入: nums = [1, 5, 1, 1, 6, 4]
输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]
示例 2:
输入: nums = [1, 3, 2, 2, 3, 1]
输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]
说明:
你可以假设所有输入都会得到有效的结果。
进阶:
你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
链接：https://leetcode-cn.com/problems/wiggle-sort-ii
*/

// 方法一排序，这里面有很个易错点
func wiggleSort(nums []int) {
	sort.Ints(nums)
	tem := make([]int, len(nums))
	copy(tem, nums)
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		mid = (len(nums) - 1) / 2
	}
	// 这里是注意点，把数据反过来重新组,因为在极端情况下可能会出错
	//https://leetcode-cn.com/problems/wiggle-sort-ii/solution/yi-bu-yi-bu-jiang-shi-jian-fu-za-du-cong-onlognjia/
	j := 0
	n := len(nums)
	for i := mid; i >= 0; i-- {
		nums[j] = tem[i]
		j++
		if n-(mid-i)-1 > mid {
			nums[j] = tem[n-(mid-i)-1]
			j++
		}

	}
	fmt.Println("num is ", nums)
}

func wiggleSort2(nums []int) {
	left, right, mid := 0, len(nums)-1, (len(nums)-1)/2
	// 找中位数，并三分
	for mid < right {
		if nums[left] > nums[mid] {
			nums[mid], nums[left] = nums[left], nums[mid]
		}
		left++

		if nums[right] <= nums[mid] {
			nums[right], nums[mid] = nums[mid], nums[right]
		}
		right--
		fmt.Println(left, right, nums)
	}

	// 组新数組
	tem := make([]int, len(nums))
	copy(tem, nums)
	fmt.Println(nums)
	fmt.Println(tem)

	mid = len(nums) / 2
	if len(nums)%2 == 0 {
		mid = (len(nums) - 1) / 2
	}
	// 这里是注意点，把数据反过来重新组,因为在极端情况下可能会出错
	//https://leetcode-cn.com/problems/wiggle-sort-ii/solution/yi-bu-yi-bu-jiang-shi-jian-fu-za-du-cong-onlognjia/
	j := 0
	n := len(nums)
	for i := mid; i >= 0; i-- {
		nums[j] = tem[i]
		j++
		if n-(mid-i)-1 > mid {
			nums[j] = tem[n-(mid-i)-1]
			j++
		}

	}
	fmt.Println("num is ", nums)
}
