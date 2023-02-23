package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 3}
	res := nextGreaterElements(nums)
	fmt.Println("res is ", res)
}

/*
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，
这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。
示例 1:
输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
注意: 输入数组的长度不会超过 10000。
*/
// 单调栈解法，技巧是把数据扩大两倍，但是不用真正的扩容
func nextGreaterElements(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	stark := make([]int, 0)
	for i := 2*len(nums) - 1; i >= 0; i-- {
		j := i
		if i >= len(nums) {
			j = i - len(nums)
		}

		for len(stark) > 0 && stark[len(stark)-1] <= nums[j] {
			stark = stark[:len(stark)-1]
		}

		if i < len(nums) {
			if len(stark) == 0 {
				res = append(res, -1)
			} else {
				res = append(res, stark[len(stark)-1])
			}
		}
		stark = append(stark, nums[j])
	}
	// res的值是反的，这里要反转
	ans := make([]int, 0)
	for i := len(res) - 1; i >= 0; i-- {
		ans = append(ans, res[i])
	}

	return ans
}
