package main

import (
	"fmt"
)

func main() {
	permutation := getPermutation(3, 3)
	permutation2 := getPermutation2(9, 78494)
	fmt.Println("permutation ", permutation, permutation2)
}

/*
给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
"123"
"132"
"213"
"231"
"312"
"321"
*/

func getPermutation(n int, k int) string {
	used, path, res := make(map[int]bool), make([]int, 0), make([]int, 0)
	var pre int

	all := Factorial(n - 1)
	start, low := k/all, k%all
	if low > 0 {
		start = start + 1
	}
	if low == 0 {
		low = all
	}
	// fmt.Println("all ", all, start, low)

	used[start] = true
	pre = low
	path = append(path, start)

	dfs(n, used, &pre, path, &res)
	ans := ""
	for i := range res {
		ans = fmt.Sprintf("%s%d", ans, res[i])
	}
	return ans
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func dfs(n int, used map[int]bool, pre *int, path []int, res *[]int) {
	if len(path) == n {
		*pre = *pre - 1
		if *pre == 0 {
			*res = append(*res, path...)
			return
		}
		return
	}
	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, i)
		dfs(n, used, pre, path, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func dfs2(n, k int, used map[int]bool, pre *int, path []int, res *[]int) {
	if len(path) == n {
		*pre = *pre + 1
		if *pre == k {
			*res = append(*res, path...)
			return
		}
		return
	}
	for i := 1; i <= n; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		path = append(path, i)
		dfs2(n, k, used, pre, path, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}

// 这种方式可以执行，但是会超时
func getPermutation2(n int, k int) string {
	used, path, res := make(map[int]bool), make([]int, 0), make([]int, 0)
	var pre int

	dfs2(n, k, used, &pre, path, &res)
	ans := ""
	for i := range res {
		ans = fmt.Sprintf("%s%d", ans, res[i])
	}
	return ans
}
