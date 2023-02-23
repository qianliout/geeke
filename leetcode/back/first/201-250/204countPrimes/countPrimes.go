package main

import (
	"fmt"
)

func main() {
	res := countPrimes(3)
	fmt.Println("res is ", res)
}

/*
统计所有小于非负整数 n 的质数的数量。
示例:
输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	primeMap := make(map[int]bool)
	for i := 2; i < n; i++ {
		// 所有的,先默认是质数
		primeMap[i] = true
	}

	for i := 2; i*i < n; i++ {
		if primeMap[i] {
			for j := 2; j*i < n; j++ {
				primeMap[j*i] = false
			}
		}
	}
	count := 0
	for _, i := range primeMap {
		if i {
			count++
		}
	}
	return count
}
