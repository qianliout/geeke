package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3}
	res := permute(num)
	fmt.Println("res:", res)
}

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。
示例:
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for i := 0; i < len(nums); i++ {
		// 这里一定要使用这种方
		//left := permute(append(append([]int{}, nums[:i]...), nums[i+1:]...))
		left := permute(append(append([]int{}, nums[:i]...), nums[i+1:]...))
		for _, l := range left {
			res = append(res, append([]int{nums[i]}, l...))
		}
	}
	return res
}
