package main

import (
	"fmt"
	. "outback/leetcode/common"
)

func main() {
	nums := []int{1}
	BitLeft(&nums, 1)
	BitLeft(&nums, 1)
	BitLeft(&nums, 2)
	BitLeft(&nums, 2)

}

/*
给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是
nums[i] 右侧小于 nums[i] 的元素的数量。
示例:
输入: [5,2,6,1]
输出: [2,1,1,0]
解释:
5 的右侧有 2 个更小的元素 (2 和 1).
2 的右侧仅有 1 个更小的元素 (1).
6 的右侧有 1 个更小的元素 (1).
1 的右侧有 0 个更小的元素.
*/
func countSmaller(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	small := make([]int, 0)
	m := make(map[int]int)
	for i := len(nums) - 1; i >= 0; i-- {
		r := bitsetLeft(&small, nums[i])
		m[i] = r
	}
	for i := range nums {
		res = append(res, m[i])
	}

	return res
}

//  使用二分查找,插入左边
func bitsetLeft(nums *[]int, tar int) int {
	if len(*nums) == 0 {
		*nums = append(*nums, tar)
		return 0
	}
	length := len(*nums)
	left, right := 0, length-1
	for left < right {
		mid := left + (right-left)/2
		if (*nums)[mid] < tar {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//  检查最后
	if (*nums)[left] == tar {

	}

	fmt.Println("num is ", *nums, left, tar)
	r := (*nums)[left:]
	*nums = append(append((*nums)[:left], tar), r...)

	return left
}
