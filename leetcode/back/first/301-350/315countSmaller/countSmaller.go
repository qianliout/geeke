package main

import (
	"fmt"

	"outback/leetcode/back/common"
)

func main() {
	nums := []int{5, 2, 6, 1}
	res := countSmaller(nums)
	fmt.Println("res is ", res)

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
	res := make([]int, len(nums))
	if len(nums) == 0 {
		return res
	}
	small := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		r := common.FindSmallIdxAndInsert(&small, nums[i])
		res[i] = r
	}

	return res
}
