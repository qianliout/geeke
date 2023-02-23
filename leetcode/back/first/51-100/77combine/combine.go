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
// 回漱法
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n <= 0 || k <= 0 {
		return res
	}
	path := make([]int, 0)
	used := make(map[int]bool)
	dfs(path, 1, k, n, &res, used)
	return res
}

func dfs(path []int, start, k, n int, res *[][]int, used map[int]bool) {
	if len(path) == k {
		*res = append(*res, append([]int{}, path...))
		return
	}
	// 这道题可以不用used判断的原因是他有序的，而且取数总是取下一个，不会向回取，当然，加上
	// used 也没有错
	//if !used[start] {
	//used[start] = true
	for i := start; i <= n; i++ {
		path = append(path, i)
		dfs(path, i+1, k, n, res, used)
		path = path[:len(path)-1]
	}
	//used[start] = false
	//}

}
