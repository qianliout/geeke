package main

import (
	"fmt"
)

func main() {
	res := combine(4, 2)
	fmt.Println("res is ", res)
}

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
示例:
输入: n = 4, k = 2
输出:
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
	res := make([][]int, 0)
	if n == 0 || k == 0 {
		return res
	}
	dfs(n, 1, k, []int{}, &res)
	return res
}

func dfs(n, start, k int, path []int, res *[][]int) {
	if len(path) == k {
		*res = append(*res, append([]int{}, path...))
		return
	}
	for i := start; i <= n; i++ {
		dfs(n, i+1, k, append(path, i), res)
	}
}
