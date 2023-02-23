package main

import (
	"fmt"
)

func main() {
	res := combinationSum3(3, 9)
	fmt.Println("res is ", res)
}

/*
找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
说明：
    所有数字都是正整数。
    解集不能包含重复的组合。
示例 1:
输

入: k = 3, n = 7
输出: [[1,2,4]]
示例 2:
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
*/

func combinationSum3(k int, n int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	used := make(map[int]bool)
	if k == 0 {
		return res
	}
	dfs(n, k, 1, used, path, &res)
	return res
}

func dfs(n, k, start int, used map[int]bool, path []int, res *[][]int) {
	if n == 0 && len(path) == k {
		*res = append(*res, append([]int{}, path...))
		return
	}
	if n < 0 {
		return
	}
	for i := start; i <= 9; i++ {
		if !used[i] {
			path = append(path, i)
			used[i] = true
			dfs(n-i, k, i, used, path, res)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}
