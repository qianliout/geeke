package main

import (
	"fmt"
	"sort"
)

/*
   给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
   输入：nums = [1,1,2]
   输出：
   [[1,1,2],
    [1,2,1],
    [2,1,1]]
*/

func main() {
	nums := []int{1, 1, 3}
	res := permuteUnique(nums)
	fmt.Println("res is ", res)
}

func permuteUnique(nums []int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	exit := make(map[int]bool)
	sort.Ints(nums)
	backtrack(nums, path, exit, &res)

	return res
}

func backtrack(nums []int, path []int, exit map[int]bool, res *[][]int) {
	if len(path) == len(nums) {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if exit[i] || (i > 0 && nums[i] == nums[i-1] && !exit[i-1]) {
			continue
		}

		exit[i] = true
		path = append(path, nums[i])
		backtrack(nums, path, exit, res)
		path = path[:len(path)-1]
		exit[i] = false
	}
}
