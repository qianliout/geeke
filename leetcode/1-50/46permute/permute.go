package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println("res is ", res)
}

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permute(nums []int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	exit := make(map[int]bool)
	backtrack(nums, path, exit, &res)
	return res
}

func backtrack(nums []int, path []int, exit map[int]bool, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if exit[i] {
			continue
		}
		exit[i] = true
		path = append(path, nums[i])
		backtrack(nums, path, exit, res)
		path = path[:len(path)-1]
		exit[i] = false
	}
}
