package main

import (
	"sort"

	"qianliout/leetcode/common/utils"
)

func main() {

}

/*
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
示例 1:
输入: [1,2,3]
输出: 6
示例 2:
输入: [1,2,3,4]
输出: 24
*/
func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	return utils.Max(nums[0]*nums[1]*nums[len(nums)-1], nums[len(nums)-1]*nums[len(nums)-2]*nums[len(nums)-3])
}
