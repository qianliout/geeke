package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	sub := subsetsWithDup(nums)
	fmt.Println("sub is ", sub)
}

/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
示例 1：
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	path, res := make([]int, 0), make([][]int, 0)
	dfs(nums, path, 0, &res)
	return res
}

func dfs(nums []int, path []int, start int, res *[][]int) {
	*res = append(*res, append([]int{}, path...))
	for i := start; i < len(nums); i++ {
		// 剪同一层的方法
		if i > start && i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, path, i+1, res)
		path = path[:len(path)-1]
	}
}
