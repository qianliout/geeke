package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{1, 1, 2}
	res := permuteUnique(num)
	fmt.Println("res is ", res)
}

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。
示例:
输入: [1,1,2]
输出:
[

	[1,1,2],
	[1,2,1],
	[2,1,1]

]
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := permuteUnique(append(append([]int{}, nums[:i]...), nums[i+1:]...))
		for _, l := range left {
			res = append(res, append([]int{nums[i]}, l...))
		}
	}

	return res
}
