package main

import (
	"fmt"
)

func main() {
	res := isPalindrome(1111111111111111112)
	fmt.Println(res)
}

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == reversion(x)
}

func reversion(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x = x / 10
	}
	return ans
}
