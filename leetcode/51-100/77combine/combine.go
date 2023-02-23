package main

import (
	"fmt"
)

func main() {
	path := combine(4, 2)
	fmt.Println("path is ", path)
}

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
示例 1：
输入：n = 4, k = 2
输出：
[

	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],

]
*/
func combine(n int, k int) [][]int {
	path := make([]int, 0)
	res := make([][]int, 0)
	dfs(0, n, k, path, &res)
	return res
}

func dfs(start, n, k int, path []int, res *[][]int) {
	if len(path) == k {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := start + 1; i <= n; i++ {
		path = append(path, i)
		dfs(i, n, k, path, res)
		path = path[:len(path)-1]
	}
}
