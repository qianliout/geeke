package main

import (
	"fmt"
)

func main() {
	res := superPow(2147483647, []int{2, 0, 0})
	//res := superPow(2, []int{2})
	fmt.Println("res is ", res)
}

/*
你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。
示例 1:
输入: a = 2, b = [3]
输出: 8
示例 2:
输入: a = 2, b = [1,0]
输出: 1024
*/
// 这道题的主要考点就在溢出

func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	base := 1337
	if len(b) == 1 {
		return pow(a, b[0], base)
	}

	last := b[len(b)-1]
	res := (pow(a, last, base)) * pow(superPow(a, b[:len(b)-1]), 10, base) % base
	return res
}

func pow2(a, b, base int) int {
	if b == 0 {
		return 1
	}

	res := 1
	for i := 0; i < b; i++ {
		res = res * (a % base)
		res = res % base
	}
	return res
}

func pow(a, b, base int) int {
	if b == 0 {
		return 1
	}
	a = a % base

	if b%2 == 1 {
		return (a * pow(a, b-1, base)) % base
	} else {
		sub := pow(a, b>>1, base)
		return (sub * sub) % base
	}
}
