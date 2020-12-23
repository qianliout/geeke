package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindrome("cbbcc"))
}

/*
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
示例 1:
输入: "aba"
输出: True
示例 2:
输入: "abca"
输出: True
解释: 你可以删除c字符。
*/
func validPalindrome(s string) bool {
	ss := []byte(s)
	for i := 0; i <= len(s)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			// fmt.Println(i, string(ss[:i])+string(ss[i+1:]), string(ss[:len(ss)-1-i])+string(ss[len(ss)-1-i+1:]))
			return valid(string(ss[:i])+string(ss[i+1:])) ||
				valid(string(ss[:len(ss)-1-i])+string(ss[len(ss)-1-i+1:]))
		}
	}
	return true
}

func valid(s string) bool {
	ss := []byte(s)
	for i := 0; i <= len(ss)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			return false
		}
	}
	return true
}
