package main

import (
	"fmt"
)

func main() {
	res := nextGreaterElements([]int{1, 2, 1})
	fmt.Println(res)
}

/*
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的
下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个
更大的数。如果不存在，则输出 -1。
示例 1:
输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
*/

// 反向找，
func nextGreaterElements(nums []int) []int {
	// 减
	stark := make([]int, 0)
	ans := make([]int, 0)
	for i := len(nums)*2 - 1; i >= 0; i-- {
		j := i % len(nums)
		
		for len(stark) > 0 && stark[len(stark)-1] <= nums[j] {
			stark = stark[:len(stark)-1]
		}
		
		if i < len(nums) {
			if len(stark) == 0 {
				ans = append(ans, -1)
			} else {
				ans = append(ans, stark[len(stark)-1])
			}
		}
		stark = append(stark, nums[j])
	}
	
	return ans
}
