package main

import (
	"fmt"
)

func main() {
	ans := countArrangement(10)
	fmt.Println("ans is ", ans)
}

/*
假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N) 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：
    第 i 位的数字能被 i 整除
    i 能被第 i 位上的数字整除
现在给定一个整数 N，请问可以构造多少个优美的排列？
示例1:
输入: 2
输出: 2
解释:
第 1 个优美的排列是 [1, 2]:
  第 1 个位置（i=1）上的数字是1，1能被 i（i=1）整除
  第 2 个位置（i=2）上的数字是2，2能被 i（i=2）整除
第 2 个优美的排列是 [2, 1]:
  第 1 个位置（i=1）上的数字是2，2能被 i（i=1）整除
  第 2 个位置（i=2）上的数字是1，i（i=2）能被 1 整除
说明:
    N 是一个正整数，并且不会超过15。
*/

func countArrangement(N int) int {
	// res := make([][]int, 0)
	path := make([]int, 0)
	// for i := 1; i <= N; i++ {
	used := make(map[int]bool)
	res := 0
	dfs(N, path, &used, &res)
	// }
	// fmt.Println("res is ", res)
	return res
}

func dfs(n int, path []int, usd *map[int]bool, res *int) {
	if len(path) == n {
		*res = *res + 1
		// *res = append(*res, append([]int{}, path...))
		return
	}

	for j := 1; j <= n; j++ {
		if !(*usd)[j] {
			if (len(path)+1)%j == 0 || j%(len(path)+1) == 0 { // 这一步剪枝很重要，不然就超时
				path = append(path, j)
				(*usd)[j] = true
				dfs(n, path, usd, res)
				(*usd)[j] = false
				path = path[:len(path)-1]
			}
		}
	}
}

func check(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	for i, n := range nums {

		if n%(i+1) == 0 {
			continue
		}
		if (i+1)%n == 0 {
			continue
		}
		return false
	}
	return true
}
