package main

import (
	"fmt"
)

func main() {
	res := isPowerOfThree(-9)
	fmt.Println("res is ", res)
}

/*
326. 3的幂
给定一个整数，写一个函数来判断它是否是 3 的幂次方。
示例 1:
输入: 27
输出: true
示例 2:
输入: 0
输出: false
示例 3:
输入: 9
输出: true
示例 4:
输入: 45
输出: false
*/
func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
