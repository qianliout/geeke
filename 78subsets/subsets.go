package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sets := subsets(nums)
	fmt.Println("sets is ", sets)
}

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
*/
func subsets(nums []int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)

	dsf(&nums, 0, path, &res)
	return res
}

func dsf(nums *[]int, start int, path []int, res *[][]int) {
	if len(path) > len(*nums) {
		return
	}
	*res = append(*res, append([]int{}, path...))

	for i := start; i < len(*nums); i++ {
		path = append(path, (*nums)[i])
		dsf(nums, i+1, path, res)
		path = path[:len(path)-1]
	}
}
