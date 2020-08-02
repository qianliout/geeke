package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPowerOfTwo(-8))
}

/*
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
示例 1:
输入: 1
输出: true
解释: 20 = 1
示例 2:
输入: 16
输出: true
解释: 24 = 16
链接：https://leetcode-cn.com/problems/power-of-two
*/
func isPowerOfTwo(n int) bool {
	return find3(n)
}
func find1(n int) bool {
	if n == 0 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}
	// 如果是负数，如-8这样的数，n==-1
	return n == 1
}
func find2(n int) bool {
	if n == 0 {
		return false
	}
	return n&(-n) == n
}

func find3(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0

}
