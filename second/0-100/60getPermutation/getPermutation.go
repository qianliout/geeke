package main

import (
	"fmt"
)

func main() {
	getPermutation(4, 9)
}

/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。
按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。
说明：
给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:
输入: n = 3, k = 3
输出: "213"
示例 2:
输入: n = 4, k = 9
输出: "2314"
*/
func getPermutation(n int, k int) string {
	fam := make(map[int]int)

	res := make([]int, 0)
	for k > 0 {
		for i := n; i > 0; i-- {
			f := Factorial(i-1, &fam)
			if k > f {
				res = append(res, i)
				k -= f
			}
		}
	}
	fmt.Println("res is ", res)
	return ""
}

func Factorial(n int, fam *map[int]int) int {
	if n, ok := (*fam)[n]; ok {
		return n
	}
	v := fac(n)
	(*fam)[n] = v
	return v
}

func fac(n int) int {
	if n == 1 {
		return 1
	}
	return n * fac(n-1)
}
