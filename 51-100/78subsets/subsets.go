package main

import (
	"fmt"
)

func main() {
	num := []int{9, 0, 3, 5, 7}
	//res := subsets(num)
	res := make([][]int, 0)
	used := make(map[int]bool)
	path := make([]int, 0)
	heller(num, 4, used, &res, path, 0)
	fmt.Println("res is ", res)

}

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func subsets(nums []int) [][]int {
	//return addNewsubsets(nums)
	return backtrack(nums)
}

func addNewsubsets(nums []int) [][]int {

	res := make([][]int, 0)

	res = append(res, []int{})
	if len(nums) == 0 {
		return res
	}

	for _, va := range nums {
		currentRes := make([][]int, 0)
		for _, v := range res {
			currentRes = append(currentRes, append(v, va))
		}
		res = append(res, currentRes...)
	}
	return res
}

func backtrack(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i <= len(nums); i++ {
		path := make([]int, 0)
		dfs(nums, path, i, &res)
	}
	return res
}

func dfs(nums, path []int, left int, res *[][]int) {
	if len(path) == left-1 {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := left; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, path, i+1, res)
		path = path[:len(path)-1]
	}
}

// 用回塑算法的话,去重是个太难题,所这里不用这种方法
func methodOfData(nums []int) {

}

func heller(nums []int, n int, used map[int]bool, res *[][]int, path []int, left int) {
	if left >= len(nums) {
		return
	}
	if len(path) == n {
		fmt.Println("path is ", path)
		*res = append(*res, append([]int{}, path...))
	}
	// 数组中不含有重复元素,所以,这里以什么去重都可以
	if !used[left] {
		path = append(path, nums[left])
		used[left] = true
		heller(nums, n, used, res, path, left+1)
		used[left] = false
		path = path[:len(path)-1]
	}
}
