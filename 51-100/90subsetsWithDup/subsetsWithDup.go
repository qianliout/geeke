package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	res := subsetsWithDup(nums)
	fmt.Println("res is ", res)
}

/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
示例:
输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	used := make(map[int]bool)
	path := make([]int, 0)
	dfs(nums, used, 0, path, &res)
	res = append(res, []int{})
	return res
}

func dfs(nums []int, used map[int]bool, index int, path []int, res *[][]int) {
	if len(path) > 0 {
		*res = append(*res, append([]int{}, path...))
		if len(path) == len(nums) {
			return
		}
	}
	for i := index; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums, used, i+1, path, res)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
