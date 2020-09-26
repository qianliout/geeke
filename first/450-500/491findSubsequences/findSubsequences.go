package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1}
	// nums := []int{4, 6, 7, 7}
	// nums:=[]int{1,-1}
	res := findSubsequences(nums)
	fmt.Println("res is ", res)
}

/*
给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
示例:
输入: [4, 6, 7, 7]
输出: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
说明:
    给定数组的长度不会超过15。
    数组中的整数范围是 [-100,100]。
    给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况 [1,1],[1,1,1],[1,1,1]就当做一种情况
*/

// 这种去重就麻烦：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1,因为前面的一个1，把后面的所有的1都算了，所以后面再算就会重复了
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	dfs(nums, path, 0, &res)
	return res

}

func dfs(nums []int, path []int, start int, res *[][]int) {
	if len(path) >= 2 {
		*res = append(*res, append([]int{}, path...))
	}
	// 每次递归新开去重的map，这里是一个新的知识点
	used := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		// 画出决策树，去重是关键
		if used[nums[i]] {
			continue
		}

		used[nums[i]] = true

		if len(path) == 0 || (len(path) > 0 && path[len(path)-1] <= nums[i]) {
			path = append(path, nums[i])
			dfs(nums, path, i+1, res)
			path = path[:len(path)-1]
			// (*used)[nums[i]] = false
		}
	}
}

/*
unc findSubsequences(nums []int) [][]int {
    res := make([][]int, 0)
	path := make([]int, 0)

	dfs(nums, path, -1, &res)
	return res

}

func dfs(nums []int, path []int, start int, res *[][]int) {
	if len(path) >= 2 {
		*res = append(*res, append([]int{}, path...))
	}
	used := make(map[int]bool)
	for i := start+1; i < len(nums); i++ {
		// 画出决策树，去重是关键
		if used[nums[i]] {
			continue
		}

		used[nums[i]] = true

		if start == -1 || (nums[start] <= nums[i]) {
			path = append(path, nums[i])
			dfs(nums, path, i, res)
			path = path[:len(path)-1]
			// (*used)[nums[i]] = false
		}
	}
}
*/
