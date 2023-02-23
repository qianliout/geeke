package main

import (
	"fmt"
)

func main() {
	a := longestPalindrome("eabcba")
	fmt.Println(a)
}

/*
给你一个字符串 s，找到 s 中最长的回文子串。
*/
func longestPalindrome(s string) string {
	ss := []byte(s)
	ans := ""
	for i := 0; i < len(ss); i++ {
		res := expand(ss, i, i+1)
		if len(res) > len(ans) {
			ans = res
		}
		res = expand(ss, i, i)
		if len(res) > len(ans) {
			ans = res
		}
	}
	return ans
}

func expand(ss []byte, left, right int) string {
	for left >= 0 && right < len(ss) && ss[left] == ss[right] {
		left--
		right++
	}
	return string(ss[left+1 : right])
}
